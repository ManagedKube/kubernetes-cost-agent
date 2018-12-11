package agent

import (
	"bytes"
	"encoding/json"
	"github.com/golang/glog"
	"log"
	"net/http"
	"time"

	"k8s.io/client-go/kubernetes"

	k8sNode "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/node"
	k8sPersistentVolume "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/persistentVolume"
	k8sPod "managedkube.com/kubernetes-cost-agent/pkg/metrics/k8s/pod"
)

var exportCycleSeconds time.Duration = 10
var exportURL = ""
var exportToken = ""
var clusterName = ""

type labels struct {
	ClusterName string `json:"clusterName"`
}

type metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Labels    labels `json:"labels"`
}

func SetExportURL(url string) {
	exportURL = url
}

func SetExportToken(token string) {
	exportToken = token
}

func SetClusterName(name string) {
	clusterName = name
}

func Run(clientset *kubernetes.Clientset) {

	k8sNode.Register()
	k8sPod.Register()
	k8sPersistentVolume.Register()
	//k8sNamespace.Register()

	go k8sNode.Watch(clientset)
	time.Sleep(5 * time.Second)
	go k8sPod.Watch(clientset)
	go k8sPersistentVolume.Watch(clientset)

	if exportURL != "" {
		go export()
	}
}

func export() {
	update()
}

func update() {
	for {
		time.Sleep(exportCycleSeconds * time.Second)
		glog.V(3).Infof("Sending exports")

		sendPods()
		sendNodes()
		sendPersistentDisk()
	}
}

func send(bytesRepresentation []uint8) {

	timeout := time.Duration(5 * time.Second)

	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", exportURL, bytes.NewBuffer(bytesRepresentation))
	req.Header.Add("Apikey", exportToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])

	if resp.StatusCode != 200 {
		glog.V(3).Infof("Error sending export to: %s, StatusCode: %s", exportURL, resp.Status)
	}
}

func sendPods() {
	podList := k8sPod.GetList()

	for _, p := range podList.Pod {

		data := struct {
			ApiVersion string           `json:"apiVersion"`
			Kind       string           `json:"kind"`
			Metadata   metadata         `json:"metadata"`
			Spec       k8sPod.PodMetric `json:"spec"`
		}{
			ApiVersion: "managedkube/v1alpha1",
			Kind:       "PodMetric",
			Metadata: metadata{
				Name:      clusterName,
				Namespace: p.Namespace_name,
				Labels: labels{
					ClusterName: clusterName,
				},
			},
			Spec: p,
		}

		bytesRepresentation, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}

		go send(bytesRepresentation)
	}
}

func sendNodes() {
	nodeList := k8sNode.GetList()

	for _, n := range nodeList.Node {

		data := struct {
			ApiVersion string           `json:"apiVersion"`
			Kind       string           `json:"kind"`
			Metadata   metadata         `json:"metadata"`
			Spec       k8sNode.NodeInfo `json:"spec"`
		}{
			ApiVersion: "managedkube/v1alpha1",
			Kind:       "NodeMetric",
			Metadata: metadata{
				Name:      clusterName,
				Namespace: "",
				Labels: labels{
					ClusterName: clusterName,
				},
			},
			Spec: n,
		}

		bytesRepresentation, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}

		go send(bytesRepresentation)
	}
}

func sendPersistentDisk() {
	pvList := k8sPersistentVolume.GetList()

	for _, n := range pvList.PersistentVolume {

		data := struct {
			ApiVersion string                               `json:"apiVersion"`
			Kind       string                               `json:"kind"`
			Metadata   metadata                             `json:"metadata"`
			Spec       k8sPersistentVolume.PersistentVolume `json:"spec"`
		}{
			ApiVersion: "managedkube/v1alpha1",
			Kind:       "PersistentVolumeeMetric",
			Metadata: metadata{
				Name:      clusterName,
				Namespace: n.Claim.Namespace,
				Labels: labels{
					ClusterName: clusterName,
				},
			},
			Spec: n,
		}

		bytesRepresentation, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}

		go send(bytesRepresentation)
	}
}
