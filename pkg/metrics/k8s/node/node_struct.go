package node

type NodeInfo struct {
	Name               string
	CpuCapacity        int64
	MemoryCapacity     int64
	ComputeCostPerHour float64
}

type NodeList struct {
	Node []NodeInfo
}
