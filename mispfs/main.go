package main

import (
	"fmt"
	"os"
)

var version = [3]byte{1, 0, 0}

func main() {
	disks, err := LoadDisks("disks")
	if err != nil {
		panic(err)
	}

	for _, disk := range disks {
		fmt.Printf("Disk: %s\n", disk.Name())
	}
}

func LoadDisks(dir string) (disks []*os.File, err error) {
	dir_contents, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading disk directory\n")
		return nil, err
	}
	var disk_names []string
	for _, file := range dir_contents {
		file_name := file.Name()
		if file.IsDir() || file_name[len(file_name)-6:] == ".mpfs" || !CheckVersion(fmt.Sprintf("%s/%s", dir, file_name)) {
			continue
		}
		disk_names = append(disk_names, file_name)
	}

	var valid_disks []*os.File
	for _, disk := range disk_names {
		file, err := os.Open(fmt.Sprintf("%s/%s", dir, disk))
		if err != nil {
			fmt.Printf("Error opening disk \"%s\"\n", disk)
			continue
		}
		valid_disks = append(valid_disks, file)
	}

	return valid_disks, nil
}

func CheckVersion(disk string) bool {
	data, err := os.ReadFile(disk)
	if err != nil {
		fmt.Printf("Error reading disk \"%s\"\n", disk)
		return false
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
