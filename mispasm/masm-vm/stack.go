package main

type Stack struct {
	Name string
	Data any
}

var stack = []Stack{}

func stack_push(reg string, data any) {
	stack = append(stack, Stack{reg, data})
}

func stack_pop(reg string) {
	for i, v := range stack {
		if v.Name == reg {
			stack = append(stack[:i], stack[i+1:]...)
			break
		}
	}
}
