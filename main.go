package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

func main() {
	// fmt.Println("Hello World")

	inputTfstate, err := getTFStateFromStdin()
	if err != nil {
		fmt.Println("Failed to create TFState from stdin: ", err)
		return
	}
	// fmt.Println("inputTfstate:")
	// fmt.Println(inputTfstate)

	TFHetznerServerResources := extractTFHetznerServerResources(inputTfstate)
	// fmt.Println("TFHetznerServerResources:")
	// fmt.Println(TFHetznerServerResources)

	//TODO: v2 - input uses an interface => hetzer, DO ...
	ansibleInventory := createAnsibleInventoryHetzner(TFHetznerServerResources)
	// fmt.Println("ansibleInventory:")
	// fmt.Println(ansibleInventory)

	yaml, err := yaml.Marshal(ansibleInventory)
	if err != nil {
		fmt.Println("failed to yaml marshal map: ", err)
		return
	}

	fmt.Println(string(yaml))
}

func getTFStateFromStdin() (tfstate, error) {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		errorMsg := "The command is intended to work with pipes.\n"
		errorMsg += "Usage:\n"
		errorMsg += "  cat terraform.tfstate | tftoinv"
		return tfstate{}, fmt.Errorf(errorMsg)
	}

	var inputBuffer bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputBuffer.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return tfstate{}, fmt.Errorf("fail to read stdin: %s", err)
	}

	// fmt.Printf("inputBuffer:\n%v\n", inputBuffer.String())

	inputTfstate := &tfstate{}
	if err := json.Unmarshal(inputBuffer.Bytes(), inputTfstate); err != nil {
		return tfstate{}, fmt.Errorf("failed to Unmarshal stdin as an terraform.state json: %s", err)
	}

	return *inputTfstate, nil
}

func extractTFHetznerServerResources(inputTfstate tfstate) []TFResource {
	tfresources := inputTfstate.Resources

	provider := "provider.hcloud"
	resourceType := "hcloud_server"

	tfHetznerResources := []TFResource{}
	for _, tfr := range tfresources {
		if (tfr.Provider == provider) && (tfr.Type == resourceType) {
			tfHetznerResources = append(tfHetznerResources, tfr)
		}
	}

	return tfHetznerResources
}

func createAnsibleInventoryHetzner(hetznerServerResources []TFResource) AnsibleInventory {
	AnsibleInventory := make(map[string]Group)

	for _, hr := range hetznerServerResources {
		instance := hr.Instances[0]

		//extract group names of hetzner
		aig := instance.Attributes.Labels.AnsibleInventoryGroups
		ansibleGroupNames := strings.Split(aig, ".")

		name := instance.Attributes.Name
		hostAddress := instance.Attributes.Ipv4Address

		for _, groupName := range ansibleGroupNames {
			_, groupExists := AnsibleInventory[groupName]
			if !groupExists {
				newHosts := make(map[string]Host)
				AnsibleInventory[groupName] = Group{Hosts: newHosts}
			}
			AnsibleInventory[groupName].Hosts[name] = Host{hostAddress, "root"}
		}
	}

	return AnsibleInventory
}

func doAnsibleJson() error {
	hosts := make(map[string]Host)
	hosts["example-host-name1"] = Host{"10.11.12.11", "root"}
	hosts["example-host-name2"] = Host{"10.11.12.11", "root"}

	json, err := json.Marshal(hosts)
	if err != nil {
		return fmt.Errorf("failed to marshal map: %s", err)
	}
	_ = json

	// fmt.Println("json:")
	// fmt.Println(string(json))
	return nil
}
