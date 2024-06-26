package main

const (
	add byte = iota
	sub
	fcal
	mul
	div
	call
	push
	pop
	mov
	label
	jmp
	mod
	cmp
	je
	jne
	jg
	jge
	jl
	jle
)

var arg_sizes = []byte{2, 2, 1, 2, 2, 1, 1, 1, 2, 0, 1, 2, 2, 2, 2, 2, 2, 2, 2}

var instructions = [19]func([]byte, []byte){}

func init_instructions(funcs *map[string]Function) {
	instructions[add] = func(arg1 []byte, arg2 []byte) {
		add_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[sub] = func(arg1 []byte, arg2 []byte) {
		sub_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[fcal] = func(arg1 []byte, arg2 []byte) {
		run_function((*funcs)[string(arg1[1:len(arg1)-1])])
	}
	instructions[mul] = func(arg1 []byte, arg2 []byte) {
		mul_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[div] = func(arg1 []byte, arg2 []byte) {
		div_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[call] = func(arg1 []byte, arg2 []byte) {
		calls[arg1[1]]([]byte{}, []byte{})
	}
	instructions[push] = func(arg1 []byte, arg2 []byte) {
		stack_push(arg1[1], int(arg1[2]))
	}
	instructions[pop] = func(arg1 []byte, arg2 []byte) {
		stack_pop(arg1[1], int(arg1[2]))
	}
	instructions[mov] = func(arg1 []byte, arg2 []byte) {
		register_set[arg1[1]](arg1[2], convert_to_value(arg2))
	}
	instructions[mod] = func(arg1 []byte, arg2 []byte) {
		mod_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[cmp] = func(arg1 []byte, arg2 []byte) {
		register_cmp[0] = arg1
		register_cmp[1] = arg2
	}
}
