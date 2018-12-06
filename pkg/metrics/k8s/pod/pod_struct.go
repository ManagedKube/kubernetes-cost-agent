package pod

type PodMetric struct {
	Namespace_name string
	Pod_name       string
	Container_name string
	Duration       string
	CostCPU        float64
	CostMemory     float64
}

type PodMetricList struct {
	Pod []PodMetric
}
