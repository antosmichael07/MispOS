package main

const (
	ADD = iota + 1
	SUB
)

var instructions = make(map[byte]func(a, b byte))

func init_instructions() {
	instructions[ADD] = func(a, b byte) {
		registers["r0"] = a + b
	}
	instructions[SUB] = func(a, b byte) {
		registers["r0"] = a - b
	}
}
