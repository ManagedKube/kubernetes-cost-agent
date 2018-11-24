package export

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/api/core/v1"
	"managedkube.com/kube-cost-agent/pkg/cost"
	k8sNode "managedkube.com/kube-cost-agent/pkg/metrics/k8s/node"
	k8sPersistenVolume "managedkube.com/kube-cost-agent/pkg/metrics/k8s/persistentVolume"
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
	NodeCostMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mk_node_cost",
		Help: "ManagedKube - Cost of the node.",
	},
		[]string{"node_name", "duration", "instance_type", "cost_per_hour"},
	)
	PersistentVolumeCostMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mk_persisten_volume_cost",
		Help: "ManagedKube - Cost of the persisten volume.",
	},
		[]string{"namespace_name", "persistent_volume_name", "duration", "disk_type", "cost_per_hour", "claim_name"},
	)
)

// Registers the Prometheus metrics
func Register() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(NamespaceCost)
	prometheus.MustRegister(PodCostMetric)
	prometheus.MustRegister(TotalNumberOfPods)
	prometheus.MustRegister(NodeCostMetric)
	prometheus.MustRegister(PersistentVolumeCostMetric)
}

// Update pod metrics
func Pods(podCost cost.PodCost, pod v1.Pod, containerName string) {
	updatePodsPrometheus(podCost, pod, containerName)
}

// Update namespace metric
func Namespace(namespaceName string, duration string, cost float64) {
	updateNamespacePrometheus(namespaceName, duration, cost)
}

// Update node metrics
func Node(nodeInfo k8sNode.NodeInfo) {
	updateNodePrometheus(nodeInfo)
}

// Update persistent volume metrics
func PersistentVolume(persistentVolume k8sPersistenVolume.PersistentVolume) {
	updatePersistentVolumePrometheus(persistentVolume)
}

// Update the prometheus metrics with the new values
func updateNodePrometheus(nodeInfo k8sNode.NodeInfo) {
	NodeCostMetric.With(prometheus.Labels{"node_name": nodeInfo.Name, "duration": "minute", "instance_type": nodeInfo.InstanceType, "cost_per_hour": fmt.Sprintf("%f", nodeInfo.ComputeCostPerHour)}).Add(nodeInfo.ComputeCostPerHour / 60)
}

// Remove a particular metric with these node labels
func RemoveNodePrometheus(nodeInfo k8sNode.NodeInfo) {
	NodeCostMetric.Delete(prometheus.Labels{"node_name": nodeInfo.Name, "duration": "minute", "instance_type": nodeInfo.InstanceType, "cost_per_hour": fmt.Sprintf("%f", nodeInfo.ComputeCostPerHour)})
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

func updatePersistentVolumePrometheus(persistentVolume k8sPersistenVolume.PersistentVolume) {
	PersistentVolumeCostMetric.With(prometheus.Labels{"namespace_name": persistentVolume.Claim.Namespace, "persistent_volume_name": persistentVolume.Name, "duration": "minute", "disk_type": persistentVolume.SpecStorageClassName, "cost_per_hour": fmt.Sprintf("%f", persistentVolume.CostPerGbHour), "claim_name": persistentVolume.Claim.Name}).Add(persistentVolume.CostPerGbHour / 60)
}

func RemovePersistentVolumePrometheus(persistentVolume k8sPersistenVolume.PersistentVolume) {

	PersistentVolumeCostMetric.Delete(prometheus.Labels{"namespace_name": persistentVolume.Claim.Namespace, "persistent_volume_name": persistentVolume.Name, "duration": "minute", "disk_type": persistentVolume.SpecStorageClassName, "cost_per_hour": fmt.Sprintf("%f", persistentVolume.CostPerGbHour), "claim_name": persistentVolume.Claim.Name})
}

func Send() {

}
