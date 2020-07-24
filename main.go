package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	outTarget := flag.String("t", "kubelife", fmt.Sprintf("output target, supported: %s", allowedOutputTargetFlags))
	flag.Parse()

	outT, err := atoOutputTarget(*outTarget)
	if err != nil {
		log.Fatal(err)
	}

	inputTfstate, err := getTFStateFromStdin()
	if err != nil {
		fmt.Println("Failed to create TFState from stdin: ", err)
		return
	}

	TFHetznerServerResources := extractTFServerResources(
		inputTfstate, "provider.hcloud", "hcloud_server",
	)

	switch outT {
	case ansible:
		//TODO: v2 - input uses an interface => hetzer, DO ...
		ansibleInventory := createAnsibleInventoryHetzner(TFHetznerServerResources)
		yaml, err := yaml.Marshal(ansibleInventory)
		if err != nil {
			fmt.Println("failed to yaml marshal map: ", err)
			return
		}

		fmt.Println(string(yaml))
		return
	case kubelife:
		kubelifePatch := createKubelifePatch(TFHetznerServerResources)
		json, err := json.Marshal(kubelifePatch)
		if err != nil {
			fmt.Println("failed to yaml marshal map: ", err)
			return
		}

		fmt.Println(string(json))
		return
	}
}

type outputTarget int

const (
	ansible outputTarget = iota
	kubelife
)

var allowedOutputTargetFlags = [...]string{"ansible", "kubelife"}

func (of outputTarget) String() string {
	return allowedOutputTargetFlags[of]
}

func atoOutputTarget(name string) (outputTarget, error) {
	index := -1
	for i, allowed := range allowedOutputTargetFlags {
		if name == allowed {
			index = i
		}
	}

	if index == -1 {
		return outputTarget(-1), fmt.Errorf("no allowed output target found for provided name \"%s\", allowed are %s", name, allowedOutputTargetFlags)
	}

	if outputTarget(index).String() != name {
		return outputTarget(-1), fmt.Errorf("found output target for provided name \"%s\", but index doesn't match, got \"%s\"", name, outputTarget(index).String())
	}

	return outputTarget(index), nil
}

func getTFStateFromStdin() (tfstate, error) {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		errorMsg := "The command is intended to work with pipes.\n"
		errorMsg += "Usage:\n"
		errorMsg += "  cat terraform.tfstate | tfto"
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

func extractTFServerResources(inputTfstate tfstate, provider, resourceType string) []TFResource {
	tfresources := inputTfstate.Resources

	extractedTfResources := []TFResource{}
	for _, tfr := range tfresources {
		if (tfr.Provider == provider) && (tfr.Type == resourceType) {
			extractedTfResources = append(extractedTfResources, tfr)
		}
	}

	return extractedTfResources
}
