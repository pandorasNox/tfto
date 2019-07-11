package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// fmt.Println("Hello World")

	inputTfstate, err := getTFStateFromStdin()
	if err != nil {
		fmt.Println("Failed to create TFState from stdin: ", err)
		return
	}
	_ = inputTfstate

	// fmt.Println("inputTfstate:")
	// fmt.Println(inputTfstate)

	err = doAnsibleJson()
	if err != nil {
		fmt.Println("failed doAnsibleJson", err)
		return
	}

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

func doAnsibleJson() error {
	hosts := make(map[string]host)
	hosts["example-host-name1"] = host{"10.11.12.11", "root"}
	hosts["example-host-name2"] = host{"10.11.12.11", "root"}

	json, err := json.Marshal(hosts)
	if err != nil {
		return fmt.Errorf("failed to marshal map: %s", err)
	}
	_ = json

	// fmt.Println("json:")
	// fmt.Println(string(json))
	return nil
}
