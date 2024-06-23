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
	msg := ""
	continu := false
	for _, v := range stack {
		if !continu {
			msg = v.Data.(string)
			continu = true
			continue
		}
		msg = fmt.Sprintf(msg, v.Data)
	}
	fmt.Print(msg)
}
