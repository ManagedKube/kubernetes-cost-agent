package pod

type PodMetric struct {
	Namespace_name string            `json:"namespace"`
	Pod_name       string            `json:"podName"`
	Container_name string            `json:"containerName"`
	Duration       string            `json:"duration"`
	CostCPU        float64           `json:"costCpu"`
	CostMemory     float64           `json:"costMemory"`
	Labels         map[string]string `json:"labels"`
}

type PodMetricList struct {
	Pod []PodMetric `json:"pod"`
}
