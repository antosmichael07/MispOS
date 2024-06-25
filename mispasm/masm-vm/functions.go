package main

const (
	GLOBAL byte = iota
	FUNC
)

type Function struct {
	Instructions []byte
	Labels       []int
}

func get_functions(data []byte) (global string, funcs map[string]Function) {
	if data[0] == GLOBAL {
		name := ""
		for i := 1; data[i] != 0; i++ {
			name += string(data[i])
		}
		global = name
	} else {
		panic("Error, no global function\n")
	}
	funcs = make(map[string]Function)
	for i := 0; i < len(data); i++ {
		if data[i] == FUNC {
			name := ""
			for i += 1; data[i] != 0; i++ {
				name += string(data[i])
			}

			instructions := []byte{}
			for i += 1; data[i] != 255; {
				_, _, _, _, arg_size := get_args(data, i)
				for j := i; j <= i+arg_size; j++ {
					instructions = append(instructions, data[j])
				}
				i += arg_size + 1
			}

			labels := []int{}
			for j := 0; j < len(instructions); {
				_, _, _, _, arg_size := get_args(instructions, j)
				if instructions[j] == LABEL {
					labels = append(labels, j)
				}
				j += arg_size + 1
			}

			funcs[name] = Function{
				Instructions: instructions,
				Labels:       labels,
			}
		}
	}

	is_global_valid := false
	for i := range funcs {
		if i == global {
			is_global_valid = true
			break
		}
	}
	if !is_global_valid {
		panic("Error, global isn't a function\n")
	}

	return global, funcs
}

func run_function(function Function) {
	for i := 0; i < len(function.Instructions); i++ {
		arg1, arg2, is_arg1, is_arg2, arg_size := get_args(function.Instructions, i)
		if function.Instructions[i] == LABEL {
			continue
		}
		if function.Instructions[i] == JMP {
			if arg1[0] != INT8 {
				panic("Invalid argument for JMP\n")
			}
			i = function.Labels[int(arg1[1])]
			continue
		}
		if function.Instructions[i] == JE || function.Instructions[i] == JNE || function.Instructions[i] == JG || function.Instructions[i] == JGE || function.Instructions[i] == JL || function.Instructions[i] == JLE {
			compare(&function, &i)(arg1, arg2)
			continue
		}
		if is_arg1 && is_arg2 {
			instructions[function.Instructions[i]](arg1, arg2)
		} else if is_arg1 {
			instructions[function.Instructions[i]](arg1, []byte{})
		} else {
			instructions[function.Instructions[i]]([]byte{}, []byte{})
		}
		i += arg_size
	}
}

func get_args(f_instructions []byte, i int) (arg1 []byte, arg2 []byte, is_arg1 bool, is_arg2 bool, arg_size int) {
	offset := 0
	if arg_sizes[f_instructions[i]] >= 1 {
		is_arg1 = true
		if f_instructions[i+1] == STRING {
			for j := i + 1; j < len(f_instructions); j++ {
				if f_instructions[j] == 0 {
					offset = j + 1
					arg_size = j - i
					arg1 = f_instructions[i+1 : j+1]
					break
				}
			}
		} else if f_instructions[i+1] <= byte(len(type_sizes)-1) {
			offset = i + int(type_sizes[f_instructions[i+1]]) + 2
			arg_size = int(type_sizes[f_instructions[i+1]] + 1)
			arg1 = f_instructions[i+1 : offset]
		} else {
			panic("Invalid argument type\n")
		}
	}

	if arg_sizes[f_instructions[i]] >= 2 {
		is_arg2 = true
		if f_instructions[offset] == STRING {
			for j := offset; j < len(f_instructions); j++ {
				if f_instructions[j] == 0 {
					arg_size = j - i
					arg2 = f_instructions[offset : j+1]
					break
				}
			}
		} else if f_instructions[offset] <= byte(len(type_sizes)-1) {
			arg_size = offset + int(type_sizes[f_instructions[offset]]) - i
			arg2 = f_instructions[offset : offset+int(type_sizes[f_instructions[offset]])+1]
		} else {
			panic("Invalid argument type\n")
		}
	}

	return arg1, arg2, is_arg1, is_arg2, arg_size
}
