package persistentVolume

import (
	"fmt"

	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"managedkube.com/kube-cost-agent/pkg/price"
)

var pvList PersistentVolumeList

func Get(clientset *kubernetes.Clientset) (PersistentVolumeList, error) {

	// Resetting
	pvList.PersistentVolume = pvList.PersistentVolume[:0]

	fmt.Println(price.GetCloud())

	pv, err := getPV(clientset)
	if err != nil {
		glog.Errorf("Failed to retrieve PV: %v", err)
		return pvList, err
	}

	for _, i := range pv.Items {

		// Converting from a type "resource.Quantity" to an int64 value
		capacity := i.Spec.Capacity[v1.ResourceStorage]

		var pv PersistentVolume
		pv.Name = i.Name
		pv.Capacity = capacity.Value()
		pv.VolumeName = ""
		//pv.StatusPhase = i.Status.Phase
		pv.SpecStorageClassName = i.Spec.StorageClassName
		pv.Claim.Name = i.Spec.ClaimRef.Name
		pv.Claim.Namespace = i.Spec.ClaimRef.Namespace
		pv.Claim.Kind = i.Spec.ClaimRef.Kind
		pv.CostPerGbHour = price.DiskPricePerHour(pv.SpecStorageClassName)

		pvList.PersistentVolume = append(pvList.PersistentVolume, pv)
	}

	return pvList, nil
}

func getPV(clientset *kubernetes.Clientset) (*v1.PersistentVolumeList, error) {

	// List PV
	pv, err := clientset.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
	if err != nil {
		glog.Errorf("Failed to retrieve PV: %v", err)
		return nil, err
	}

	return pv, nil

}
