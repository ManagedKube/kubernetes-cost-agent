package price

import (
	"testing"
)

func TestPricingJSONParsing(t *testing.T) {

	fileLocationPrefix = "./prices/"

	cloud = "test"
	region = "us-east-1"

	loadInstancePricing()

	want := 21

	got := len(instances.Instance)

	if got != want {
		t.Errorf("Expected %v but got %v", want, got)
	}

	// for _, i := range instances.Instance {
	// 	fmt.Println(i.Name)want := "t2.micro"
	// 	fmt.Println(i.Cpu)
	// 	fmt.Println(i.HourlyCost.OnDemand)
	// }
}

func TestPricingJSONParsing2(t *testing.T) {

	fileLocationPrefix = "./prices/"

	cloud = "test"
	region = "ap-southeast-1"

	loadInstancePricing()

	want := 21

	got := len(instances.Instance)

	if got != want {
		t.Errorf("Expected %v but got %v", want, got)
	}

	// for _, i := range instances.Instance {
	// 	fmt.Println(i.Name)want := "t2.micro"
	// 	fmt.Println(i.Cpu)
	// 	fmt.Println(i.HourlyCost.OnDemand)
	// }
}

func TestPricingJSONValues(t *testing.T) {

	fileLocationPrefix = "./prices/"

	instance := Instance{
		Name:   "t2.micro",
		Memory: 1073741824,
		Cpu:    1,
		HourlyCost: HourlyCost{
			OnDemand:   0.0116,
			ReduceCost: 0.0035,
		},
	}

	cloud = "test"
	region = "us-east-1"

	loadInstancePricing()

	if instances.Instance[0].Name != instance.Name {
		t.Errorf("Expected %s got %s", instance.Name, instances.Instance[0].Name)
	}
	if instances.Instance[0].Memory != instance.Memory {
		t.Errorf("Expected %v got %v", instance.Memory, instances.Instance[0].Memory)
	}
	if instances.Instance[0].Cpu != instance.Cpu {
		t.Errorf("Expected %v got %v", instance.Cpu, instances.Instance[0].Cpu)
	}
	if instances.Instance[0].HourlyCost.OnDemand != instance.HourlyCost.OnDemand {
		t.Errorf("Expected %v got %v", instance.HourlyCost.OnDemand, instances.Instance[0].HourlyCost.OnDemand)
	}
	if instances.Instance[0].HourlyCost.ReduceCost != instance.HourlyCost.ReduceCost {
		t.Errorf("Expected %v got %v", instance.HourlyCost.ReduceCost, instances.Instance[0].HourlyCost.ReduceCost)
	}
}

func TestAutoDetectCloud(t *testing.T) {
	labels := make(map[string]string)
	labels["cloud.google.com/gke-os-distribution"] = "cos"
	labels["cloud.google.com/gke-preemptible"] = "true"
	labels["failure-domain.beta.kubernetes.io/region"] = "us-central1"
	labels["failure-domain.beta.kubernetes.io/zone"] = "us-central1-a"
	labels["kubernetes.io/hostname"] = "gke-gar1-default-pool-eb7a3f28-60ht"
	labels["beta.kubernetes.io/instance-type"] = "n1-standard-1"
	labels["beta.kubernetes.io/fluentd-ds-ready"] = "true"
	labels["beta.kubernetes.io/os"] = "linux"
	labels["cloud.google.com/gke-nodepool"] = "default-pool"
	labels["beta.kubernetes.io/arch"] = "amd64"

	got := AutoDetectCloud(labels)

	want := "gcp"

	if want != got {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestNodePricePerHourOnDemand(t *testing.T) {

	fileLocationPrefix = "./prices/"

	region := "us-central1"
	instanceType := "n1-standard-1"
	reduceCostInstance := "false"

	got := NodePricePerHour(region, instanceType, reduceCostInstance)
	want := 0.0475

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestNodePricePerHourReduceCost(t *testing.T) {

	fileLocationPrefix = "./prices/"

	region := "us-central1"
	instanceType := "n1-standard-1"
	reduceCostInstance := "true"

	got := NodePricePerHour(region, instanceType, reduceCostInstance)
	want := 0.01

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}
