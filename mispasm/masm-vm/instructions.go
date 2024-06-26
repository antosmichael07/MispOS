package main

const (
	ADD byte = iota
	SUB
	FCAL
	MUL
	DIV
	CALL
	PUSH
	POP
	MOV
	LABEL
	JMP
	MOD
	CMP
	JE
	JNE
	JG
	JGE
	JL
	JLE
)

var arg_sizes = []byte{2, 2, 1, 2, 2, 1, 1, 1, 2, 0, 1, 2, 2, 2, 2, 2, 2, 2, 2}

var instructions = [19]func([]byte, []byte){}

func init_instructions(funcs *map[string]Function) {
	instructions[ADD] = func(arg1 []byte, arg2 []byte) {
		add_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[SUB] = func(arg1 []byte, arg2 []byte) {
		sub_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[FCAL] = func(arg1 []byte, arg2 []byte) {
		run_function((*funcs)[string(arg1[1:len(arg1)-1])])
	}
	instructions[MUL] = func(arg1 []byte, arg2 []byte) {
		mul_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[DIV] = func(arg1 []byte, arg2 []byte) {
		div_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[CALL] = func(arg1 []byte, arg2 []byte) {
		calls[arg1[1]]([]byte{}, []byte{})
	}
	instructions[PUSH] = func(arg1 []byte, arg2 []byte) {
		stack_push(arg1[1], int(arg1[2]))
	}
	instructions[POP] = func(arg1 []byte, arg2 []byte) {
		stack_pop(arg1[1], int(arg1[2]))
	}
	instructions[MOV] = func(arg1 []byte, arg2 []byte) {
		register_set[arg1[1]](arg1[2], convert_to_value(arg2))
	}
	instructions[MOD] = func(arg1 []byte, arg2 []byte) {
		mod_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[CMP] = func(arg1 []byte, arg2 []byte) {
		REGISTER_CMP[0] = arg1
		REGISTER_CMP[1] = arg2
	}
}
