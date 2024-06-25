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
	instructions[ADD] = get_math_operation_func(0)
	instructions[SUB] = get_math_operation_func(1)
	instructions[FCAL] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != STRING {
			panic("Invalid argument type for CAL\n")
		} else {
			run_function((*funcs)[string(arg1[1:len(arg1)-1])])
		}
	}
	instructions[MUL] = get_math_operation_func(2)
	instructions[DIV] = get_math_operation_func(3)
	instructions[CALL] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != INT8 {
			panic("Invalid argument type for CALL\n")
		} else {
			calls[arg1[1]]([]byte{}, []byte{})
		}
	}
	instructions[PUSH] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != REG {
			panic("Invalid argument type for PUSH\n")
		} else {
			stack_push(arg1[1], int(arg1[2]))
		}
	}
	instructions[POP] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != REG {
			panic("Invalid argument type for POP\n")
		} else {
			stack_pop(arg1[1], int(arg1[2]))
		}
	}
	instructions[MOV] = func(arg1 []byte, arg2 []byte) {
		register_set(arg1[1], int(arg1[2]), ConvertToValue(arg2))
	}
	instructions[MOD] = get_math_operation_func(4)
	instructions[CMP] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != arg2[0] && arg1[0] != REG && arg2[0] != REG {
			panic("Invalid argument type for CMP\n")
		}

		REGISTER_CMP[0] = arg1
		REGISTER_CMP[1] = arg2
	}
}
