package agent

import (
	k8sNode "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/node"
	k8sPersistentVolume "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/persistentVolume"
	k8sPod "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/pod"
)

type Labels struct {
	ClusterName string `json:"clusterName"`
}

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Labels    Labels `json:"labels"`
}

type PodExport struct {
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   Metadata         `json:"metadata"`
	Spec       k8sPod.PodMetric `json:"spec"`
}

type NodeExport struct {
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   Metadata         `json:"metadata"`
	Spec       k8sNode.NodeInfo `json:"spec"`
}

type PersistentDiskExport struct {
	ApiVersion string                               `json:"apiVersion"`
	Kind       string                               `json:"kind"`
	Metadata   Metadata                             `json:"metadata"`
	Spec       k8sPersistentVolume.PersistentVolume `json:"spec"`
}
