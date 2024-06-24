package main

var REGISTER_CMP = [2][]byte{}

func compare(function *Function, index *int) func([]byte, []byte) {
	return func(arg1 []byte, arg2 []byte) {
		if function.Instructions[*index] == byte(JE) {
			if ConvertToValue(REGISTER_CMP[0]) == ConvertToValue(REGISTER_CMP[1]) {
				*index = function.Labels[int(arg1[1])]
			} else {
				*index = function.Labels[int(arg2[1])]
			}
		}
	}
}
