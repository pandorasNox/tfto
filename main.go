package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("  cat yourfile.txt | searchr -pattern=<your_pattern>")
		return
	}

	// reader := bufio.NewReader(os.Stdin)

	// line := 1
	// for {
	// 	input, err := reader.ReadString('\n')
	// 	if err != nil && err == io.EOF {
	// 		break
	// 	}

	// 	// color := "\x1b[39m"
	// 	// if strings.Contains(input, pattern) {
	// 	// 	color = "\x1b[31m"
	// 	// }

	// 	// fmt.Printf("%s%2d: %s", color, line, input)
	// 	fmt.Println("line: ", line, " | contains: ", input)
	// 	line++
	// }
	// // match(*pattern, reader)

	holeInput := ""
	var b bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		holeInput = holeInput + scanner.Text()
		b.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("hole input:\n%v\n", holeInput)
	fmt.Printf("hole buffer:\n%v\n", b.String())
}
