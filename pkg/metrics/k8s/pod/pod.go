package pod

import (
	"log"
	"strconv"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"managedkube.com/kube-cost-agent/pkg/cost"
	k8sNamespace "managedkube.com/kube-cost-agent/pkg/metrics/k8s/namespace"
	k8sNode "managedkube.com/kube-cost-agent/pkg/metrics/k8s/node"
)

var (
	PodCostMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mk_pod_cost",
		Help: "ManagedKube - Cost of the pod.",
	},
		[]string{"namespace_name", "pod_name", "container_name", "duration"},
	)
	TotalNumberOfPods = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mk_total_number_of_pods",
		Help: "ManagedKube - The total number of running pods.",
	},
		[]string{"namespace_name"},
	)
)

// Registers the Prometheus metrics
func Register() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(PodCostMetric)
	prometheus.MustRegister(TotalNumberOfPods)
}

// Get list of pods from the k8s API
func GetAllPods(clientset *kubernetes.Clientset) (*v1.PodList, error) {

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

	// fmt.Println(reflect.TypeOf(pods))

	return pods, nil
}

func Watch(clientset *kubernetes.Clientset) {

	// setup list options
	listOptions := metav1.ListOptions{
		LabelSelector: "",
		FieldSelector: "",
	}

	watcher, err := clientset.CoreV1().Pods("").Watch(listOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Start channel to watch
	ch := watcher.ResultChan()

	// Loop through events
	for event := range ch {
		p, ok := event.Object.(*v1.Pod)
		if !ok {
			log.Fatal("unexpected type")
		}
		// fmt.Println(reflect.TypeOf(event))
		// fmt.Println(event.Type)

		containerName := ""
		var podCost cost.PodCost

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

				nodeInfo, err := k8sNode.GetNodeInfo(p.Spec.NodeName)
				if err != nil {
					glog.Errorf("Failed to retrieve nodes: %v", err)
					return
				}

				var podUsageMemory int64 = memoryLimit
				var podUsageCpu int64 = cpuLimit

				// Calculate the cost of this container
				podCost = cost.CalculatePodCost(nodeInfo, podUsageMemory, podUsageCpu)

				containerName = c.Name
			}
		}

		// Switch on events
		switch event.Type {
		case watch.Added:
			glog.V(3).Infof("Added - Pod Name: %s", p.Name)

			// Update namespace cost
			k8sNamespace.Add(p.Namespace, podCost.MinuteCpu+podCost.MinuteMemory)
			k8sNamespace.Export()

			PodCostMetric.With(prometheus.Labels{"namespace_name": p.Namespace, "pod_name": p.Name, "container_name": containerName, "duration": "minute"}).Add(podCost.MinuteCpu + podCost.MinuteMemory)

		case watch.Modified:
			glog.V(3).Infof("Modified - Pod Name: %s", p.Name)
		case watch.Deleted:
			glog.V(3).Infof("Deleted - Pod Name: %s", p.Name)

			k8sNamespace.Subtract(p.Namespace, podCost.MinuteCpu+podCost.MinuteMemory)

			PodCostMetric.Delete(prometheus.Labels{"namespace_name": p.Namespace, "pod_name": p.Name, "container_name": containerName, "duration": "minute"})
		case watch.Error:
			glog.V(3).Infof("Error - Pod Name: %s", p.Name)
		}
	}
}
