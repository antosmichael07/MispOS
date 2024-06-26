package main

import "fmt"

const (
	printf byte = iota
)

var calls = [1]func([]byte, []byte){}

func init_calls() {
	calls[printf] = func(arg1 []byte, arg2 []byte) {
		print_stack()
	}
}

func print_stack() {
	msg := ""
	continu := false
	for _, v := range stack {
		if !continu {
			msg = v.data.(string)
			continu = true
			continue
		}
		msg = fmt.Sprintf(msg, v.data)
	}
	fmt.Print(msg)
}
