package cost

import (
	"strconv"
	"testing"

	k8sNode "managedkube.com/kube-cost-agent/pkg/metrics/k8s/node"
)

func TestCalculatePodCost(t *testing.T) {

	type testFields struct {
		node           k8sNode.NodeInfo
		podUsageMemory int64
		podUsageCpu    int64
		podCost        PodCost
	}

	testVars := []testFields{{
		node: k8sNode.NodeInfo{
			Name:               "one",
			CpuCapacity:        1000,
			MemoryCapacity:     3885420544,
			ComputeCostPerHour: 0.0475,
		},
		podUsageMemory: 4194304000,
		podUsageCpu:    4000,
		podCost: PodCost{
			MinuteMemory: 0.000427,
			HourMemory:   0.025638,
			DayMemory:    0.615314,
			MonthMemory:  18.459417,
			MinuteCpu:    0.001583,
			HourCpu:      0.095000,
			DayCpu:       2.280000,
			MonthCpu:     68.400000,
		},
	}, {
		node: k8sNode.NodeInfo{
			Name:               "two",
			CpuCapacity:        1000,
			MemoryCapacity:     3885420544,
			ComputeCostPerHour: 0.0475,
		},
		podUsageMemory: 4194304000,
		podUsageCpu:    4000,
		podCost: PodCost{
			MinuteMemory: 0.000427,
			HourMemory:   0.025638,
			DayMemory:    0.615314,
			MonthMemory:  18.459417,
			MinuteCpu:    0.001583,
			HourCpu:      0.095000,
			DayCpu:       2.280000,
			MonthCpu:     68.400000},
	}, {
		node: k8sNode.NodeInfo{
			Name:               "three",
			CpuCapacity:        1000,
			MemoryCapacity:     3885420544,
			ComputeCostPerHour: 0.0475,
		},
		podUsageMemory: 4194304000,
		podUsageCpu:    4000,
		podCost: PodCost{
			MinuteMemory: 0.000427,
			HourMemory:   0.025638,
			DayMemory:    0.615314,
			MonthMemory:  18.459417,
			MinuteCpu:    0.001583,
			HourCpu:      0.095000,
			DayCpu:       2.280000,
			MonthCpu:     68.400000},
	}, {
		node: k8sNode.NodeInfo{
			Name:               "four",
			CpuCapacity:        1000,
			MemoryCapacity:     3885420544,
			ComputeCostPerHour: 0.0475,
		},
		podUsageMemory: 524288000,
		podUsageCpu:    500,
		podCost: PodCost{
			MinuteMemory: 0.000053,
			HourMemory:   0.003205,
			DayMemory:    0.076914,
			MonthMemory:  2.307427,
			MinuteCpu:    0.000198,
			HourCpu:      0.011875,
			DayCpu:       0.285000,
			MonthCpu:     8.550000},
	}}

	for _, v := range testVars {
		//fmt.Println(v.podUsageCpu)

		got := CalculatePodCost(v.node, v.podUsageMemory, v.podUsageCpu)
		want := v.podCost

		if strconv.FormatFloat(got.MinuteMemory, 'f', 6, 64) != strconv.FormatFloat(want.MinuteMemory, 'f', 6, 64) {
			t.Errorf("test: %s - got %.g, want %.g", v.node.Name, got.MinuteMemory, want.MinuteMemory)
		}
		if strconv.FormatFloat(got.HourMemory, 'f', 6, 64) != strconv.FormatFloat(want.HourMemory, 'f', 6, 64) {
			t.Errorf("test: %s - got %.g, want %.g", v.node.Name, got.HourMemory, want.HourMemory)
		}
		if strconv.FormatFloat(got.DayMemory, 'f', 6, 64) != strconv.FormatFloat(want.DayMemory, 'f', 6, 64) {
			t.Errorf("test: %.s - got %.g, want %.g", v.node.Name, got.DayMemory, want.DayMemory)
		}
		if strconv.FormatFloat(got.MonthMemory, 'f', 6, 64) != strconv.FormatFloat(want.MonthMemory, 'f', 6, 64) {
			t.Errorf("test: %s - got %.g, want %.g", v.node.Name, got.MonthMemory, want.MonthMemory)
		}
		if strconv.FormatFloat(got.MinuteCpu, 'f', 6, 64) != strconv.FormatFloat(want.MinuteCpu, 'f', 6, 64) {
			t.Errorf("test: %s - got %.g, want %.g", v.node.Name, got.MinuteCpu, want.MinuteCpu)
		}
		if strconv.FormatFloat(got.HourCpu, 'f', 6, 64) != strconv.FormatFloat(want.HourCpu, 'f', 6, 64) {
			t.Errorf("test: %s - got %.g, want %.g", v.node.Name, got.HourCpu, want.HourCpu)
		}
		if strconv.FormatFloat(got.DayCpu, 'f', 6, 64) != strconv.FormatFloat(want.DayCpu, 'f', 6, 64) {
			t.Errorf("test: %s - got %.g, want %.g", v.node.Name, got.DayCpu, want.DayCpu)
		}
		if strconv.FormatFloat(got.MonthCpu, 'f', 6, 64) != strconv.FormatFloat(want.MonthCpu, 'f', 6, 64) {
			t.Errorf("test: %s - got %.g, want %.g", v.node.Name, got.MonthCpu, want.MonthCpu)
		}
	}
}
