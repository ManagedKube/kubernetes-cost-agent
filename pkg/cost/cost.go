package cost

import (
	"strconv"

	"github.com/golang/glog"
	k8sNode "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/node"
)

func CalculatePodCost(node k8sNode.NodeInfo, podUsageMemory int64, podUsageCpu int64) PodCost {

	cost := PodCost{}

	computeCostPerHourMemory := node.ComputeCostPerHour * 0.5
	computeCostPerHourCpu := node.ComputeCostPerHour * 0.5

	percentUsedMemory := float64(podUsageMemory) / float64(node.MemoryCapacity)
	percentUsedCpu := float64(podUsageCpu) / float64(node.CpuCapacity)

	cost.HourMemory = computeCostPerHourMemory * float64(percentUsedMemory)
	cost.HourCpu = computeCostPerHourCpu * float64(percentUsedCpu)

	cost.MinuteMemory = cost.HourMemory / 60
	cost.MinuteCpu = cost.HourCpu / 60

	cost.DayMemory = cost.HourMemory * 24
	cost.DayCpu = cost.HourCpu * 24

	cost.MonthMemory = cost.DayMemory * 30
	cost.MonthCpu = cost.DayCpu * 30

	glog.V(3).Infof("Cost per minute memory: %s", strconv.FormatFloat(cost.MinuteMemory, 'f', 6, 64))
	glog.V(3).Infof("Cost per minute cpu: %s", strconv.FormatFloat(cost.MinuteCpu, 'f', 6, 64))

	glog.V(3).Infof("Cost per hour memory: %s", strconv.FormatFloat(cost.HourMemory, 'f', 6, 64))
	glog.V(3).Infof("Cost per hour cpu: %s", strconv.FormatFloat(cost.HourCpu, 'f', 6, 64))

	glog.V(3).Infof("Cost per day memory: %s", strconv.FormatFloat(cost.DayMemory, 'f', 6, 64))
	glog.V(3).Infof("Cost per day cpu: %s", strconv.FormatFloat(cost.DayCpu, 'f', 6, 64))

	glog.V(3).Infof("Cost per month memory: %s", strconv.FormatFloat(cost.MonthMemory, 'f', 6, 64))
	glog.V(3).Infof("Cost per month cpu: %s", strconv.FormatFloat(cost.MonthCpu, 'f', 6, 64))

	return cost
}
