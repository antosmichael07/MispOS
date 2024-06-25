package main

func get_math_operation_func(op byte) func([]byte, []byte) {
	switch op {
	case 0:
		return func(arg1 []byte, arg2 []byte) {
			if (arg1[0] < INT8 || arg2[0] < INT8 || arg1[0] > FLOAT64 || arg2[0] > FLOAT64) && (arg1[0] != REG && arg2[0] != REG) {
				panic("Invalid arguments for ADD\n")
			} else if arg1[0] != arg2[0] && arg1[0] != REG && arg2[0] != REG {
				panic("Arguments for ADD are different types\n")
			} else {
				switch arg1[0] {
				case INT8:
					if arg2[0] != REG {
						register_set(RBI, 0, ByteToInt8(arg1[1])+ByteToInt8(arg2[1]))
					} else if arg2[1] == BI {
						register_set(RBI, 0, ByteToInt8(arg1[1])+register_get(arg2[1], int(arg2[2])).(int8))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case INT16:
					if arg2[0] != REG {
						register_set(RSI, 0, BytesToInt16(arg1[1:])+BytesToInt16(arg2[1:]))
					} else if arg2[1] == SI {
						register_set(RSI, 0, BytesToInt16(arg1[1:])+register_get(arg2[1], int(arg2[2])).(int16))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case INT32:
					if arg2[0] != REG {
						register_set(RLI, 0, BytesToInt32(arg1[1:])+BytesToInt32(arg2[1:]))
					} else if arg2[1] == LI {
						register_set(RLI, 0, BytesToInt32(arg1[1:])+register_get(arg2[1], int(arg2[2])).(int32))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case INT64:
					if arg2[0] != REG {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])+BytesToInt64(arg2[1:]))
					} else if arg2[1] == LLI {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])+register_get(arg2[1], int(arg2[2])).(int64))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case UINT8:
					if arg2[0] != REG {
						register_set(RBUI, 0, ByteToUint8(arg1[1])+ByteToUint8(arg2[1]))
					} else if arg2[1] == BUI {
						register_set(RBUI, 0, ByteToUint8(arg1[1])+register_get(arg2[1], int(arg2[2])).(uint8))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case UINT16:
					if arg2[0] != REG {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])+BytesToUint16(arg2[1:]))
					} else if arg2[1] == SUI {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])+register_get(arg2[1], int(arg2[2])).(uint16))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case UINT32:
					if arg2[0] != REG {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])+BytesToUint32(arg2[1:]))
					} else if arg2[1] == LUI {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])+register_get(arg2[1], int(arg2[2])).(uint32))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case UINT64:
					if arg2[0] != REG {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])+BytesToUint64(arg2[1:]))
					} else if arg2[1] == LLUI {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])+register_get(arg2[1], int(arg2[2])).(uint64))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case FLOAT32:
					if arg2[0] != REG {
						register_set(RLF, 0, BytesToFloat32(arg1[1:])+BytesToFloat32(arg2[1:]))
					} else if arg2[1] == LF {
						register_set(RLF, 0, BytesToFloat32(arg1[1:])+register_get(arg2[1], int(arg2[2])).(float32))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case FLOAT64:
					if arg2[0] != REG {
						register_set(RLLF, 0, BytesToFloat64(arg1[1:])+BytesToFloat64(arg2[1:]))
					} else if arg2[1] == LLF {
						register_set(RLLF, 0, BytesToFloat64(arg1[1:])+register_get(arg2[1], int(arg2[2])).(float64))
					} else {
						panic("Invalid argument for ADD\n")
					}
				case REG:
					if arg1[1] == S || arg1[1] == RS {
						panic("Invalid argument for ADD\n")
					}

					val1 := ConvertToValue(arg1)
					val2 := ConvertToValue(arg2)

					switch arg1[1] {
					case BI, RBI:
						register_set(RBI, 0, val1.(int8)+val2.(int8))
					case SI, RSI:
						register_set(RSI, 0, val1.(int16)+val2.(int16))
					case LI, RLI:
						register_set(RLI, 0, val1.(int32)+val2.(int32))
					case LLI, RLLI:
						register_set(RLLI, 0, val1.(int64)+val2.(int64))
					case BUI, RBUI:
						register_set(RBUI, 0, val1.(uint8)+val2.(uint8))
					case SUI, RSUI:
						register_set(RSUI, 0, val1.(uint16)+val2.(uint16))
					case LUI, RLUI:
						register_set(RLUI, 0, val1.(uint32)+val2.(uint32))
					case LLUI, RLLUI:
						register_set(RLLUI, 0, val1.(uint64)+val2.(uint64))
					case LF, RLF:
						register_set(RLF, 0, val1.(float32)+val2.(float32))
					case LLF, RLLF:
						register_set(RLLF, 0, val1.(float64)+val2.(float64))
					}
				}
			}
		}
	case 1:
		return func(arg1 []byte, arg2 []byte) {
			if (arg1[0] < INT8 || arg2[0] < INT8 || arg1[0] > FLOAT64 || arg2[0] > FLOAT64) && (arg1[0] != REG && arg2[0] != REG) {
				panic("Invalid arguments for SUB\n")
			} else if arg1[0] != arg2[0] && arg1[0] != REG && arg2[0] != REG {
				panic("Arguments for SUB are different types\n")
			} else {
				switch arg1[0] {
				case INT8:
					if arg2[0] != REG {
						register_set(RBI, 0, ByteToInt8(arg1[1])-ByteToInt8(arg2[1]))
					} else if arg2[1] == BI {
						register_set(RBI, 0, ByteToInt8(arg1[1])-register_get(arg2[1], int(arg2[2])).(int8))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case INT16:
					if arg2[0] != REG {
						register_set(RSI, 0, BytesToInt16(arg1[1:])-BytesToInt16(arg2[1:]))
					} else if arg2[1] == SI {
						register_set(RSI, 0, BytesToInt16(arg1[1:])-register_get(arg2[1], int(arg2[2])).(int16))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case INT32:
					if arg2[0] != REG {
						register_set(RLI, 0, BytesToInt32(arg1[1:])-BytesToInt32(arg2[1:]))
					} else if arg2[1] == LI {
						register_set(RLI, 0, BytesToInt32(arg1[1:])-register_get(arg2[1], int(arg2[2])).(int32))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case INT64:
					if arg2[0] != REG {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])-BytesToInt64(arg2[1:]))
					} else if arg2[1] == LLI {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])-register_get(arg2[1], int(arg2[2])).(int64))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case UINT8:
					if arg2[0] != REG {
						register_set(RBUI, 0, ByteToUint8(arg1[1])-ByteToUint8(arg2[1]))
					} else if arg2[1] == BUI {
						register_set(RBUI, 0, ByteToUint8(arg1[1])-register_get(arg2[1], int(arg2[2])).(uint8))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case UINT16:
					if arg2[0] != REG {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])-BytesToUint16(arg2[1:]))
					} else if arg2[1] == SUI {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])-register_get(arg2[1], int(arg2[2])).(uint16))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case UINT32:
					if arg2[0] != REG {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])-BytesToUint32(arg2[1:]))
					} else if arg2[1] == LUI {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])-register_get(arg2[1], int(arg2[2])).(uint32))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case UINT64:
					if arg2[0] != REG {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])-BytesToUint64(arg2[1:]))
					} else if arg2[1] == LLUI {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])-register_get(arg2[1], int(arg2[2])).(uint64))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case FLOAT32:
					if arg2[0] != REG {
						register_set(RLF, 0, BytesToFloat32(arg1[1:])-BytesToFloat32(arg2[1:]))
					} else if arg2[1] == LF {
						register_set(RLF, 0, BytesToFloat32(arg1[1:])-register_get(arg2[1], int(arg2[2])).(float32))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case FLOAT64:
					if arg2[0] != REG {
						register_set(RLLF, 0, BytesToFloat64(arg1[1:])-BytesToFloat64(arg2[1:]))
					} else if arg2[1] == LLF {
						register_set(RLLF, 0, BytesToFloat64(arg1[1:])-register_get(arg2[1], int(arg2[2])).(float64))
					} else {
						panic("Invalid argument for SUB\n")
					}
				case REG:
					if arg1[1] == S || arg1[1] == RS {
						panic("Invalid argument for SUB\n")
					}

					val1 := ConvertToValue(arg1)
					val2 := ConvertToValue(arg2)

					switch arg1[1] {
					case BI, RBI:
						register_set(RBI, 0, val1.(int8)-val2.(int8))
					case SI, RSI:
						register_set(RSI, 0, val1.(int16)-val2.(int16))
					case LI, RLI:
						register_set(RLI, 0, val1.(int32)-val2.(int32))
					case LLI, RLLI:
						register_set(RLLI, 0, val1.(int64)-val2.(int64))
					case BUI, RBUI:
						register_set(RBUI, 0, val1.(uint8)-val2.(uint8))
					case SUI, RSUI:
						register_set(RSUI, 0, val1.(uint16)-val2.(uint16))
					case LUI, RLUI:
						register_set(RLUI, 0, val1.(uint32)-val2.(uint32))
					case LLUI, RLLUI:
						register_set(RLLUI, 0, val1.(uint64)-val2.(uint64))
					case LF, RLF:
						register_set(RLF, 0, val1.(float32)-val2.(float32))
					case LLF, RLLF:
						register_set(RLLF, 0, val1.(float64)-val2.(float64))
					}
				}
			}
		}
	case 2:
		return func(arg1 []byte, arg2 []byte) {
			if (arg1[0] < INT8 || arg2[0] < INT8 || arg1[0] > FLOAT64 || arg2[0] > FLOAT64) && (arg1[0] != REG && arg2[0] != REG) {
				panic("Invalid arguments for MUL\n")
			} else if arg1[0] != arg2[0] && arg1[0] != REG && arg2[0] != REG {
				panic("Arguments for MUL are different types\n")
			} else {
				switch arg1[0] {
				case INT8:
					if arg2[0] != REG {
						register_set(RBI, 0, ByteToInt8(arg1[1])*ByteToInt8(arg2[1]))
					} else if arg2[1] == BI {
						register_set(RBI, 0, ByteToInt8(arg1[1])*register_get(arg2[1], int(arg2[2])).(int8))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case INT16:
					if arg2[0] != REG {
						register_set(RSI, 0, BytesToInt16(arg1[1:])*BytesToInt16(arg2[1:]))
					} else if arg2[1] == SI {
						register_set(RSI, 0, BytesToInt16(arg1[1:])*register_get(arg2[1], int(arg2[2])).(int16))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case INT32:
					if arg2[0] != REG {
						register_set(RLI, 0, BytesToInt32(arg1[1:])*BytesToInt32(arg2[1:]))
					} else if arg2[1] == LI {
						register_set(RLI, 0, BytesToInt32(arg1[1:])*register_get(arg2[1], int(arg2[2])).(int32))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case INT64:
					if arg2[0] != REG {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])*BytesToInt64(arg2[1:]))
					} else if arg2[1] == LLI {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])*register_get(arg2[1], int(arg2[2])).(int64))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case UINT8:
					if arg2[0] != REG {
						register_set(RBUI, 0, ByteToUint8(arg1[1])*ByteToUint8(arg2[1]))
					} else if arg2[1] == BUI {
						register_set(RBUI, 0, ByteToUint8(arg1[1])*register_get(arg2[1], int(arg2[2])).(uint8))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case UINT16:
					if arg2[0] != REG {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])*BytesToUint16(arg2[1:]))
					} else if arg2[1] == SUI {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])*register_get(arg2[1], int(arg2[2])).(uint16))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case UINT32:
					if arg2[0] != REG {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])*BytesToUint32(arg2[1:]))
					} else if arg2[1] == LUI {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])*register_get(arg2[1], int(arg2[2])).(uint32))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case UINT64:
					if arg2[0] != REG {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])*BytesToUint64(arg2[1:]))
					} else if arg2[1] == LLUI {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])*register_get(arg2[1], int(arg2[2])).(uint64))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case FLOAT32:
					if arg2[0] != REG {
						register_set(RLF, 0, BytesToFloat32(arg1[1:])*BytesToFloat32(arg2[1:]))
					} else if arg2[1] == LF {
						register_set(RLF, 0, BytesToFloat32(arg1[1:])*register_get(arg2[1], int(arg2[2])).(float32))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case FLOAT64:
					if arg2[0] != REG {
						register_set(RLLF, 0, BytesToFloat64(arg1[1:])*BytesToFloat64(arg2[1:]))
					} else if arg2[1] == LLF {
						register_set(RLLF, 0, BytesToFloat64(arg1[1:])*register_get(arg2[1], int(arg2[2])).(float64))
					} else {
						panic("Invalid argument for MUL\n")
					}
				case REG:
					if arg1[1] == S || arg1[1] == RS {
						panic("Invalid argument for MUL\n")
					}

					val1 := ConvertToValue(arg1)
					val2 := ConvertToValue(arg2)

					switch arg1[1] {
					case BI, RBI:
						register_set(RBI, 0, val1.(int8)*val2.(int8))
					case SI, RSI:
						register_set(RSI, 0, val1.(int16)*val2.(int16))
					case LI, RLI:
						register_set(RLI, 0, val1.(int32)*val2.(int32))
					case LLI, RLLI:
						register_set(RLLI, 0, val1.(int64)*val2.(int64))
					case BUI, RBUI:
						register_set(RBUI, 0, val1.(uint8)*val2.(uint8))
					case SUI, RSUI:
						register_set(RSUI, 0, val1.(uint16)*val2.(uint16))
					case LUI, RLUI:
						register_set(RLUI, 0, val1.(uint32)*val2.(uint32))
					case LLUI, RLLUI:
						register_set(RLLUI, 0, val1.(uint64)*val2.(uint64))
					case LF, RLF:
						register_set(RLF, 0, val1.(float32)*val2.(float32))
					case LLF, RLLF:
						register_set(RLLF, 0, val1.(float64)*val2.(float64))
					}
				}
			}
		}
	case 3:
		return func(arg1 []byte, arg2 []byte) {
			if (arg1[0] < INT8 || arg2[0] < INT8 || arg1[0] > FLOAT64 || arg2[0] > FLOAT64) && (arg1[0] != REG && arg2[0] != REG) {
				panic("Invalid arguments for DIV\n")
			} else if arg1[0] != arg2[0] && arg1[0] != REG && arg2[0] != REG {
				panic("Arguments for DIV are different types\n")
			} else {
				switch arg1[0] {
				case INT8:
					if arg2[0] != REG {
						register_set(RBI, 0, ByteToInt8(arg1[1])/ByteToInt8(arg2[1]))
					} else if arg2[1] == BI {
						register_set(RBI, 0, ByteToInt8(arg1[1])/register_get(arg2[1], int(arg2[2])).(int8))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case INT16:
					if arg2[0] != REG {
						register_set(RSI, 0, BytesToInt16(arg1[1:])/BytesToInt16(arg2[1:]))
					} else if arg2[1] == SI {
						register_set(RSI, 0, BytesToInt16(arg1[1:])/register_get(arg2[1], int(arg2[2])).(int16))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case INT32:
					if arg2[0] != REG {
						register_set(RLI, 0, BytesToInt32(arg1[1:])/BytesToInt32(arg2[1:]))
					} else if arg2[1] == LI {
						register_set(RLI, 0, BytesToInt32(arg1[1:])/register_get(arg2[1], int(arg2[2])).(int32))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case INT64:
					if arg2[0] != REG {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])/BytesToInt64(arg2[1:]))
					} else if arg2[1] == LLI {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])/register_get(arg2[1], int(arg2[2])).(int64))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case UINT8:
					if arg2[0] != REG {
						register_set(RBUI, 0, ByteToUint8(arg1[1])/ByteToUint8(arg2[1]))
					} else if arg2[1] == BUI {
						register_set(RBUI, 0, ByteToUint8(arg1[1])/register_get(arg2[1], int(arg2[2])).(uint8))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case UINT16:
					if arg2[0] != REG {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])/BytesToUint16(arg2[1:]))
					} else if arg2[1] == SUI {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])/register_get(arg2[1], int(arg2[2])).(uint16))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case UINT32:
					if arg2[0] != REG {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])/BytesToUint32(arg2[1:]))
					} else if arg2[1] == LUI {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])/register_get(arg2[1], int(arg2[2])).(uint32))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case UINT64:
					if arg2[0] != REG {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])/BytesToUint64(arg2[1:]))
					} else if arg2[1] == LLUI {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])/register_get(arg2[1], int(arg2[2])).(uint64))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case FLOAT32:
					if arg2[0] != REG {
						register_set(RLF, 0, BytesToFloat32(arg1[1:])/BytesToFloat32(arg2[1:]))
					} else if arg2[1] == LF {
						register_set(RLF, 0, BytesToFloat32(arg1[1:])/register_get(arg2[1], int(arg2[2])).(float32))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case FLOAT64:
					if arg2[0] != REG {
						register_set(RLLF, 0, BytesToFloat64(arg1[1:])/BytesToFloat64(arg2[1:]))
					} else if arg2[1] == LLF {
						register_set(RLLF, 0, BytesToFloat64(arg1[1:])/register_get(arg2[1], int(arg2[2])).(float64))
					} else {
						panic("Invalid argument for DIV\n")
					}
				case REG:
					if arg1[1] == S || arg1[1] == RS {
						panic("Invalid argument for DIV\n")
					}

					val1 := ConvertToValue(arg1)
					val2 := ConvertToValue(arg2)

					switch arg1[1] {
					case BI, RBI:
						register_set(RBI, 0, val1.(int8)/val2.(int8))
					case SI, RSI:
						register_set(RSI, 0, val1.(int16)/val2.(int16))
					case LI, RLI:
						register_set(RLI, 0, val1.(int32)/val2.(int32))
					case LLI, RLLI:
						register_set(RLLI, 0, val1.(int64)/val2.(int64))
					case BUI, RBUI:
						register_set(RBUI, 0, val1.(uint8)/val2.(uint8))
					case SUI, RSUI:
						register_set(RSUI, 0, val1.(uint16)/val2.(uint16))
					case LUI, RLUI:
						register_set(RLUI, 0, val1.(uint32)/val2.(uint32))
					case LLUI, RLLUI:
						register_set(RLLUI, 0, val1.(uint64)/val2.(uint64))
					case LF, RLF:
						register_set(RLF, 0, val1.(float32)/val2.(float32))
					case LLF, RLLF:
						register_set(RLLF, 0, val1.(float64)/val2.(float64))
					}
				}
			}
		}
	case 4:
		return func(arg1 []byte, arg2 []byte) {
			if (arg1[0] < INT8 || arg2[0] < INT8 || arg1[0] > UINT64 || arg2[0] > UINT64) && (arg1[0] != REG && arg2[0] != REG) {
				panic("Invalid arguments for MOD\n")
			} else if arg1[0] != arg2[0] && arg1[0] != REG && arg2[0] != REG {
				panic("Arguments for MOD are different types\n")
			} else {
				switch arg1[0] {
				case INT8:
					if arg2[0] != REG {
						register_set(RBI, 0, ByteToInt8(arg1[1])%ByteToInt8(arg2[1]))
					} else if arg2[1] == BI {
						register_set(RBI, 0, ByteToInt8(arg1[1])%register_get(arg2[1], int(arg2[2])).(int8))
					} else {
						panic("Invalid argument for MOD\n")
					}
				case INT16:
					if arg2[0] != REG {
						register_set(RSI, 0, BytesToInt16(arg1[1:])%BytesToInt16(arg2[1:]))
					} else if arg2[1] == SI {
						register_set(RSI, 0, BytesToInt16(arg1[1:])%register_get(arg2[1], int(arg2[2])).(int16))
					} else {
						panic("Invalid argument for MOD\n")
					}
				case INT32:
					if arg2[0] != REG {
						register_set(RLI, 0, BytesToInt32(arg1[1:])%BytesToInt32(arg2[1:]))
					} else if arg2[1] == LI {
						register_set(RLI, 0, BytesToInt32(arg1[1:])%register_get(arg2[1], int(arg2[2])).(int32))
					} else {
						panic("Invalid argument for MOD\n")
					}
				case INT64:
					if arg2[0] != REG {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])%BytesToInt64(arg2[1:]))
					} else if arg2[1] == LLI {
						register_set(RLLI, 0, BytesToInt64(arg1[1:])%register_get(arg2[1], int(arg2[2])).(int64))
					} else {
						panic("Invalid argument for MOD\n")
					}
				case UINT8:
					if arg2[0] != REG {
						register_set(RBUI, 0, ByteToUint8(arg1[1])%ByteToUint8(arg2[1]))
					} else if arg2[1] == BUI {
						register_set(RBUI, 0, ByteToUint8(arg1[1])%register_get(arg2[1], int(arg2[2])).(uint8))
					} else {
						panic("Invalid argument for MOD\n")
					}
				case UINT16:
					if arg2[0] != REG {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])%BytesToUint16(arg2[1:]))
					} else if arg2[1] == SUI {
						register_set(RSUI, 0, BytesToUint16(arg1[1:])%register_get(arg2[1], int(arg2[2])).(uint16))
					} else {
						panic("Invalid argument for MOD\n")
					}
				case UINT32:
					if arg2[0] != REG {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])%BytesToUint32(arg2[1:]))
					} else if arg2[1] == LUI {
						register_set(RLUI, 0, BytesToUint32(arg1[1:])%register_get(arg2[1], int(arg2[2])).(uint32))
					} else {
						panic("Invalid argument for MOD\n")
					}
				case UINT64:
					if arg2[0] != REG {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])%BytesToUint64(arg2[1:]))
					} else if arg2[1] == LLUI {
						register_set(RLLUI, 0, BytesToUint64(arg1[1:])%register_get(arg2[1], int(arg2[2])).(uint64))
					} else {
						panic("Invalid argument for MOD\n")
					}
				case REG:
					if (arg1[1] >= S && arg1[1] <= LLF) || (arg1[1] >= RS && arg1[1] <= RLLF) {
						panic("Invalid argument for MOD\n")
					}

					val1 := ConvertToValue(arg1)
					val2 := ConvertToValue(arg2)

					switch arg1[1] {
					case BI, RBI:
						register_set(RBI, 0, val1.(int8)%val2.(int8))
					case SI, RSI:
						register_set(RSI, 0, val1.(int16)%val2.(int16))
					case LI, RLI:
						register_set(RLI, 0, val1.(int32)%val2.(int32))
					case LLI, RLLI:
						register_set(RLLI, 0, val1.(int64)%val2.(int64))
					case BUI, RBUI:
						register_set(RBUI, 0, val1.(uint8)%val2.(uint8))
					case SUI, RSUI:
						register_set(RSUI, 0, val1.(uint16)%val2.(uint16))
					case LUI, RLUI:
						register_set(RLUI, 0, val1.(uint32)%val2.(uint32))
					case LLUI, RLLUI:
						register_set(RLLUI, 0, val1.(uint64)%val2.(uint64))
					}
				}
			}
		}
	}
	return nil
}
