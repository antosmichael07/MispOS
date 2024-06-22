package main

import (
	"fmt"
	"os"
)

var registers = make(map[string]any)

func main() {
	data := read_exec(os.Args[1])
	global, funcs := get_functions(data)
	init_instructions(&funcs)
	init_calls()

	run_function(funcs[global])
}

func read_exec(loc string) (data []byte) {
	data, err := os.ReadFile(loc)
	if err != nil {
		fmt.Printf("Error reading executable\n")
		panic(err)
	}
	return data
}
