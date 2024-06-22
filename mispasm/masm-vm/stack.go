package main

var stack = make(map[string]any)

func stack_push(reg string, data any) {
	stack[reg] = data
}

func stack_pop(reg string) {
	delete(stack, reg)
}
