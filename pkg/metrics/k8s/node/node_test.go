package node

import (
	"testing"
)

func TestGetLabelValues(t *testing.T) {

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

	got := getLabelValue(labels, "failure-domain.beta.kubernetes.io/region")
	want := "us-central1"

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}
