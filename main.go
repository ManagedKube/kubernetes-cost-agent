package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"managedkube.com/kube-cost-agent/pkg/cost"
	"managedkube.com/kube-cost-agent/pkg/export"
	k8sNode "managedkube.com/kube-cost-agent/pkg/metrics/k8s/node"
	k8sPod "managedkube.com/kube-cost-agent/pkg/metrics/k8s/pod"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// optional - local kubeconfig for testing
var kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig file")

func main() {

	// send logs to stderr so we can use 'kubectl logs'
	flag.Set("logtostderr", "true")
	flag.Set("v", "3")
	flag.Parse()

	config, err := getConfig(*kubeconfig)
	if err != nil {
		glog.Errorf("Failed to load client config: %v", err)
		return
	}

	// build the Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Errorf("Failed to create kubernetes client: %v", err)
		return
	}

	export.Register()

	go update(clientset)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9101", nil)
}

func getConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	return rest.InClusterConfig()
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func update(clientset *kubernetes.Clientset) {

	namespaceCostMap := make(map[string]float64)
	var podMetricList k8sPod.PodMetricList
	var prunePodMetricList k8sPod.PodMetricList

	for {
		nodeList, err := k8sNode.AllNodes(clientset)
		if err != nil {
			glog.Errorf("Failed to retrieve nodes: %v", err)
			return
		}

		fmt.Println("nodeList.Node")
		for _, n := range nodeList.Node {
			fmt.Println(n)
		}

		pods, err := k8sPod.GetAllPods(clientset)
		if err != nil {
			glog.Errorf("Failed to retrieve pods: %v", err)
			return
		}

		// Search for pods that was in the last cycle's list and not in the new list this cycle
		for _, pm := range podMetricList.Pod {
			var didFindPod bool = false

			// check if this pod is in the updated pods list
			for _, p := range pods.Items {

				if pm.Namespace_name == p.Namespace {

					if pm.Pod_name == p.Name {
						didFindPod = true
					}
				}
			}

			if !didFindPod {
				// Remove this entry to the remove pod list
				prunePodMetricList.Pod = append(prunePodMetricList.Pod, pm)
			}
		}

		// Remove the metrics
		for _, pm := range prunePodMetricList.Pod {
			glog.V(3).Infof("Removing pod from the export list: %s/%s", pm.Pod_name, pm.Container_name)

			// for each of these items remove them from the prometheus export
			export.RemovePodPrometheus(pm)
		}

		// Resetting slices
		podMetricList.Pod = podMetricList.Pod[:0]
		prunePodMetricList.Pod = prunePodMetricList.Pod[:0]

		//fmt.Println(reflect.TypeOf(pods.Items))

		// loop through each pod and calculate the cost
		for _, p := range pods.Items {

			if p.Status.Phase == "Running" {
				//PrettyPrint(p)
				//fmt.Println(reflect.TypeOf(p))
				glog.V(3).Infof("Found pods: %s/%s/%s/%s", p.Namespace, p.Name, p.UID, p.Spec.NodeName)

				for _, c := range p.Spec.Containers {
					glog.V(3).Infof("Found container: %s", c.Name)

					var cpuLimit int64 = c.Resources.Limits.Cpu().MilliValue()
					var cpuRequest int64 = c.Resources.Requests.Cpu().MilliValue()
					var memoryLimit int64 = c.Resources.Limits.Memory().Value()
					var memoryRequest int64 = c.Resources.Requests.Memory().Value()

					glog.V(3).Infof("CPU Limit: %s", strconv.FormatInt(cpuLimit, 10))
					glog.V(3).Infof("Memory Limit: %s", strconv.FormatInt(memoryLimit, 10))
					glog.V(3).Infof("CPU Requests: %s", strconv.FormatInt(cpuRequest, 10))
					glog.V(3).Infof("Memory Requests: %s", strconv.FormatInt(memoryRequest, 10))

					//fmt.Println(reflect.TypeOf(cpuLimit))

					nodeInfo, err := k8sNode.GetNodeInfo(nodeList, p.Spec.NodeName)
					if err != nil {
						glog.Errorf("Failed to retrieve nodes: %v", err)
						return
					}

					var podUsageMemory int64 = memoryLimit
					var podUsageCpu int64 = cpuLimit

					// Calculate the cost of this container
					podCost := cost.CalculatePodCost(nodeInfo, podUsageMemory, podUsageCpu)

					// Keeping track of last iterations pod list
					// Used for pruning the metrics list
					var metric k8sPod.PodMetric
					metric.Namespace_name = p.Namespace
					metric.Pod_name = p.Name
					metric.Container_name = c.Name
					metric.Duration = "minute"

					podMetricList.Pod = append(podMetricList.Pod, metric)

					// Export pod metrics
					export.Pods(podCost, p, c.Name)

					// Add this pod to the total
					namespaceCostMap[p.Namespace] += podCost.MinuteCpu + podCost.MinuteMemory
				}
			}

		}

		for k, ns := range namespaceCostMap {
			// fmt.Println(k)
			// fmt.Println(strconv.FormatFloat(ns, 'f', 6, 64))
			export.NamespaceCost.With(prometheus.Labels{"namespace_name": k, "duration": "minute"}).Add(ns)

			// reset counter
			namespaceCostMap[k] = 0
		}

		time.Sleep(60 * time.Second)
	}
}
