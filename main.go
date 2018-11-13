package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"managedkube.com/kube-cost-agent/pkg/cost"
	"managedkube.com/kube-cost-agent/pkg/node"

	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	nodes, err := getAllNodes(clientset)

	for _, n := range nodes.Items {
		glog.V(3).Infof("Found nodes: %s/%s", n.Name, n.UID)
	}

	//recordMetrics()
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

func getAllNodes(clientset *kubernetes.Clientset) (*v1.NodeList, error) {

	// list nodes
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		glog.Errorf("Failed to retrieve nodes: %v", err)
		return nil, err
	}

	return nodes, nil
}

func getAllPods(clientset *kubernetes.Clientset) (*v1.PodList, error) {

	//fmt.Println(reflect.TypeOf(clientset))

	// setup list options
	listOptions := metav1.ListOptions{
		LabelSelector: "",
		FieldSelector: "",
	}

	// list pods
	pods, err := clientset.CoreV1().Pods("").List(listOptions)
	if err != nil {
		glog.Errorf("Failed to retrieve pods: %v", err)
		return nil, err
	}

	//fmt.Print(pods.Items[0])
	//fmt.Printf("%+v\n", pods)

	//fmt.Println(reflect.TypeOf(pods))

	return pods, nil
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func update(clientset *kubernetes.Clientset) {

	// divisor := resource.Quantity{}
	// divisor = resource.MustParse("1")

	namespaceCostMap := make(map[string]float64)
	var nodeList node.NodeList
	var podMetricList podMetricList

	for {

		nodes, err := getAllNodes(clientset)
		if err != nil {
			glog.Errorf("Failed to retrieve nodes: %v", err)
			return
		}

		fmt.Println(reflect.TypeOf(nodes))

		for _, n := range nodes.Items {
			//PrettyPrint(n.Status.Capacity)
			glog.V(3).Infof("Found nodes: %s/%s", n.Name, n.UID)

			var node node.NodeInfo
			node.Name = n.Name
			node.CpuCapacity = n.Status.Capacity.Cpu().MilliValue()
			node.MemoryCapacity = n.Status.Capacity.Memory().Value()
			node.ComputeCostPerHour = 0.0475

			glog.V(3).Infof("Node CPU Capacity: %s", strconv.FormatInt(node.CpuCapacity, 10))
			glog.V(3).Infof("Node Memory Capacity: %s", strconv.FormatInt(node.MemoryCapacity, 10))

			nodeList.Node = append(nodeList.Node, node)
		}

		fmt.Println("nodeList.Node")
		for _, n := range nodeList.Node {
			fmt.Println(n)
		}

		pods, err := getAllPods(clientset)
		if err != nil {
			glog.Errorf("Failed to retrieve pods: %v", err)
			return
		}

		// Reset pod metrics counters
		// Would prefer to remove the metrics when it goes to zero.  Havent found a way to do that with
		// the prometheus libs
		for _, p := range podMetricList.pod {
			podCostMetric.With(prometheus.Labels{"namespace_name": p.namespace_name, "pod_name": p.pod_name, "container_name": p.container_name, "duration": "minute"}).Set(0)
			podCostMetric.With(prometheus.Labels{"namespace_name": p.namespace_name, "pod_name": p.pod_name, "container_name": p.container_name, "duration": "hour"}).Set(0)
			podCostMetric.With(prometheus.Labels{"namespace_name": p.namespace_name, "pod_name": p.pod_name, "container_name": p.container_name, "duration": "day"}).Set(0)
			podCostMetric.With(prometheus.Labels{"namespace_name": p.namespace_name, "pod_name": p.pod_name, "container_name": p.container_name, "duration": "month"}).Set(0)
		}

		for _, p := range pods.Items {
			if p.Status.Phase == "Running" {
				//PrettyPrint(p)
				//fmt.Println(reflect.TypeOf(p.Spec.Containers))
				glog.V(3).Infof("Found pods: %s/%s/%s/%s", p.Namespace, p.Name, p.UID, p.Spec.NodeName)

				for _, c := range p.Spec.Containers {
					glog.V(3).Infof("Found container: %s", c.Name)
					//fmt.Println(reflect.TypeOf(c.Resources.Limits.Memory))
					//fmt.Println(reflect.TypeOf(c))
					//PrettyPrint(c.Resources.Limits)
					//fmt.Println(c.Resources.Limits.Memory.Value())
					// for k, l := range c.Resources.Limits {
					// 	fmt.Println(k)
					// 	fmt.Println(l)
					// 	// PrettyPrint(k)
					// 	// PrettyPrint(l)
					// }

					var cpuLimit int64 = c.Resources.Limits.Cpu().MilliValue()
					var cpuRequest int64 = c.Resources.Requests.Cpu().MilliValue()
					var memoryLimit int64 = c.Resources.Limits.Memory().Value()
					var memoryRequest int64 = c.Resources.Requests.Memory().Value()

					glog.V(3).Infof("CPU Limit: %s", strconv.FormatInt(cpuLimit, 10))
					glog.V(3).Infof("Memory Limit: %s", strconv.FormatInt(memoryLimit, 10))
					glog.V(3).Infof("CPU Requests: %s", strconv.FormatInt(cpuRequest, 10))
					glog.V(3).Infof("Memory Requests: %s", strconv.FormatInt(memoryRequest, 10))

					//fmt.Println(reflect.TypeOf(cpuLimit))

					nodeInfo, err := getNodeInfo(nodeList, p.Spec.NodeName)
					if err != nil {
						glog.Errorf("Failed to retrieve nodes: %v", err)
						return
					}

					var podUsageMemory int64 = memoryLimit
					var podUsageCpu int64 = cpuLimit

					//cost := calculatePodCost(nodeInfo, podUsageMemory, podUsageCpu)
					podCost := cost.CalculatePodCost(nodeInfo, podUsageMemory, podUsageCpu)

					podCostMetric.With(prometheus.Labels{"namespace_name": p.Namespace, "pod_name": p.Name, "container_name": c.Name, "duration": "minute"}).Set(podCost.MinuteCpu + podCost.MinuteMemory)
					podCostMetric.With(prometheus.Labels{"namespace_name": p.Namespace, "pod_name": p.Name, "container_name": c.Name, "duration": "hour"}).Set(podCost.HourCpu + podCost.HourMemory)
					podCostMetric.With(prometheus.Labels{"namespace_name": p.Namespace, "pod_name": p.Name, "container_name": c.Name, "duration": "day"}).Set(podCost.DayCpu + podCost.DayMemory)
					podCostMetric.With(prometheus.Labels{"namespace_name": p.Namespace, "pod_name": p.Name, "container_name": c.Name, "duration": "month"}).Set(podCost.MonthCpu + podCost.MonthMemory)

					var metric podMetric
					metric.namespace_name = p.Namespace
					metric.pod_name = p.Name
					metric.container_name = c.Name

					podMetricList.pod = append(podMetricList.pod, metric)

					// Add this pod to the total
					namespaceCostMap[p.Namespace] += podCost.MinuteCpu + podCost.MinuteMemory
				}
			}

		}

		// hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
		// namespaceCost.With(prometheus.Labels{"namespace_name": "foo", "duration": "bar"}).Set(4.2)
		// namespaceCost.With(prometheus.Labels{"namespace_name": "foo2", "duration": "bar"}).Set(5.2)

		for k, ns := range namespaceCostMap {
			// fmt.Println(k)
			// fmt.Println(strconv.FormatFloat(ns, 'f', 6, 64))
			namespaceCost.With(prometheus.Labels{"namespace_name": k, "duration": "minute"}).Set(ns)

			// reset counter
			namespaceCostMap[k] = 0
		}

		time.Sleep(60 * time.Second)
	}
}

// // https://github.com/kubernetes/kubernetes/blob/master/pkg/api/resource/helpers.go
// // convertResourceCPUToInt converts cpu value to the format of divisor and returns
// // ceiling of the value.
// func convertResourceCPUToInt(cpu *resource.Quantity, divisor resource.Quantity) (int64, error) {
// 	c := int64(math.Ceil(float64(cpu.MilliValue()) / float64(divisor.MilliValue())))
// 	//b := float64(math.Ceil(float64(cpu.Value()) / float64(divisor.Value())))
// 	fmt.Println(cpu.MilliValue())
// 	return c, nil
// }
//
// // convertResourceMemoryToInt converts memory value to the format of divisor and returns
// // ceiling of the value.
// func convertResourceMemoryToInt(memory *resource.Quantity, divisor resource.Quantity) (int64, error) {
// 	m := int64(math.Ceil(float64(memory.Value()) / float64(divisor.Value())))
// 	return m, nil
// }
//
// // convertResourceEphemeralStorageToInt converts ephemeral storage value to the format of divisor and returns
// // ceiling of the value.
// func convertResourceEphemeralStorageToInt(ephemeralStorage *resource.Quantity, divisor resource.Quantity) (int64, error) {
// 	m := int64(math.Ceil(float64(ephemeralStorage.Value()) / float64(divisor.Value())))
// 	return m, nil
// }

func recordMetrics() {
	go func() {
		for {
			cpuTemp.Set(65.3)
			hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
			namespaceCost.With(prometheus.Labels{"namespace_name": "foo", "duration": "bar"}).Set(4.2)
			namespaceCost.With(prometheus.Labels{"namespace_name": "foo2", "duration": "bar"}).Set(5.2)
			time.Sleep(2 * time.Second)
		}
	}()
}

type podMetric struct {
	namespace_name string
	pod_name       string
	container_name string
	duration       string
}

type podMetricList struct {
	pod []podMetric
}

func getNodeInfo(nodes node.NodeList, nodeName string) (node.NodeInfo, error) {

	info := node.NodeInfo{}

	for _, n := range nodes.Node {
		if n.Name == nodeName {
			info = n
		}
	}

	return info, nil
}

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors.",
	},
		[]string{"device"},
	)
	namespaceCost = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mk_namespace_cost",
		Help: "ManagedKube - The cost of the namespace.",
	},
		[]string{"namespace_name", "duration"},
	)
	podCostMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mk_pod_cost",
		Help: "ManagedKube - The cost of the pod.",
	},
		[]string{"namespace_name", "pod_name", "container_name", "duration"},
	)
	totalNumberOfPods = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mk_total_number_of_pods",
		Help: "ManagedKube - The total number of running pods.",
	},
		[]string{"namespace_name"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
	prometheus.MustRegister(namespaceCost)
	prometheus.MustRegister(podCostMetric)
	prometheus.MustRegister(totalNumberOfPods)
}
