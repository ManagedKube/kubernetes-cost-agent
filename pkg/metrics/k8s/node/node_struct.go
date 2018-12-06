package node

type NodeInfo struct {
	Name               string  `json:"name"`
	CpuCapacity        int64   `json:"cpuCapacity"`
	MemoryCapacity     int64   `json:"memoryCapacity"`
	ComputeCostPerHour float64 `json:"computeCostPerHour"`
	Region             string  `json:"region"`
	Zone               string  `json:"zone"`
	InstanceType       string  `json:"instanceType"`
	ReduceCostInstance string  `json:"reduceCostInstance"`
}

type NodeList struct {
	Node []NodeInfo `json:"node"`
}
