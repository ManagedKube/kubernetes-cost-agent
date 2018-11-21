package node

import (
	"strconv"

	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"managedkube.com/kube-cost-agent/pkg/price"
)

var nodeList NodeList

// Retrieves all of the nodes in a k8s cluster
func AllNodes(clientset *kubernetes.Clientset) (NodeList, error) {

	// Resetting NodeList
	nodeList.Node = nodeList.Node[:0]

	nodes, err := getNodes(clientset)
	if err != nil {
		glog.Errorf("Failed to retrieve nodes: %v", err)
		return nodeList, err
	}

	for _, n := range nodes.Items {
		//PrettyPrint(n.Status.Capacity)
		glog.V(3).Infof("Found nodes: %s/%s", n.Name, n.UID)
		//fmt.Println(reflect.TypeOf(n.Labels))

		var node NodeInfo
		node.Name = n.Name
		node.CpuCapacity = n.Status.Capacity.Cpu().MilliValue()
		node.MemoryCapacity = n.Status.Capacity.Memory().Value()
		node.Region = getLabelValue(n.Labels, "failure-domain.beta.kubernetes.io/region")
		node.Zone = getLabelValue(n.Labels, "failure-domain.beta.kubernetes.io/zone")
		node.InstanceType = getLabelValue(n.Labels, "beta.kubernetes.io/instance-type")
		node.ReduceCostInstance = getLabelValue(n.Labels, "cloud.google.com/gke-preemptible")
		node.ComputeCostPerHour = price.NodePricePerHour(node.Region, node.InstanceType, node.ReduceCostInstance)

		glog.V(3).Infof("Node CPU Capacity: %s", strconv.FormatInt(node.CpuCapacity, 10))
		glog.V(3).Infof("Node Memory Capacity: %s", strconv.FormatInt(node.MemoryCapacity, 10))
		glog.V(3).Infof("Node HourlyCost: %v", node.ComputeCostPerHour)

		nodeList.Node = append(nodeList.Node, node)

	}

	return nodeList, nil
}

// Get list of nodes from k8s API
func getNodes(clientset *kubernetes.Clientset) (*v1.NodeList, error) {

	// list nodes
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		glog.Errorf("Failed to retrieve nodes: %v", err)
		return nil, err
	}

	return nodes, nil
}

// From the node list return the NodeInfo which matches the nodeName
func GetNodeInfo(nodes NodeList, nodeName string) (NodeInfo, error) {

	info := NodeInfo{}

	for _, n := range nodes.Node {
		if n.Name == nodeName {
			info = n
		}
	}

	return info, nil
}

func getLabelValue(labels map[string]string, labelName string) string {

	labelKey := ""

	for key, val := range labels {
		if key == labelName {
			labelKey = val
		}
	}

	return labelKey
}
