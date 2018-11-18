package export

import (
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/api/core/v1"
	"managedkube.com/kube-cost-agent/pkg/cost"
	k8sPod "managedkube.com/kube-cost-agent/pkg/metrics/k8s/pod"
)

var (
	NamespaceCost = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mk_namespace_cost",
		Help: "ManagedKube - Cost of the namespace.",
	},
		[]string{"namespace_name", "duration"},
	)
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
	prometheus.MustRegister(NamespaceCost)
	prometheus.MustRegister(PodCostMetric)
	prometheus.MustRegister(TotalNumberOfPods)
}

// Updates the pods metrics
func Pods(podCost cost.PodCost, pod v1.Pod, containerName string) {
	updatePodsPrometheus(podCost, pod, containerName)
}

// Update the namespace metric
func Namespace(namespaceName string, duration string, cost float64) {
	updateNamespacePrometheus(namespaceName, duration, cost)
}

// Updates the prometheus metric with the new values
func updatePodsPrometheus(podCost cost.PodCost, pod v1.Pod, containerName string) {

	PodCostMetric.With(prometheus.Labels{"namespace_name": pod.Namespace, "pod_name": pod.Name, "container_name": containerName, "duration": "minute"}).Add(podCost.MinuteCpu + podCost.MinuteMemory)
}

// Removes a particular metric with these pod labels
func RemovePodPrometheus(pod k8sPod.PodMetric) {

	PodCostMetric.Delete(prometheus.Labels{"namespace_name": pod.Namespace_name, "pod_name": pod.Pod_name, "container_name": pod.Container_name, "duration": pod.Duration})
}

func updateNamespacePrometheus(namespaceName string, duration string, cost float64) {
	NamespaceCost.With(prometheus.Labels{"namespace_name": namespaceName, "duration": duration}).Add(cost)
}

func Send() {

}
