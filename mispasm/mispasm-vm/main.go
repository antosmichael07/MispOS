package main

import (
	"fmt"
	"os"
)

var should_close = false

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
		fmt.Printf("Error reading executable file\n")
		panic(err)
	}
	return data
}
