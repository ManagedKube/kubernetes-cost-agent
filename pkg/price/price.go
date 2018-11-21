package price

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
)

var cloud string
var region string
var instances Instances

var gcpRegions = [...]string{
	"asia-east1",
	"asia-northeast1",
	"asia-southeast1",
	"europe-west1",
	"europe-west3",
	"northamerica-northeast1",
	"us-central1",
	"us-east4",
	"us-west2",
	"asia-east2",
	"asia-south1",
	"australia-southeast1",
	"europe-west2",
	"europe-west4",
	"southamerica-east1",
	"us-east1",
	"us-west1",
}

var awsRegions = [...]string{
	"ap-northeast-1",
	"ap-northeast-3",
	"ap-southeast-1",
	"ca-central-1",
	"cn-northwest-1",
	"eu-west-1",
	"sa-east-1",
	"us-east-2",
	"us-gov-west-1",
	"us-west-2",
	"ap-northeast-2",
	"ap-south-1",
	"ap-southeast-2",
	"cn-north-1",
	"eu-central-1",
	"eu-west-2",
	"us-east-1",
	"us-gov-east-1",
	"us-west-1",
	"us-west-3",
}

var fileLocationPrefix = "./pkg/price/prices/"

// Load the instance pricing
func loadInstancePricing() {

	// Clear instances
	instances = Instances{}

	// Set path to work in main.go and the tests
	var fileLocation = fileLocationPrefix + cloud + "/" + region + "/instance.json"
	glog.V(3).Infof("Opening price file: %s", fileLocation)

	// Open jsonFile
	jsonFile, err := os.Open(fileLocation)

	if err != nil {
		panic(err)
	}

	// read opened json file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// unmarshal our byteArray which contains the json
	json.Unmarshal(byteValue, &instances)

	defer jsonFile.Close()
}

// Detect the cloud based on the Kubernetes node meta data info
func AutoDetectCloud(labels map[string]string) string {

	// for key, val := range labels {
	// 	fmt.Println(key)
	// 	fmt.Println(val)
	// }

	return "gcp"
}

func NodePricePerHour(regionLocal string, instanceType string, reduceCostInstance string) float64 {

	region = regionLocal

	nodePricePerHour := 0.00

	// Check if the region is in the cloud list
	for _, i := range gcpRegions {
		if region == i {
			glog.V(3).Infof("This is in the GCP Region")
			cloud = "gcp"
		}
	}

	for _, i := range awsRegions {
		if region == i {
			glog.V(3).Infof("This is in the AWS Region")
			cloud = "gcp"
		}
	}

	loadInstancePricing()

	// Find the instance type and get the price
	for _, i := range instances.Instance {
		if i.Name == instanceType {
			if reduceCostInstance == "true" {
				nodePricePerHour = i.HourlyCost.ReduceCost
			} else {
				nodePricePerHour = i.HourlyCost.OnDemand
			}
		}
	}

	return nodePricePerHour
}
