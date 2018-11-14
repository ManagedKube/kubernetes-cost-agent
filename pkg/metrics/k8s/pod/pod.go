package pod

import (
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Get list of pods from the k8s API
func GetAllPods(clientset *kubernetes.Clientset) (*v1.PodList, error) {

	//fmt.Println(reflect.TypeOf(clientset))

	// setup list options
	listOptions := metav1.ListOptions{
		LabelSelector: "",
		FieldSelector: "",
	}

	// list pods
	pods, err := clientset.CoreV1().Pods("").List(listOptions)
	if err != nil {
		glog.Errorf("Failed to retrieve pods: %v", err)
		return nil, err
	}

	//fmt.Print(pods.Items[0])
	//fmt.Printf("%+v\n", pods)

	// fmt.Println(reflect.TypeOf(pods))

	return pods, nil
}
