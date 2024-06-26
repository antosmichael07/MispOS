package main

var add_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[RBI](0, val1.(int8)+val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[RSI](0, val1.(int16)+val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[RLI](0, val1.(int32)+val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[RLLI](0, val1.(int64)+val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[RBUI](0, val1.(uint8)+val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[RSUI](0, val1.(uint16)+val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[RLUI](0, val1.(uint32)+val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[RLLUI](0, val1.(uint64)+val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[RLF](0, val1.(float32)+val2.(float32))
	},
	func(val1 any, val2 any) {
		register_set[RLLF](0, val1.(float64)+val2.(float64))
	},
}
var add_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBI](0, byte_to_int8(arg1[1])+byte_to_int8(arg2[1]))
		} else {
			register_set[RBI](0, byte_to_int8(arg1[1])+register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSI](0, bytes_to_int16(arg1[1:])+bytes_to_int16(arg2[1:]))
		} else {
			register_set[RSI](0, bytes_to_int16(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLI](0, bytes_to_int32(arg1[1:])+bytes_to_int32(arg2[1:]))
		} else {
			register_set[RLI](0, bytes_to_int32(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])+bytes_to_int64(arg2[1:]))
		} else {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBUI](0, byte_to_uint8(arg1[1])+byte_to_uint8(arg2[1]))
		} else {
			register_set[RBUI](0, byte_to_uint8(arg1[1])+register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])+bytes_to_uint16(arg2[1:]))
		} else {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])+bytes_to_uint32(arg2[1:]))
		} else {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])+bytes_to_uint64(arg2[1:]))
		} else {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLF](0, bytes_to_float32(arg1[1:])+bytes_to_float32(arg2[1:]))
		} else {
			register_set[RLF](0, bytes_to_float32(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(float32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLF](0, bytes_to_float64(arg1[1:])+bytes_to_float64(arg2[1:]))
		} else {
			register_set[RLLF](0, bytes_to_float64(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(float64))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		add_math_operation_reg[arg1[1]%10](convert_to_value(arg1), convert_to_value(arg2))
	},
}

var sub_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[RBI](0, val1.(int8)-val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[RSI](0, val1.(int16)-val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[RLI](0, val1.(int32)-val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[RLLI](0, val1.(int64)-val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[RBUI](0, val1.(uint8)-val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[RSUI](0, val1.(uint16)-val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[RLUI](0, val1.(uint32)-val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[RLLUI](0, val1.(uint64)-val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[RLF](0, val1.(float32)-val2.(float32))
	},
	func(val1 any, val2 any) {
		register_set[RLLF](0, val1.(float64)-val2.(float64))
	},
}

var sub_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBI](0, byte_to_int8(arg1[1])-byte_to_int8(arg2[1]))
		} else {
			register_set[RBI](0, byte_to_int8(arg1[1])-register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSI](0, bytes_to_int16(arg1[1:])-bytes_to_int16(arg2[1:]))
		} else {
			register_set[RSI](0, bytes_to_int16(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLI](0, bytes_to_int32(arg1[1:])-bytes_to_int32(arg2[1:]))
		} else {
			register_set[RLI](0, bytes_to_int32(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])-bytes_to_int64(arg2[1:]))
		} else {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBUI](0, byte_to_uint8(arg1[1])-byte_to_uint8(arg2[1]))
		} else {
			register_set[RBUI](0, byte_to_uint8(arg1[1])-register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])-bytes_to_uint16(arg2[1:]))
		} else {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])-bytes_to_uint32(arg2[1:]))
		} else {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])-bytes_to_uint64(arg2[1:]))
		} else {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLF](0, bytes_to_float32(arg1[1:])-bytes_to_float32(arg2[1:]))
		} else {
			register_set[RLF](0, bytes_to_float32(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(float32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLF](0, bytes_to_float64(arg1[1:])-bytes_to_float64(arg2[1:]))
		} else {
			register_set[RLLF](0, bytes_to_float64(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(float64))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		sub_math_operation_reg[arg1[1]%10](convert_to_value(arg1), convert_to_value(arg2))
	},
}

var mul_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[RBI](0, val1.(int8)*val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[RSI](0, val1.(int16)*val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[RLI](0, val1.(int32)*val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[RLLI](0, val1.(int64)*val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[RBUI](0, val1.(uint8)*val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[RSUI](0, val1.(uint16)*val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[RLUI](0, val1.(uint32)*val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[RLLUI](0, val1.(uint64)*val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[RLF](0, val1.(float32)*val2.(float32))
	},
	func(val1 any, val2 any) {
		register_set[RLLF](0, val1.(float64)*val2.(float64))
	},
}

var mul_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBI](0, byte_to_int8(arg1[1])*byte_to_int8(arg2[1]))
		} else {
			register_set[RBI](0, byte_to_int8(arg1[1])*register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSI](0, bytes_to_int16(arg1[1:])*bytes_to_int16(arg2[1:]))
		} else {
			register_set[RSI](0, bytes_to_int16(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLI](0, bytes_to_int32(arg1[1:])*bytes_to_int32(arg2[1:]))
		} else {
			register_set[RLI](0, bytes_to_int32(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])*bytes_to_int64(arg2[1:]))
		} else {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBUI](0, byte_to_uint8(arg1[1])*byte_to_uint8(arg2[1]))
		} else {
			register_set[RBUI](0, byte_to_uint8(arg1[1])*register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])*bytes_to_uint16(arg2[1:]))
		} else {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])*bytes_to_uint32(arg2[1:]))
		} else {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])*bytes_to_uint64(arg2[1:]))
		} else {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLF](0, bytes_to_float32(arg1[1:])*bytes_to_float32(arg2[1:]))
		} else {
			register_set[RLF](0, bytes_to_float32(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(float32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLF](0, bytes_to_float64(arg1[1:])*bytes_to_float64(arg2[1:]))
		} else {
			register_set[RLLF](0, bytes_to_float64(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(float64))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		mul_math_operation_reg[arg1[1]%10](convert_to_value(arg1), convert_to_value(arg2))
	},
}

var div_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[RBI](0, val1.(int8)/val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[RSI](0, val1.(int16)/val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[RLI](0, val1.(int32)/val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[RLLI](0, val1.(int64)/val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[RBUI](0, val1.(uint8)/val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[RSUI](0, val1.(uint16)/val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[RLUI](0, val1.(uint32)/val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[RLLUI](0, val1.(uint64)/val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[RLF](0, val1.(float32)/val2.(float32))
	},
	func(val1 any, val2 any) {
		register_set[RLLF](0, val1.(float64)/val2.(float64))
	},
}

var div_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBI](0, byte_to_int8(arg1[1])/byte_to_int8(arg2[1]))
		} else {
			register_set[RBI](0, byte_to_int8(arg1[1])/register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSI](0, bytes_to_int16(arg1[1:])/bytes_to_int16(arg2[1:]))
		} else {
			register_set[RSI](0, bytes_to_int16(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLI](0, bytes_to_int32(arg1[1:])/bytes_to_int32(arg2[1:]))
		} else {
			register_set[RLI](0, bytes_to_int32(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])/bytes_to_int64(arg2[1:]))
		} else {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBUI](0, byte_to_uint8(arg1[1])/byte_to_uint8(arg2[1]))
		} else {
			register_set[RBUI](0, byte_to_uint8(arg1[1])/register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])/bytes_to_uint16(arg2[1:]))
		} else {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])/bytes_to_uint32(arg2[1:]))
		} else {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])/bytes_to_uint64(arg2[1:]))
		} else {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLF](0, bytes_to_float32(arg1[1:])/bytes_to_float32(arg2[1:]))
		} else {
			register_set[RLF](0, bytes_to_float32(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(float32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLF](0, bytes_to_float64(arg1[1:])/bytes_to_float64(arg2[1:]))
		} else {
			register_set[RLLF](0, bytes_to_float64(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(float64))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		div_math_operation_reg[arg1[1]%10](convert_to_value(arg1), convert_to_value(arg2))
	},
}

var mod_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[RBI](0, val1.(int8)%val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[RSI](0, val1.(int16)%val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[RLI](0, val1.(int32)%val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[RLLI](0, val1.(int64)%val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[RBUI](0, val1.(uint8)%val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[RSUI](0, val1.(uint16)%val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[RLUI](0, val1.(uint32)%val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[RLLUI](0, val1.(uint64)%val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[RLF](0, float32(int(val1.(float32))%int(val2.(float32))))
	},
	func(val1 any, val2 any) {
		register_set[RLLF](0, float64(int(val1.(float64))%int(val2.(float64))))
	},
}

var mod_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBI](0, byte_to_int8(arg1[1])%byte_to_int8(arg2[1]))
		} else {
			register_set[RBI](0, byte_to_int8(arg1[1])%register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSI](0, bytes_to_int16(arg1[1:])%bytes_to_int16(arg2[1:]))
		} else {
			register_set[RSI](0, bytes_to_int16(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLI](0, bytes_to_int32(arg1[1:])%bytes_to_int32(arg2[1:]))
		} else {
			register_set[RLI](0, bytes_to_int32(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])%bytes_to_int64(arg2[1:]))
		} else {
			register_set[RLLI](0, bytes_to_int64(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RBUI](0, byte_to_uint8(arg1[1])%byte_to_uint8(arg2[1]))
		} else {
			register_set[RBUI](0, byte_to_uint8(arg1[1])%register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])%bytes_to_uint16(arg2[1:]))
		} else {
			register_set[RSUI](0, bytes_to_uint16(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])%bytes_to_uint32(arg2[1:]))
		} else {
			register_set[RLUI](0, bytes_to_uint32(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])%bytes_to_uint64(arg2[1:]))
		} else {
			register_set[RLLUI](0, bytes_to_uint64(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLF](0, float32(int(bytes_to_float32(arg1[1:]))%int(bytes_to_float32(arg2[1:]))))
		} else {
			register_set[RLF](0, float32(int(bytes_to_float32(arg1[1:]))%int(register_get[arg2[1]](int(arg2[2])).(float32))))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != REG {
			register_set[RLLF](0, float64(int(bytes_to_float64(arg1[1:]))%int(bytes_to_float64(arg2[1:]))))
		} else {
			register_set[RLLF](0, float64(int(bytes_to_float64(arg1[1:]))%int(register_get[arg2[1]](int(arg2[2])).(float64))))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		mod_math_operation_reg[arg1[1]%10](convert_to_value(arg1), convert_to_value(arg2))
	},
}
