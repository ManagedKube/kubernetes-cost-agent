package node

import (
	"fmt"
	"log"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"managedkube.com/kubernetes-cost-agent/pkg/price"
)

var nodeList NodeList

var (
	NodeCostMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mk_node_cost",
		Help: "ManagedKube - Cost of the node.",
	},
		[]string{"node_name", "duration", "instance_type", "cost_per_hour"},
	)
)

// Registers the Prometheus metrics
func Register() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(NodeCostMetric)
}

// From the node list return the NodeInfo which matches the nodeName
func GetNodeInfo(nodeName string) (NodeInfo, error) {

	info := NodeInfo{}

	for _, n := range nodeList.Node {
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

func Watch(clientset *kubernetes.Clientset) {

	timeout := int64(99999999)

	watcher, err := clientset.CoreV1().Nodes().Watch(metav1.ListOptions{TimeoutSeconds: &timeout})
	if err != nil {
		log.Fatal(err)
	}

	// Start channel to watch
	ch := watcher.ResultChan()

	// Loop through events
	for event := range ch {
		n, ok := event.Object.(*v1.Node)
		if !ok {
			log.Fatal("unexpected type")
		}
		// fmt.Println(reflect.TypeOf(event))
		// fmt.Println(event.Type)

		var node NodeInfo
		node.Name = n.Name
		node.CpuCapacity = n.Status.Capacity.Cpu().MilliValue()
		node.MemoryCapacity = n.Status.Capacity.Memory().Value()
		node.Region = getLabelValue(n.Labels, "failure-domain.beta.kubernetes.io/region")
		node.Zone = getLabelValue(n.Labels, "failure-domain.beta.kubernetes.io/zone")
		node.InstanceType = getLabelValue(n.Labels, "beta.kubernetes.io/instance-type")
		node.ReduceCostInstance = getLabelValue(n.Labels, "cloud.google.com/gke-preemptible")
		node.ComputeCostPerHour = price.NodePricePerHour(node.Region, node.InstanceType, node.ReduceCostInstance)

		// Switch on events
		switch event.Type {
		case watch.Added:
			glog.V(3).Infof("Added - Node Name: %s", node.Name)
			addToList(node)
			NodeCostMetric.With(prometheus.Labels{"node_name": node.Name, "duration": "minute", "instance_type": node.InstanceType, "cost_per_hour": fmt.Sprintf("%f", node.ComputeCostPerHour)}).Add(node.ComputeCostPerHour / 60)
		case watch.Modified:
			glog.V(3).Infof("Modified - Node Name: %s", node.Name)
		case watch.Deleted:
			glog.V(3).Infof("Deleted - Node Name: %s", node.Name)
			removeFromList(node)
			NodeCostMetric.Delete(prometheus.Labels{"node_name": node.Name, "duration": "minute", "instance_type": node.InstanceType, "cost_per_hour": fmt.Sprintf("%f", node.ComputeCostPerHour)})
		case watch.Error:
			glog.V(3).Infof("Error - Node Name: %s", node.Name)
		}
	}
}

func addToList(node NodeInfo) {
	isInList := false

	for _, v := range nodeList.Node {
		if v.Name == node.Name {
			isInList = true
		}
	}

	if !isInList {
		nodeList.Node = append(nodeList.Node, node)
	}
}

func removeFromList(node NodeInfo) {
	for index, v := range nodeList.Node {
		if v.Name == node.Name {
			nodeList.Node[index] = nodeList.Node[len(nodeList.Node)-1]
			nodeList.Node = nodeList.Node[:len(nodeList.Node)-1]
		}
	}
}
