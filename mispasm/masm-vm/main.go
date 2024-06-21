package main

import (
	"fmt"
	"os"
)

const (
	GLOBAL = iota
	FUNC
)

func main() {
	data := read_exec(os.Args[1])
	global := ""
	funcs := map[string]func(){}

	for i := 0; i < len(data); i++ {
		if data[i] == GLOBAL {
			if global != "" {
				fmt.Printf("Error, two globals initialized\n")
				return
			}
			name := ""
			offset := i + 1
			for ; data[offset] != 0; offset++ {
				name += string(data[offset])
			}
			global = name
			i = offset
		}
		if data[i] == FUNC {
			name := ""
			offset := i + 1
			for ; data[offset] != 0; offset++ {
				name += string(data[offset])
			}
			funcs[name] = func() {}
			i = offset
		}
	}

	is_global_valid := false
	for i := range funcs {
		if i == global {
			is_global_valid = true
			break
		}
	}
	if !is_global_valid {
		fmt.Printf("Error, global isn't a function\n")
		return
	}

	fmt.Printf("Global: %s\n", global)

	for i := range funcs {
		fmt.Printf("Func: %s\n", i)
	}
}

func read_exec(loc string) (data []byte) {
	data, err := os.ReadFile(loc)
	if err != nil {
		fmt.Printf("Error reading executable\n")
		panic(err)
	}
	return data
}
