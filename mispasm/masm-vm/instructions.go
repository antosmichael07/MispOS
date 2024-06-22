package main

const (
	ADD = iota + 1
	SUB
	CAL
)

var arg_sizes = []byte{0, 2, 2, 1}

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
				registers["rbi0"] = ByteToInt8(arg1[1]) + ByteToInt8(arg2[1])
			case byte(INT16):
				registers["rsi0"] = BytesToInt16(arg1[1:]) + BytesToInt16(arg2[1:])
			case byte(INT32):
				registers["rli0"] = BytesToInt32(arg1[1:]) + BytesToInt32(arg2[1:])
			case byte(INT64):
				registers["rlli0"] = BytesToInt64(arg1[1:]) + BytesToInt64(arg2[1:])
			case byte(UINT8):
				registers["rbui0"] = ByteToUint8(arg1[1]) + ByteToUint8(arg2[1])
			case byte(UINT16):
				registers["rsui0"] = BytesToUint16(arg1[1:]) + BytesToUint16(arg2[1:])
			case byte(UINT32):
				registers["rlui0"] = BytesToUint32(arg1[1:]) + BytesToUint32(arg2[1:])
			case byte(UINT64):
				registers["rllui0"] = BytesToUint64(arg1[1:]) + BytesToUint64(arg2[1:])
			case byte(FLOAT32):
				registers["rlf0"] = BytesToFloat32(arg1[1:]) + BytesToFloat32(arg2[1:])
			case byte(FLOAT64):
				registers["rllf0"] = BytesToFloat64(arg1[1:]) + BytesToFloat64(arg2[1:])
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
				registers["rbi0"] = byte(ByteToInt8(arg1[1]) - ByteToInt8(arg2[1]))
			case byte(INT16):
				registers["rsi0"] = byte(BytesToInt16(arg1[1:]) - BytesToInt16(arg2[1:]))
			case byte(INT32):
				registers["rli0"] = byte(BytesToInt32(arg1[1:]) - BytesToInt32(arg2[1:]))
			case byte(INT64):
				registers["rlli0"] = byte(BytesToInt64(arg1[1:]) - BytesToInt64(arg2[1:]))
			case byte(UINT8):
				registers["rbui0"] = byte(ByteToUint8(arg1[1]) - ByteToUint8(arg2[1]))
			case byte(UINT16):
				registers["rsui0"] = byte(BytesToUint16(arg1[1:]) - BytesToUint16(arg2[1:]))
			case byte(UINT32):
				registers["rlui0"] = byte(BytesToUint32(arg1[1:]) - BytesToUint32(arg2[1:]))
			case byte(UINT64):
				registers["rllui0"] = byte(BytesToUint64(arg1[1:]) - BytesToUint64(arg2[1:]))
			case byte(FLOAT32):
				registers["rlf0"] = byte(BytesToFloat32(arg1[1:]) - BytesToFloat32(arg2[1:]))
			case byte(FLOAT64):
				registers["rllf0"] = byte(BytesToFloat64(arg1[1:]) - BytesToFloat64(arg2[1:]))
			}
		}
	}
	instructions[CAL] = func(arg1 []byte, arg2 []byte) {
		if arg1[0] != 12 {
			panic("Invalid argument type for CAL\n")
		} else {
			run_function((*funcs)[string(arg1[1:len(arg1)-1])])
		}
	}
}
