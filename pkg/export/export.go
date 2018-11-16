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

func updatePodsPrometheus(podCost cost.PodCost, pod v1.Pod, containerName string) {

	PodCostMetric.With(prometheus.Labels{"namespace_name": pod.Namespace, "pod_name": pod.Name, "container_name": containerName, "duration": "minute"}).Add(podCost.MinuteCpu + podCost.MinuteMemory)
	PodCostMetric.With(prometheus.Labels{"namespace_name": pod.Namespace, "pod_name": pod.Name, "container_name": containerName, "duration": "hour"}).Add(podCost.HourCpu + podCost.HourMemory)
	PodCostMetric.With(prometheus.Labels{"namespace_name": pod.Namespace, "pod_name": pod.Name, "container_name": containerName, "duration": "day"}).Add(podCost.DayCpu + podCost.DayMemory)
	PodCostMetric.With(prometheus.Labels{"namespace_name": pod.Namespace, "pod_name": pod.Name, "container_name": containerName, "duration": "month"}).Add(podCost.MonthCpu + podCost.MonthMemory)
}

func resetPodsPrometheus(k8sPod.PodMetricList) {

}

func Send() {

}
