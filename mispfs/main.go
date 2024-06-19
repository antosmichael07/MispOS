package main

import (
	"fmt"
	"os"
)

var version = [3]byte{1, 0, 0}

func main() {
	if !CheckVersion(os.Args[1]) {
		fmt.Printf("Disk doesn't have the same version as your read-writer\n")
		return
	}

	fmt.Printf("Good disk you have here\n")
}

func CheckVersion(disk string) bool {
	data, err := os.ReadFile(disk)
	if err != nil {
		fmt.Printf("Error reading disk\n")
		panic(err)
	}

	if len(data) < 3 {
		return false
	}

	for i, v := range version {
		if data[i] != v {
			return false
		}
	}

	return true
}
