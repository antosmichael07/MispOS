package main

type Stack struct {
	Name  byte
	Index int
	Data  any
}

var stack = []Stack{}

func stack_push(reg byte, index int) {
	stack = append(stack, Stack{reg, index, register_get(reg, index)})
}

func stack_pop(reg byte, index int) {
	for i, v := range stack {
		if v.Name == reg && v.Index == index {
			stack = append(stack[:i], stack[i+1:]...)
			break
		}
	}
}
