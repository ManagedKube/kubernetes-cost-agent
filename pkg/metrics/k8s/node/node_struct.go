package node

type NodeInfo struct {
	Name               string
	ClusterName	       string
	CpuCapacity        int64
	MemoryCapacity     int64
	ComputeCostPerHour float64
	Region             string
	Zone               string
	InstanceType       string
	ReduceCostInstance string
}

type NodeList struct {
	Node []NodeInfo
}
