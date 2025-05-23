package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-" {
		input, err := os.ReadFile("/dev/stdin")
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}
		fmt.Print(string(input))
		return
	}
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))
}
