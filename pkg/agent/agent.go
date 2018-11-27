package agent

import (
	"time"

	"k8s.io/client-go/kubernetes"

	k8sNode "managedkube.com/kube-cost-agent/pkg/metrics/k8s/node"
	k8sPersistentVolume "managedkube.com/kube-cost-agent/pkg/metrics/k8s/persistentVolume"
	k8sPod "managedkube.com/kube-cost-agent/pkg/metrics/k8s/pod"
)

func Update(clientset *kubernetes.Clientset) {

	k8sNode.Register()
	k8sPod.Register()
	k8sPersistentVolume.Register()
	//k8sNamespace.Register()

	go k8sNode.Watch(clientset)
	time.Sleep(5 * time.Second)
	go k8sPod.Watch(clientset)
	go k8sPersistentVolume.Watch(clientset)

}
