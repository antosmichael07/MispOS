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

	bootable_disks := GetBootableDisks(disks)

	for _, disk := range disks {
		fmt.Printf("Disk: %s\n", disk.Name())
	}

	fmt.Printf("\n")

	for _, disk := range bootable_disks {
		fmt.Printf("Bootable disk: %s\n", disk.Name())
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

	for _, disk := range disk_names {
		file, err := os.Open(fmt.Sprintf("%s/%s", dir, disk))

		if err != nil {
			fmt.Printf("Error opening disk \"%s\"\n", disk)
			continue
		}

		disks = append(disks, file)
	}

	return disks, nil
}

func GetBootableDisks(disks []*os.File) (bootable_disks []*os.File) {
	for _, disk := range disks {
		data := make([]byte, 1)
		_, err := disk.Read(data)

		if err != nil {
			fmt.Printf("Error reading disk \"%s\"\n", disk.Name())
			continue
		}

		if data[0] & 0b10000000 == 0b10000000 {
			bootable_disks = append(bootable_disks, disk)
		}
	}

	return bootable_disks
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

	data[0] &= 0b01111111

	for i, v := range version {
		if data[i] != v {
			return false
		}
	}

	return true
}
