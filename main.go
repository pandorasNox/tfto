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

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("  cat terraform.tfstate | tftoinv")
		return
	}

	var inputBufer bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputBufer.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return
	}

	// fmt.Printf("inputBufer:\n%v\n", inputBufer.String())

	inputTfstate := &tfstate{}
	if err := json.Unmarshal([]byte(inputBufer.String()), inputTfstate); err != nil {
		fmt.Println("failed to Unmarshal stdin as an terraform.state json")
		// return fmt.Errorf("failed to Unmarshal Pod from incoming AdmissionReview: %s", err)
		return
	}

	fmt.Println("inputTfstate:")
	fmt.Println(inputTfstate)
}
