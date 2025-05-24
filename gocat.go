package main

import (
	"fmt"
	"log"
	"os"
)

func logFatalIfErr(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func getFilenameFromArgs() string {
	if len(os.Args) < 2 || os.Args[1] == "-" {
		return "/dev/stdin"
	}
	return os.Args[1]
}

func main() {
	filename := getFilenameFromArgs()
	data, err := os.ReadFile(filename)
	logFatalIfErr("Error reading file:", err)
	fmt.Println(string(data))
}
