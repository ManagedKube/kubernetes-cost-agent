package persistentVolume

import (
	"fmt"
	"log"
	"strconv"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"managedkube.com/kube-cost-agent/pkg/price"
)

// var pvList PersistentVolumeList

var (
	PersistentVolumeCostMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mk_persisten_volume_cost",
		Help: "ManagedKube - Cost of the persisten volume.",
	},
		[]string{"namespace_name", "persistent_volume_name", "duration", "disk_type", "cost_per_hour", "claim_name", "disk_size"},
	)
)

// Registers the Prometheus metrics
func Register() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(PersistentVolumeCostMetric)
}

// func Get(clientset *kubernetes.Clientset) (PersistentVolumeList, error) {
//
// 	// Resetting
// 	pvList.PersistentVolume = pvList.PersistentVolume[:0]
//
// 	//fmt.Println(price.GetCloud())
//
// 	pv, err := getPV(clientset)
// 	if err != nil {
// 		glog.Errorf("Failed to retrieve PV: %v", err)
// 		return pvList, err
// 	}
//
// 	for _, i := range pv.Items {
//
// 		// Converting from a type "resource.Quantity" to an int64 value
// 		capacity := i.Spec.Capacity[v1.ResourceStorage]
//
// 		var pv PersistentVolume
// 		pv.Name = i.Name
// 		pv.Capacity = capacity.Value()
// 		pv.VolumeName = ""
// 		//pv.StatusPhase = i.Status.Phase
// 		pv.SpecStorageClassName = i.Spec.StorageClassName
// 		pv.Claim.Name = i.Spec.ClaimRef.Name
// 		pv.Claim.Namespace = i.Spec.ClaimRef.Namespace
// 		pv.Claim.Kind = i.Spec.ClaimRef.Kind
// 		pv.CostPerGbHour = price.DiskPricePerHour(pv.SpecStorageClassName)
//
// 		pvList.PersistentVolume = append(pvList.PersistentVolume, pv)
// 	}
//
// 	return pvList, nil
// }
//
// func getPV(clientset *kubernetes.Clientset) (*v1.PersistentVolumeList, error) {
//
// 	// List PV
// 	pv, err := clientset.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
// 	if err != nil {
// 		glog.Errorf("Failed to retrieve PV: %v", err)
// 		return nil, err
// 	}
//
// 	return pv, nil
//
// }

func Watch(clientset *kubernetes.Clientset) {

	watcher, err := clientset.CoreV1().PersistentVolumes().Watch(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Start channel to watch
	ch := watcher.ResultChan()

	// Loop through events
	for event := range ch {
		pv, ok := event.Object.(*v1.PersistentVolume)
		if !ok {
			log.Fatal("unexpected type")
		}
		// fmt.Println(reflect.TypeOf(event))
		// fmt.Println(event.Type)

		// Converting from a type "resource.Quantity" to an int64 value
		capacity := pv.Spec.Capacity[v1.ResourceStorage]

		var persistentVolume PersistentVolume
		persistentVolume.Name = pv.Name
		persistentVolume.Capacity = capacity.Value()
		persistentVolume.VolumeName = ""
		//pv.StatusPhase = i.Status.Phase
		persistentVolume.SpecStorageClassName = pv.Spec.StorageClassName
		persistentVolume.Claim.Name = pv.Spec.ClaimRef.Name
		persistentVolume.Claim.Namespace = pv.Spec.ClaimRef.Namespace
		persistentVolume.Claim.Kind = pv.Spec.ClaimRef.Kind
		persistentVolume.CostPerGbHour = price.DiskPricePerHour(persistentVolume.SpecStorageClassName) / 30 / 24

		// Calculate disk cost
		diskCostPerHour := (float64(persistentVolume.Capacity) / 1000000000) * persistentVolume.CostPerGbHour
		diskCostPerMinute := diskCostPerHour / 60

		// Switch on events
		switch event.Type {
		case watch.Added:
			glog.V(3).Infof("Added - PVC Name: %s", pv.Name)
			PersistentVolumeCostMetric.With(prometheus.Labels{"namespace_name": persistentVolume.Claim.Namespace, "persistent_volume_name": persistentVolume.Name, "duration": "minute", "disk_type": persistentVolume.SpecStorageClassName, "cost_per_hour": fmt.Sprintf("%f", persistentVolume.CostPerGbHour), "claim_name": persistentVolume.Claim.Name, "disk_size": strconv.FormatInt(persistentVolume.Capacity, 10)}).Add(diskCostPerMinute)
		case watch.Modified:
			glog.V(3).Infof("Modified - PVC Name: %s", pv.Name)
		case watch.Deleted:
			glog.V(3).Infof("Deleted - PVC Name: %s", pv.Name)
			PersistentVolumeCostMetric.Delete(prometheus.Labels{"namespace_name": persistentVolume.Claim.Namespace, "persistent_volume_name": persistentVolume.Name, "duration": "minute", "disk_type": persistentVolume.SpecStorageClassName, "cost_per_hour": fmt.Sprintf("%f", persistentVolume.CostPerGbHour), "claim_name": persistentVolume.Claim.Name, "disk_size": strconv.FormatInt(persistentVolume.Capacity, 10)})
		case watch.Error:
			glog.V(3).Infof("Error - PVC Name: %s", pv.Name)
		}
	}
}
