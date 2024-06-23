package main

const (
	ADD = iota + 1
	SUB
	FCAL
	MUL
	DIV
	CALL
	PUSH
	POP
	MOV
)

var arg_sizes = []byte{0, 2, 2, 1, 2, 2, 1, 1, 1, 2}

var instructions = make(map[byte]func([]byte, []byte))

func init_instructions(funcs *map[string][]byte) {
	instructions[ADD] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] < byte(INT8) || arg2[0] < byte(INT8) || arg1[0] > byte(FLOAT64) || arg2[0] > byte(FLOAT64) {
			panic("Invalid arguments for ADD\n")
		} else if arg1[0] != arg2[0] {
			panic("Arguments for ADD are different types\n")
		} else {
			switch arg1[0] {
			case byte(INT8):
				register_set(RBI, 0, ByteToInt8(arg1[1])+ByteToInt8(arg2[1]))
			case byte(INT16):
				register_set(RSI, 0, BytesToInt16(arg1[1:])+BytesToInt16(arg2[1:]))
			case byte(INT32):
				register_set(RLI, 0, BytesToInt32(arg1[1:])+BytesToInt32(arg2[1:]))
			case byte(INT64):
				register_set(RLLI, 0, BytesToInt64(arg1[1:])+BytesToInt64(arg2[1:]))
			case byte(UINT8):
				register_set(RBUI, 0, ByteToUint8(arg1[1])+ByteToUint8(arg2[1]))
			case byte(UINT16):
				register_set(RSUI, 0, BytesToUint16(arg1[1:])+BytesToUint16(arg2[1:]))
			case byte(UINT32):
				register_set(RLUI, 0, BytesToUint32(arg1[1:])+BytesToUint32(arg2[1:]))
			case byte(UINT64):
				register_set(RLLUI, 0, BytesToUint64(arg1[1:])+BytesToUint64(arg2[1:]))
			case byte(FLOAT32):
				register_set(RLF, 0, BytesToFloat32(arg1[1:])+BytesToFloat32(arg2[1:]))
			case byte(FLOAT64):
				register_set(RLLF, 0, BytesToFloat64(arg1[1:])+BytesToFloat64(arg2[1:]))
			}
		}
	}
	instructions[SUB] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] < byte(INT8) || arg2[0] < byte(INT8) || arg1[0] > byte(FLOAT64) || arg2[0] > byte(FLOAT64) {
			panic("Invalid arguments for SUB\n")
		} else if arg1[0] != arg2[0] {
			panic("Arguments for SUB are different types\n")
		} else {
			switch arg1[0] {
			case byte(INT8):
				register_set(RBI, 0, ByteToInt8(arg1[1])-ByteToInt8(arg2[1]))
			case byte(INT16):
				register_set(RSI, 0, BytesToInt16(arg1[1:])-BytesToInt16(arg2[1:]))
			case byte(INT32):
				register_set(RLI, 0, BytesToInt32(arg1[1:])-BytesToInt32(arg2[1:]))
			case byte(INT64):
				register_set(RLLI, 0, BytesToInt64(arg1[1:])-BytesToInt64(arg2[1:]))
			case byte(UINT8):
				register_set(RBUI, 0, ByteToUint8(arg1[1])-ByteToUint8(arg2[1]))
			case byte(UINT16):
				register_set(RSUI, 0, BytesToUint16(arg1[1:])-BytesToUint16(arg2[1:]))
			case byte(UINT32):
				register_set(RLUI, 0, BytesToUint32(arg1[1:])-BytesToUint32(arg2[1:]))
			case byte(UINT64):
				register_set(RLLUI, 0, BytesToUint64(arg1[1:])-BytesToUint64(arg2[1:]))
			case byte(FLOAT32):
				register_set(RLF, 0, BytesToFloat32(arg1[1:])-BytesToFloat32(arg2[1:]))
			case byte(FLOAT64):
				register_set(RLLF, 0, BytesToFloat64(arg1[1:])-BytesToFloat64(arg2[1:]))
			}
		}
	}
	instructions[FCAL] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != byte(STRING) {
			panic("Invalid argument type for CAL\n")
		} else {
			run_function((*funcs)[string(arg1[1:len(arg1)-1])])
		}
	}
	instructions[MUL] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] < byte(INT8) || arg2[0] < byte(INT8) || arg1[0] > byte(FLOAT64) || arg2[0] > byte(FLOAT64) {
			panic("Invalid arguments for MUL\n")
		} else if arg1[0] != arg2[0] {
			panic("Arguments for MUL are different types\n")
		} else {
			switch arg1[0] {
			case byte(INT8):
				register_set(RBI, 0, ByteToInt8(arg1[1])*ByteToInt8(arg2[1]))
			case byte(INT16):
				register_set(RSI, 0, BytesToInt16(arg1[1:])*BytesToInt16(arg2[1:]))
			case byte(INT32):
				register_set(RLI, 0, BytesToInt32(arg1[1:])*BytesToInt32(arg2[1:]))
			case byte(INT64):
				register_set(RLLI, 0, BytesToInt64(arg1[1:])*BytesToInt64(arg2[1:]))
			case byte(UINT8):
				register_set(RBUI, 0, ByteToUint8(arg1[1])*ByteToUint8(arg2[1]))
			case byte(UINT16):
				register_set(RSUI, 0, BytesToUint16(arg1[1:])*BytesToUint16(arg2[1:]))
			case byte(UINT32):
				register_set(RLUI, 0, BytesToUint32(arg1[1:])*BytesToUint32(arg2[1:]))
			case byte(UINT64):
				register_set(RLLUI, 0, BytesToUint64(arg1[1:])*BytesToUint64(arg2[1:]))
			case byte(FLOAT32):
				register_set(RLF, 0, BytesToFloat32(arg1[1:])*BytesToFloat32(arg2[1:]))
			case byte(FLOAT64):
				register_set(RLLF, 0, BytesToFloat64(arg1[1:])*BytesToFloat64(arg2[1:]))
			}
		}
	}
	instructions[DIV] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] < byte(INT8) || arg2[0] < byte(INT8) || arg1[0] > byte(FLOAT64) || arg2[0] > byte(FLOAT64) {
			panic("Invalid arguments for DIV\n")
		} else if arg1[0] != arg2[0] {
			panic("Arguments for DIV are different types\n")
		} else {
			switch arg1[0] {
			case byte(INT8):
				register_set(RBI, 0, ByteToInt8(arg1[1])/ByteToInt8(arg2[1]))
			case byte(INT16):
				register_set(RSI, 0, BytesToInt16(arg1[1:])/BytesToInt16(arg2[1:]))
			case byte(INT32):
				register_set(RLI, 0, BytesToInt32(arg1[1:])/BytesToInt32(arg2[1:]))
			case byte(INT64):
				register_set(RLLI, 0, BytesToInt64(arg1[1:])/BytesToInt64(arg2[1:]))
			case byte(UINT8):
				register_set(RBUI, 0, ByteToUint8(arg1[1])/ByteToUint8(arg2[1]))
			case byte(UINT16):
				register_set(RSUI, 0, BytesToUint16(arg1[1:])/BytesToUint16(arg2[1:]))
			case byte(UINT32):
				register_set(RLUI, 0, BytesToUint32(arg1[1:])/BytesToUint32(arg2[1:]))
			case byte(UINT64):
				register_set(RLLUI, 0, BytesToUint64(arg1[1:])/BytesToUint64(arg2[1:]))
			case byte(FLOAT32):
				register_set(RLF, 0, BytesToFloat32(arg1[1:])/BytesToFloat32(arg2[1:]))
			case byte(FLOAT64):
				register_set(RLLF, 0, BytesToFloat64(arg1[1:])/BytesToFloat64(arg2[1:]))
			}
		}
	}
	instructions[CALL] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != byte(INT8) {
			panic("Invalid argument type for CALL\n")
		} else {
			calls[int(arg1[1])]([]byte{}, []byte{})
		}
	}
	instructions[PUSH] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != byte(INT16) {
			panic("Invalid argument type for PUSH\n")
		} else {
			stack_push(arg1[1], int(arg1[2]))
		}
	}
	instructions[POP] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != byte(INT16) {
			panic("Invalid argument type for POP\n")
		} else {
			stack_pop(arg1[1], int(arg1[2]))
		}
	}
	instructions[MOV] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != byte(INT16) {
			panic("Invalid arguments for MOV\n")
		} else {
			register_set(arg1[1], int(arg1[2]), ConvertToValue(arg2))
		}
	}
}
