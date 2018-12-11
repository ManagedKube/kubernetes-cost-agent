package agent

import (
	k8sNode "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/node"
	k8sPersistentVolume "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/persistentVolume"
	k8sPod "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/pod"
)

type PodExport struct {
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   metadata         `json:"metadata"`
	Spec       k8sPod.PodMetric `json:"spec"`
}

type NodeExport struct {
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   metadata         `json:"metadata"`
	Spec       k8sNode.NodeInfo `json:"spec"`
}

type PersistentDiskExport struct {
	ApiVersion string                               `json:"apiVersion"`
	Kind       string                               `json:"kind"`
	Metadata   metadata                             `json:"metadata"`
	Spec       k8sPersistentVolume.PersistentVolume `json:"spec"`
}
