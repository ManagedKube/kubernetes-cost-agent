package price

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var cloud string
var region string
var instances Instances

// Load the instance pricing
func loadInstancePricing() {

	// Open jsonFile
	jsonFile, err := os.Open("./prices/" + cloud + "/" + region + "/instance.json")

	if err != nil {
		panic(err)
	}

	// read opened json file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// unmarshal our byteArray which contains the json
	json.Unmarshal(byteValue, &instances)

	defer jsonFile.Close()
}
