package service

import (
	"fmt"
	"log"
	"reflect"

	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

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
		fmt.Println(reflect.TypeOf(event))
		fmt.Println(event.Type)

		// Switch on events
		switch event.Type {
		case watch.Added:
			glog.V(3).Infof("Added - PVC Name: %s", pv.Name)
		case watch.Modified:
			glog.V(3).Infof("Modified - PVC Name: %s", pv.Name)
		case watch.Deleted:
			glog.V(3).Infof("Deleted - PVC Name: %s", pv.Name)
		case watch.Error:
			glog.V(3).Infof("Error - PVC Name: %s", pv.Name)
		}
	}
}
