package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Println("Usage: os-lab-4 <file>")
	os.Exit(1)
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	file, err := os.Open(os.Args[1])
	handle(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_ = ParseCommand(line)
	}
}
