package main

var REGISTER_CMP = [2][]byte{}

func compare(function *Function, index *int) func([]byte, []byte) {
	return func(arg1 []byte, arg2 []byte) {
		switch function.Instructions[*index] {
		case byte(JE):
			if ConvertToValue(REGISTER_CMP[0]) == ConvertToValue(REGISTER_CMP[1]) {
				*index = function.Labels[int(arg1[1])]
			} else {
				*index = function.Labels[int(arg2[1])]
			}
		case byte(JNE):
			if ConvertToValue(REGISTER_CMP[0]) != ConvertToValue(REGISTER_CMP[1]) {
				*index = function.Labels[int(arg1[1])]
			} else {
				*index = function.Labels[int(arg2[1])]
			}
		case byte(JG):
			switch REGISTER_CMP[0][0] {
			case byte(INT8):
				if ConvertToValue(REGISTER_CMP[0]).(int8) > ConvertToValue(REGISTER_CMP[1]).(int8) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT16):
				if ConvertToValue(REGISTER_CMP[0]).(int16) > ConvertToValue(REGISTER_CMP[1]).(int16) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT32):
				if ConvertToValue(REGISTER_CMP[0]).(int32) > ConvertToValue(REGISTER_CMP[1]).(int32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT64):
				if ConvertToValue(REGISTER_CMP[0]).(int64) > ConvertToValue(REGISTER_CMP[1]).(int64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT8):
				if ConvertToValue(REGISTER_CMP[0]).(uint8) > ConvertToValue(REGISTER_CMP[1]).(uint8) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT16):
				if ConvertToValue(REGISTER_CMP[0]).(uint16) > ConvertToValue(REGISTER_CMP[1]).(uint16) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT32):
				if ConvertToValue(REGISTER_CMP[0]).(uint32) > ConvertToValue(REGISTER_CMP[1]).(uint32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT64):
				if ConvertToValue(REGISTER_CMP[0]).(uint64) > ConvertToValue(REGISTER_CMP[1]).(uint64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(FLOAT32):
				if ConvertToValue(REGISTER_CMP[0]).(float32) > ConvertToValue(REGISTER_CMP[1]).(float32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(FLOAT64):
				if ConvertToValue(REGISTER_CMP[0]).(float64) > ConvertToValue(REGISTER_CMP[1]).(float64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(REG):
				switch REGISTER_CMP[0][1] {
				case byte(BI), byte(RBI):
					if ConvertToValue(REGISTER_CMP[0]).(int8) > ConvertToValue(REGISTER_CMP[1]).(int8) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(SI), byte(RSI):
					if ConvertToValue(REGISTER_CMP[0]).(int16) > ConvertToValue(REGISTER_CMP[1]).(int16) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LI), byte(RLI):
					if ConvertToValue(REGISTER_CMP[0]).(int32) > ConvertToValue(REGISTER_CMP[1]).(int32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLI), byte(RLLI):
					if ConvertToValue(REGISTER_CMP[0]).(int64) > ConvertToValue(REGISTER_CMP[1]).(int64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(BUI), byte(RBUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint8) > ConvertToValue(REGISTER_CMP[1]).(uint8) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(SUI), byte(RSUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint16) > ConvertToValue(REGISTER_CMP[1]).(uint16) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LUI), byte(RLUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint32) > ConvertToValue(REGISTER_CMP[1]).(uint32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLUI), byte(RLLUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint64) > ConvertToValue(REGISTER_CMP[1]).(uint64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LF), byte(RLF):
					if ConvertToValue(REGISTER_CMP[0]).(float32) > ConvertToValue(REGISTER_CMP[1]).(float32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLF), byte(RLLF):
					if ConvertToValue(REGISTER_CMP[0]).(float64) > ConvertToValue(REGISTER_CMP[1]).(float64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case byte(JGE):
			switch REGISTER_CMP[0][0] {
			case byte(INT8):
				if ConvertToValue(REGISTER_CMP[0]).(int8) >= ConvertToValue(REGISTER_CMP[1]).(int8) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT16):
				if ConvertToValue(REGISTER_CMP[0]).(int16) >= ConvertToValue(REGISTER_CMP[1]).(int16) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT32):
				if ConvertToValue(REGISTER_CMP[0]).(int32) >= ConvertToValue(REGISTER_CMP[1]).(int32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT64):
				if ConvertToValue(REGISTER_CMP[0]).(int64) >= ConvertToValue(REGISTER_CMP[1]).(int64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT8):
				if ConvertToValue(REGISTER_CMP[0]).(uint8) >= ConvertToValue(REGISTER_CMP[1]).(uint8) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT16):
				if ConvertToValue(REGISTER_CMP[0]).(uint16) >= ConvertToValue(REGISTER_CMP[1]).(uint16) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT32):
				if ConvertToValue(REGISTER_CMP[0]).(uint32) >= ConvertToValue(REGISTER_CMP[1]).(uint32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT64):
				if ConvertToValue(REGISTER_CMP[0]).(uint64) >= ConvertToValue(REGISTER_CMP[1]).(uint64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(FLOAT32):
				if ConvertToValue(REGISTER_CMP[0]).(float32) >= ConvertToValue(REGISTER_CMP[1]).(float32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(FLOAT64):
				if ConvertToValue(REGISTER_CMP[0]).(float64) >= ConvertToValue(REGISTER_CMP[1]).(float64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(REG):
				switch REGISTER_CMP[0][1] {
				case byte(BI), byte(RBI):
					if ConvertToValue(REGISTER_CMP[0]).(int8) >= ConvertToValue(REGISTER_CMP[1]).(int8) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(SI), byte(RSI):
					if ConvertToValue(REGISTER_CMP[0]).(int16) >= ConvertToValue(REGISTER_CMP[1]).(int16) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LI), byte(RLI):
					if ConvertToValue(REGISTER_CMP[0]).(int32) >= ConvertToValue(REGISTER_CMP[1]).(int32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLI), byte(RLLI):
					if ConvertToValue(REGISTER_CMP[0]).(int64) >= ConvertToValue(REGISTER_CMP[1]).(int64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(BUI), byte(RBUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint8) >= ConvertToValue(REGISTER_CMP[1]).(uint8) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(SUI), byte(RSUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint16) >= ConvertToValue(REGISTER_CMP[1]).(uint16) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LUI), byte(RLUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint32) >= ConvertToValue(REGISTER_CMP[1]).(uint32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLUI), byte(RLLUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint64) >= ConvertToValue(REGISTER_CMP[1]).(uint64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LF), byte(RLF):
					if ConvertToValue(REGISTER_CMP[0]).(float32) >= ConvertToValue(REGISTER_CMP[1]).(float32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLF), byte(RLLF):
					if ConvertToValue(REGISTER_CMP[0]).(float64) >= ConvertToValue(REGISTER_CMP[1]).(float64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case byte(JL):
			switch REGISTER_CMP[0][0] {
			case byte(INT8):
				if ConvertToValue(REGISTER_CMP[0]).(int8) < ConvertToValue(REGISTER_CMP[1]).(int8) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT16):
				if ConvertToValue(REGISTER_CMP[0]).(int16) < ConvertToValue(REGISTER_CMP[1]).(int16) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT32):
				if ConvertToValue(REGISTER_CMP[0]).(int32) < ConvertToValue(REGISTER_CMP[1]).(int32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT64):
				if ConvertToValue(REGISTER_CMP[0]).(int64) < ConvertToValue(REGISTER_CMP[1]).(int64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT8):
				if ConvertToValue(REGISTER_CMP[0]).(uint8) < ConvertToValue(REGISTER_CMP[1]).(uint8) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT16):
				if ConvertToValue(REGISTER_CMP[0]).(uint16) < ConvertToValue(REGISTER_CMP[1]).(uint16) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT32):
				if ConvertToValue(REGISTER_CMP[0]).(uint32) < ConvertToValue(REGISTER_CMP[1]).(uint32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT64):
				if ConvertToValue(REGISTER_CMP[0]).(uint64) < ConvertToValue(REGISTER_CMP[1]).(uint64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(FLOAT32):
				if ConvertToValue(REGISTER_CMP[0]).(float32) < ConvertToValue(REGISTER_CMP[1]).(float32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(FLOAT64):
				if ConvertToValue(REGISTER_CMP[0]).(float64) < ConvertToValue(REGISTER_CMP[1]).(float64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(REG):
				switch REGISTER_CMP[0][1] {
				case byte(BI), byte(RBI):
					if ConvertToValue(REGISTER_CMP[0]).(int8) < ConvertToValue(REGISTER_CMP[1]).(int8) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(SI), byte(RSI):
					if ConvertToValue(REGISTER_CMP[0]).(int16) < ConvertToValue(REGISTER_CMP[1]).(int16) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LI), byte(RLI):
					if ConvertToValue(REGISTER_CMP[0]).(int32) < ConvertToValue(REGISTER_CMP[1]).(int32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLI), byte(RLLI):
					if ConvertToValue(REGISTER_CMP[0]).(int64) < ConvertToValue(REGISTER_CMP[1]).(int64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(BUI), byte(RBUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint8) < ConvertToValue(REGISTER_CMP[1]).(uint8) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(SUI), byte(RSUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint16) < ConvertToValue(REGISTER_CMP[1]).(uint16) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LUI), byte(RLUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint32) < ConvertToValue(REGISTER_CMP[1]).(uint32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLUI), byte(RLLUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint64) < ConvertToValue(REGISTER_CMP[1]).(uint64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LF), byte(RLF):
					if ConvertToValue(REGISTER_CMP[0]).(float32) < ConvertToValue(REGISTER_CMP[1]).(float32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLF), byte(RLLF):
					if ConvertToValue(REGISTER_CMP[0]).(float64) < ConvertToValue(REGISTER_CMP[1]).(float64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case byte(JLE):
			switch REGISTER_CMP[0][0] {
			case byte(INT8):
				if ConvertToValue(REGISTER_CMP[0]).(int8) <= ConvertToValue(REGISTER_CMP[1]).(int8) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT16):
				if ConvertToValue(REGISTER_CMP[0]).(int16) <= ConvertToValue(REGISTER_CMP[1]).(int16) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT32):
				if ConvertToValue(REGISTER_CMP[0]).(int32) <= ConvertToValue(REGISTER_CMP[1]).(int32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(INT64):
				if ConvertToValue(REGISTER_CMP[0]).(int64) <= ConvertToValue(REGISTER_CMP[1]).(int64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT8):
				if ConvertToValue(REGISTER_CMP[0]).(uint8) <= ConvertToValue(REGISTER_CMP[1]).(uint8) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT16):
				if ConvertToValue(REGISTER_CMP[0]).(uint16) <= ConvertToValue(REGISTER_CMP[1]).(uint16) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT32):
				if ConvertToValue(REGISTER_CMP[0]).(uint32) <= ConvertToValue(REGISTER_CMP[1]).(uint32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(UINT64):
				if ConvertToValue(REGISTER_CMP[0]).(uint64) <= ConvertToValue(REGISTER_CMP[1]).(uint64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(FLOAT32):
				if ConvertToValue(REGISTER_CMP[0]).(float32) <= ConvertToValue(REGISTER_CMP[1]).(float32) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(FLOAT64):
				if ConvertToValue(REGISTER_CMP[0]).(float64) <= ConvertToValue(REGISTER_CMP[1]).(float64) {
					*index = function.Labels[int(arg1[1])]
				} else {
					*index = function.Labels[int(arg2[1])]
				}
			case byte(REG):
				switch REGISTER_CMP[0][1] {
				case byte(BI), byte(RBI):
					if ConvertToValue(REGISTER_CMP[0]).(int8) <= ConvertToValue(REGISTER_CMP[1]).(int8) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(SI), byte(RSI):
					if ConvertToValue(REGISTER_CMP[0]).(int16) <= ConvertToValue(REGISTER_CMP[1]).(int16) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LI), byte(RLI):
					if ConvertToValue(REGISTER_CMP[0]).(int32) <= ConvertToValue(REGISTER_CMP[1]).(int32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLI), byte(RLLI):
					if ConvertToValue(REGISTER_CMP[0]).(int64) <= ConvertToValue(REGISTER_CMP[1]).(int64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(BUI), byte(RBUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint8) <= ConvertToValue(REGISTER_CMP[1]).(uint8) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(SUI), byte(RSUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint16) <= ConvertToValue(REGISTER_CMP[1]).(uint16) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LUI), byte(RLUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint32) <= ConvertToValue(REGISTER_CMP[1]).(uint32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLUI), byte(RLLUI):
					if ConvertToValue(REGISTER_CMP[0]).(uint64) <= ConvertToValue(REGISTER_CMP[1]).(uint64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LF), byte(RLF):
					if ConvertToValue(REGISTER_CMP[0]).(float32) <= ConvertToValue(REGISTER_CMP[1]).(float32) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				case byte(LLF), byte(RLLF):
					if ConvertToValue(REGISTER_CMP[0]).(float64) <= ConvertToValue(REGISTER_CMP[1]).(float64) {
						*index = function.Labels[int(arg1[1])]
					} else {
						*index = function.Labels[int(arg2[1])]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		}
	}
}
