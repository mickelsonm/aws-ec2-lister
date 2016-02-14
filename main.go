package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var resp EC2Response

	err := json.NewDecoder(os.Stdin).Decode(&resp)
	if err != nil {
		log.Fatal(fmt.Errorf("error reading json string: %s", err))
	}

	var output string
	for _, res := range resp.Reservations {
		if output == "" {
			output = fmt.Sprintf("%30s%14s%11s%16s%18s\n\n",
				"Server",
				"ID",
				"State",
				"Private IP",
				"Public IP",
			)
		}
		for _, inst := range res.Instances {
			output += fmt.Sprintf("%30s%14s%11s%16s%18s\n",
				getInstanceNameFromTags(inst.Tags),
				inst.ID,
				inst.State.Name,
				inst.PrivateIP,
				inst.PublicIP,
			)
		}
	}

	if output != "" {
		fmt.Println(output)
	}
}

//getInstanceNameFromTags ...
func getInstanceNameFromTags(tags []EC2InstanceTag) string {
	if len(tags) > 0 {
		for _, tag := range tags {
			if tag.Key == "Name" {
				return tag.Value
			}
		}
	}
	return "---"
}

//EC2Response ...
type EC2Response struct {
	Reservations []EC2Result `json:"Reservations"`
}

//EC2Result ...
type EC2Result struct {
	Instances []EC2Instance `json:"Instances"`
}

//EC2Instance ...
type EC2Instance struct {
	ID        string           `json:"InstanceId"`
	PrivateIP string           `json:"PrivateIpAddress"`
	PublicIP  string           `json:"PublicIpAddress"`
	Tags      []EC2InstanceTag `json:"Tags"`
	State     EC2InstanceState `json:"State"`
}

//EC2InstanceState ...
type EC2InstanceState struct {
	Name string `json:"Name"`
}

//EC2InstanceTag ...
type EC2InstanceTag struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
