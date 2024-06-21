package main

type Function struct {
	Instructions []byte
	Caller       int
}

const (
	GLOBAL = iota
	FUNC
)

func get_functions(data []byte) (global string, funcs map[string]Function) {
	funcs = make(map[string]Function)
	for i := 0; i < len(data); i++ {
		if data[i] == GLOBAL {
			if global != "" {
				panic("Error, two globals initialized\n")
			}
			name := ""
			offset := i + 1
			for ; data[offset] != 0; offset++ {
				name += string(data[offset])
			}
			global = name
			i = offset
		}
		if data[i] == FUNC {
			name := ""
			for i += 1; data[i] != 0; i++ {
				name += string(data[i])
			}

			instructions := []byte{}
			for i += 1; data[i] != 0; i++ {
				instructions = append(instructions, data[i])
			}

			funcs[name] = Function{Instructions: instructions, Caller: -1}
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
		if function.Instructions[i] >= 1 {
			instructions[function.Instructions[i]](function.Instructions[i+1], function.Instructions[i+2])
			i += 2
		}
	}
}
