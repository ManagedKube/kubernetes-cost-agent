package agent

import (
	"strconv"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"managedkube.com/kube-cost-agent/pkg/cost"
	"managedkube.com/kube-cost-agent/pkg/export"

	k8sNode "managedkube.com/kube-cost-agent/pkg/metrics/k8s/node"
	k8sPersistentVolume "managedkube.com/kube-cost-agent/pkg/metrics/k8s/persistentVolume"
	k8sPod "managedkube.com/kube-cost-agent/pkg/metrics/k8s/pod"
)

func Update(clientset *kubernetes.Clientset) {

	namespaceCostMap := make(map[string]float64)
	var podMetricList k8sPod.PodMetricList
	var prunePodMetricList k8sPod.PodMetricList

	var nodeMetricList k8sNode.NodeList
	var pruneNodeMetricList k8sNode.NodeList

	var persistentVolumeMetricList k8sPersistentVolume.PersistentVolumeList
	var prunePersistenVolumeMetricList k8sPersistentVolume.PersistentVolumeList

	// Main run loop
	for {
		nodeList, err := k8sNode.AllNodes(clientset)
		if err != nil {
			glog.Errorf("Failed to retrieve nodes: %v", err)
			return
		}

		pvcList, err := k8sPersistentVolume.Get(clientset)
		// pvList, err := k8sPersistentVolume.Get(clientset)
		if err != nil {
			glog.Errorf("Failed to retrieve PV: %v", err)
			return
		}

		// Prune node metrics
		for _, nm := range nodeMetricList.Node {
			var didFindNode bool = false

			// check if this node is in the updated node list
			for _, n := range nodeList.Node {

				if nm.Name == n.Name {
					didFindNode = true
				}
			}

			if !didFindNode {
				// Remove this entry to the remove node list
				pruneNodeMetricList.Node = append(pruneNodeMetricList.Node, nm)
			}
		}

		// Remove the node metrics
		for _, nm := range pruneNodeMetricList.Node {
			glog.V(3).Infof("Removing node from the export list: %s", nm.Name)

			// for each of these items remove them from the prometheus export
			export.RemoveNodePrometheus(nm)
		}

		// Prune persistent volume metrics
		for _, i := range persistentVolumeMetricList.PersistentVolume {
			var didFindNode bool = false

			// check if this node is in the updated node list
			for _, n := range pvcList.PersistentVolume {

				if i.Name == n.Name {
					didFindNode = true
				}
			}

			if !didFindNode {
				// Add this entry to the remove list
				prunePersistenVolumeMetricList.PersistentVolume = append(prunePersistenVolumeMetricList.PersistentVolume, i)
			}
		}

		// Remove the persistent volume metric
		for _, i := range prunePersistenVolumeMetricList.PersistentVolume {
			glog.V(3).Infof("Removing persistent volume from the export list: %s/%s", i.Name, i.Claim.Name)

			// for each of these items remove them from the prometheus export
			export.RemovePersistentVolumePrometheus(i)
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

		// Remove the pod metrics
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

		// export node metrics
		for _, n := range nodeList.Node {
			export.Node(n)

			// Adding to the list to be used for comparing for prune
			nodeMetricList.Node = append(nodeMetricList.Node, n)
		}

		// export namespace metrics
		for k, ns := range namespaceCostMap {
			export.Namespace(k, "minute", ns)

			// reset counter
			namespaceCostMap[k] = 0
		}

		// export persistent volue metric
		for _, i := range pvcList.PersistentVolume {
			export.PersistentVolume(i)

			// Adding to the list to be used for comparing for prune
			persistentVolumeMetricList.PersistentVolume = append(persistentVolumeMetricList.PersistentVolume, i)
		}

		time.Sleep(60 * time.Second)
	}
}
