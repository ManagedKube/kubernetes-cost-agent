package service

type Service struct {
	Name      string
	Namespace string
	ClusterIp string
}

type ServiceList struct {
	Service []Service
}
