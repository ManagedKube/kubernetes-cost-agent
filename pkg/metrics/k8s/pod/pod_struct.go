package pod

type PodMetric struct {
	Namespace_name string
	Pod_name       string
	Container_name string
	Duration       string
}

type PodMetricList struct {
	Pod []PodMetric
}
