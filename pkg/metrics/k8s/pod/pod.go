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
	"managedkube.com/kubernetes-cost-agent/pkg/cost"
	k8sNode "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/node"
)

var podList PodMetricList

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

func GetList() PodMetricList {
	return podList
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

	timeout := int64(99999999)

	// setup list options
	listOptions := metav1.ListOptions{
		LabelSelector:  "",
		FieldSelector:  "",
		TimeoutSeconds: &timeout,
	}

	watcher, err := clientset.CoreV1().Pods("").Watch(listOptions)
	if err != nil {
		glog.V(3).Infof("Watch Error...")
		log.Fatal(err)
		//utilruntime.HandleError(err)
		//os.Exit(3)
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

		// Switch on events
		switch event.Type {
		case watch.Added:
			glog.V(3).Infof("Added - Pod Name: %s", p.Name)

			if p.Status.Phase == "Running" {
				//PrettyPrint(p)
				//fmt.Println(reflect.TypeOf(p))
				glog.V(3).Infof("Found pods: %s/%s/%s/%s", p.Namespace, p.Name, p.UID, p.Spec.NodeName)

				for _, c := range p.Spec.Containers {
					glog.V(3).Infof("Found container: %s", c.Name)

					var podMetric PodMetric
					podMetric.Namespace_name = p.Namespace
					podMetric.Pod_name = p.Name
					podMetric.Container_name = c.Name
					podMetric.Duration = "minute"

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
					podCost := cost.CalculatePodCost(nodeInfo, podUsageMemory, podUsageCpu)

					podMetric.CostCPU = podCost.MinuteCpu
					podMetric.CostMemory = podCost.MinuteMemory

					PodCostMetric.With(prometheus.Labels{"namespace_name": podMetric.Namespace_name, "pod_name": podMetric.Pod_name, "container_name": podMetric.Container_name, "duration": podMetric.Duration}).Add(podCost.MinuteCpu + podCost.MinuteMemory)

					addToListPodMetricList(podMetric)
				}
			}

		case watch.Modified:
			glog.V(3).Infof("Modified - Pod Name: %s", p.Name)

			if p.Status.Phase == "Running" {
				//PrettyPrint(p)
				//fmt.Println(reflect.TypeOf(p))
				glog.V(3).Infof("Found pods: %s/%s/%s/%s", p.Namespace, p.Name, p.UID, p.Spec.NodeName)

				for _, c := range p.Spec.Containers {
					glog.V(3).Infof("Found container: %s", c.Name)

					var podMetric PodMetric
					podMetric.Namespace_name = p.Namespace
					podMetric.Pod_name = p.Name
					podMetric.Container_name = c.Name
					podMetric.Duration = "minute"

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
					podCost := cost.CalculatePodCost(nodeInfo, podUsageMemory, podUsageCpu)

					podMetric.CostCPU = podCost.MinuteCpu
					podMetric.CostMemory = podCost.MinuteMemory

					PodCostMetric.With(prometheus.Labels{"namespace_name": podMetric.Namespace_name, "pod_name": podMetric.Pod_name, "container_name": podMetric.Container_name, "duration": podMetric.Duration}).Add(podCost.MinuteCpu + podCost.MinuteMemory)

					addToListPodMetricList(podMetric)
				}
			}

		case watch.Deleted:
			glog.V(3).Infof("Deleted - Pod Name: %s", p.Name)

			if p.Status.Phase == "Running" {
				//PrettyPrint(p)
				//fmt.Println(reflect.TypeOf(p))
				glog.V(3).Infof("Found pods: %s/%s/%s/%s", p.Namespace, p.Name, p.UID, p.Spec.NodeName)

				for _, c := range p.Spec.Containers {
					glog.V(3).Infof("Found container: %s", c.Name)

					var podMetric PodMetric
					podMetric.Namespace_name = p.Namespace
					podMetric.Pod_name = p.Name
					podMetric.Container_name = c.Name
					podMetric.Duration = "minute"

					//PodCostMetric.Delete(prometheus.Labels{"namespace_name": p.Namespace, "pod_name": p.Name, "container_name": c.Name, "duration": "minute"})
					PodCostMetric.Delete(prometheus.Labels{"namespace_name": podMetric.Namespace_name, "pod_name": podMetric.Pod_name, "container_name": podMetric.Container_name, "duration": podMetric.Duration})

					removeFromPodMetricList(podMetric)
				}
			}

		case watch.Error:
			glog.V(3).Infof("Error - Pod Name: %s", p.Name)
		}
	}
}

func addToListPodMetricList(podMetric PodMetric) {

	isInList := false

	for _, v := range podList.Pod {

		if v.Namespace_name == podMetric.Namespace_name {
			if v.Pod_name == podMetric.Pod_name {
				if v.Container_name == podMetric.Container_name {
					isInList = true
				}
			}
		}
	}

	if !isInList {
		podList.Pod = append(podList.Pod, podMetric)
	}
}

func removeFromPodMetricList(podMetric PodMetric) {

	for index, i := range podList.Pod {

		if i.Namespace_name == podMetric.Namespace_name {
			if i.Pod_name == podMetric.Pod_name {
				if i.Container_name == podMetric.Container_name {
					// Remove item
					podList.Pod[index] = podList.Pod[len(podList.Pod)-1]
					podList.Pod = podList.Pod[:len(podList.Pod)-1]
				}
			}
		}
	}
}
