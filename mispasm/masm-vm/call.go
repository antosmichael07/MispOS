package main

import "fmt"

const (
	PRINTF = iota + 1
)

var calls = make(map[int]func([]byte, []byte))

func init_calls() {
	calls[PRINTF] = func(arg1 []byte, arg2 []byte) {
		print_stack()
	}
}

func print_stack() {
	for _, v := range stack {
		fmt.Printf("%v\n", v)
	}
}
