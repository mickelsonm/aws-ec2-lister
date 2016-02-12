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
		log.Fatal(err)
	}

	showName := func(tags []EC2InstanceTag) string {
		if len(tags) > 0 {
			for _, tag := range tags {
				if tag.Key == "Name" {
					return tag.Value
				}
			}
		}
		return "---"
	}

	var output string
	for _, res := range resp.Reservations {
		if output == "" {
			output = fmt.Sprintf("%30s %16s %16s %16s\n\n",
				"Server",
				"ID",
				"Private IP",
				"Public IP",
			)
		}
		for _, inst := range res.Instances {
			output += fmt.Sprintf("%30s %16s %16s %16s\n",
				showName(inst.Tags),
				inst.ID,
				inst.PrivateIP,
				inst.PublicIP,
			)
		}
	}

	if output != "" {
		fmt.Println(output)
	}
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
}

//EC2InstanceTag ...
type EC2InstanceTag struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
