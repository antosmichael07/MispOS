package main

var REGISTER_CMP = [2][]byte{}

func compare(function *Function, index *int) func([]byte, []byte) {
	return func(arg1 []byte, arg2 []byte) {
		switch function.Instructions[*index] {
		case JE:
			if convert_to_value(REGISTER_CMP[0]) == convert_to_value(REGISTER_CMP[1]) {
				*index = function.Labels[arg1[1]]
			} else {
				*index = function.Labels[arg2[1]]
			}
		case JNE:
			if convert_to_value(REGISTER_CMP[0]) != convert_to_value(REGISTER_CMP[1]) {
				*index = function.Labels[arg1[1]]
			} else {
				*index = function.Labels[arg2[1]]
			}
		case JG:
			switch REGISTER_CMP[0][0] {
			case INT8:
				if convert_to_value(REGISTER_CMP[0]).(int8) > convert_to_value(REGISTER_CMP[1]).(int8) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT16:
				if convert_to_value(REGISTER_CMP[0]).(int16) > convert_to_value(REGISTER_CMP[1]).(int16) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT32:
				if convert_to_value(REGISTER_CMP[0]).(int32) > convert_to_value(REGISTER_CMP[1]).(int32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT64:
				if convert_to_value(REGISTER_CMP[0]).(int64) > convert_to_value(REGISTER_CMP[1]).(int64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT8:
				if convert_to_value(REGISTER_CMP[0]).(uint8) > convert_to_value(REGISTER_CMP[1]).(uint8) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT16:
				if convert_to_value(REGISTER_CMP[0]).(uint16) > convert_to_value(REGISTER_CMP[1]).(uint16) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT32:
				if convert_to_value(REGISTER_CMP[0]).(uint32) > convert_to_value(REGISTER_CMP[1]).(uint32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT64:
				if convert_to_value(REGISTER_CMP[0]).(uint64) > convert_to_value(REGISTER_CMP[1]).(uint64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case FLOAT32:
				if convert_to_value(REGISTER_CMP[0]).(float32) > convert_to_value(REGISTER_CMP[1]).(float32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case FLOAT64:
				if convert_to_value(REGISTER_CMP[0]).(float64) > convert_to_value(REGISTER_CMP[1]).(float64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case REG:
				switch REGISTER_CMP[0][1] {
				case BI, RBI:
					if convert_to_value(REGISTER_CMP[0]).(int8) > convert_to_value(REGISTER_CMP[1]).(int8) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case SI, RSI:
					if convert_to_value(REGISTER_CMP[0]).(int16) > convert_to_value(REGISTER_CMP[1]).(int16) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LI, RLI:
					if convert_to_value(REGISTER_CMP[0]).(int32) > convert_to_value(REGISTER_CMP[1]).(int32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLI, RLLI:
					if convert_to_value(REGISTER_CMP[0]).(int64) > convert_to_value(REGISTER_CMP[1]).(int64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case BUI, RBUI:
					if convert_to_value(REGISTER_CMP[0]).(uint8) > convert_to_value(REGISTER_CMP[1]).(uint8) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case SUI, RSUI:
					if convert_to_value(REGISTER_CMP[0]).(uint16) > convert_to_value(REGISTER_CMP[1]).(uint16) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LUI, RLUI:
					if convert_to_value(REGISTER_CMP[0]).(uint32) > convert_to_value(REGISTER_CMP[1]).(uint32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLUI, RLLUI:
					if convert_to_value(REGISTER_CMP[0]).(uint64) > convert_to_value(REGISTER_CMP[1]).(uint64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LF, RLF:
					if convert_to_value(REGISTER_CMP[0]).(float32) > convert_to_value(REGISTER_CMP[1]).(float32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLF, RLLF:
					if convert_to_value(REGISTER_CMP[0]).(float64) > convert_to_value(REGISTER_CMP[1]).(float64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case JGE:
			switch REGISTER_CMP[0][0] {
			case INT8:
				if convert_to_value(REGISTER_CMP[0]).(int8) >= convert_to_value(REGISTER_CMP[1]).(int8) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT16:
				if convert_to_value(REGISTER_CMP[0]).(int16) >= convert_to_value(REGISTER_CMP[1]).(int16) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT32:
				if convert_to_value(REGISTER_CMP[0]).(int32) >= convert_to_value(REGISTER_CMP[1]).(int32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT64:
				if convert_to_value(REGISTER_CMP[0]).(int64) >= convert_to_value(REGISTER_CMP[1]).(int64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT8:
				if convert_to_value(REGISTER_CMP[0]).(uint8) >= convert_to_value(REGISTER_CMP[1]).(uint8) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT16:
				if convert_to_value(REGISTER_CMP[0]).(uint16) >= convert_to_value(REGISTER_CMP[1]).(uint16) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT32:
				if convert_to_value(REGISTER_CMP[0]).(uint32) >= convert_to_value(REGISTER_CMP[1]).(uint32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT64:
				if convert_to_value(REGISTER_CMP[0]).(uint64) >= convert_to_value(REGISTER_CMP[1]).(uint64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case FLOAT32:
				if convert_to_value(REGISTER_CMP[0]).(float32) >= convert_to_value(REGISTER_CMP[1]).(float32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case FLOAT64:
				if convert_to_value(REGISTER_CMP[0]).(float64) >= convert_to_value(REGISTER_CMP[1]).(float64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case REG:
				switch REGISTER_CMP[0][1] {
				case BI, RBI:
					if convert_to_value(REGISTER_CMP[0]).(int8) >= convert_to_value(REGISTER_CMP[1]).(int8) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case SI, RSI:
					if convert_to_value(REGISTER_CMP[0]).(int16) >= convert_to_value(REGISTER_CMP[1]).(int16) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LI, RLI:
					if convert_to_value(REGISTER_CMP[0]).(int32) >= convert_to_value(REGISTER_CMP[1]).(int32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLI, RLLI:
					if convert_to_value(REGISTER_CMP[0]).(int64) >= convert_to_value(REGISTER_CMP[1]).(int64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case BUI, RBUI:
					if convert_to_value(REGISTER_CMP[0]).(uint8) >= convert_to_value(REGISTER_CMP[1]).(uint8) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case SUI, RSUI:
					if convert_to_value(REGISTER_CMP[0]).(uint16) >= convert_to_value(REGISTER_CMP[1]).(uint16) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LUI, RLUI:
					if convert_to_value(REGISTER_CMP[0]).(uint32) >= convert_to_value(REGISTER_CMP[1]).(uint32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLUI, RLLUI:
					if convert_to_value(REGISTER_CMP[0]).(uint64) >= convert_to_value(REGISTER_CMP[1]).(uint64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LF, RLF:
					if convert_to_value(REGISTER_CMP[0]).(float32) >= convert_to_value(REGISTER_CMP[1]).(float32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLF, RLLF:
					if convert_to_value(REGISTER_CMP[0]).(float64) >= convert_to_value(REGISTER_CMP[1]).(float64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case JL:
			switch REGISTER_CMP[0][0] {
			case INT8:
				if convert_to_value(REGISTER_CMP[0]).(int8) < convert_to_value(REGISTER_CMP[1]).(int8) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT16:
				if convert_to_value(REGISTER_CMP[0]).(int16) < convert_to_value(REGISTER_CMP[1]).(int16) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT32:
				if convert_to_value(REGISTER_CMP[0]).(int32) < convert_to_value(REGISTER_CMP[1]).(int32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT64:
				if convert_to_value(REGISTER_CMP[0]).(int64) < convert_to_value(REGISTER_CMP[1]).(int64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT8:
				if convert_to_value(REGISTER_CMP[0]).(uint8) < convert_to_value(REGISTER_CMP[1]).(uint8) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT16:
				if convert_to_value(REGISTER_CMP[0]).(uint16) < convert_to_value(REGISTER_CMP[1]).(uint16) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT32:
				if convert_to_value(REGISTER_CMP[0]).(uint32) < convert_to_value(REGISTER_CMP[1]).(uint32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT64:
				if convert_to_value(REGISTER_CMP[0]).(uint64) < convert_to_value(REGISTER_CMP[1]).(uint64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case FLOAT32:
				if convert_to_value(REGISTER_CMP[0]).(float32) < convert_to_value(REGISTER_CMP[1]).(float32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case FLOAT64:
				if convert_to_value(REGISTER_CMP[0]).(float64) < convert_to_value(REGISTER_CMP[1]).(float64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case REG:
				switch REGISTER_CMP[0][1] {
				case BI, RBI:
					if convert_to_value(REGISTER_CMP[0]).(int8) < convert_to_value(REGISTER_CMP[1]).(int8) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case SI, RSI:
					if convert_to_value(REGISTER_CMP[0]).(int16) < convert_to_value(REGISTER_CMP[1]).(int16) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LI, RLI:
					if convert_to_value(REGISTER_CMP[0]).(int32) < convert_to_value(REGISTER_CMP[1]).(int32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLI, RLLI:
					if convert_to_value(REGISTER_CMP[0]).(int64) < convert_to_value(REGISTER_CMP[1]).(int64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case BUI, RBUI:
					if convert_to_value(REGISTER_CMP[0]).(uint8) < convert_to_value(REGISTER_CMP[1]).(uint8) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case SUI, RSUI:
					if convert_to_value(REGISTER_CMP[0]).(uint16) < convert_to_value(REGISTER_CMP[1]).(uint16) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LUI, RLUI:
					if convert_to_value(REGISTER_CMP[0]).(uint32) < convert_to_value(REGISTER_CMP[1]).(uint32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLUI, RLLUI:
					if convert_to_value(REGISTER_CMP[0]).(uint64) < convert_to_value(REGISTER_CMP[1]).(uint64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LF, RLF:
					if convert_to_value(REGISTER_CMP[0]).(float32) < convert_to_value(REGISTER_CMP[1]).(float32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLF, RLLF:
					if convert_to_value(REGISTER_CMP[0]).(float64) < convert_to_value(REGISTER_CMP[1]).(float64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case JLE:
			switch REGISTER_CMP[0][0] {
			case INT8:
				if convert_to_value(REGISTER_CMP[0]).(int8) <= convert_to_value(REGISTER_CMP[1]).(int8) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT16:
				if convert_to_value(REGISTER_CMP[0]).(int16) <= convert_to_value(REGISTER_CMP[1]).(int16) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT32:
				if convert_to_value(REGISTER_CMP[0]).(int32) <= convert_to_value(REGISTER_CMP[1]).(int32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case INT64:
				if convert_to_value(REGISTER_CMP[0]).(int64) <= convert_to_value(REGISTER_CMP[1]).(int64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT8:
				if convert_to_value(REGISTER_CMP[0]).(uint8) <= convert_to_value(REGISTER_CMP[1]).(uint8) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT16:
				if convert_to_value(REGISTER_CMP[0]).(uint16) <= convert_to_value(REGISTER_CMP[1]).(uint16) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT32:
				if convert_to_value(REGISTER_CMP[0]).(uint32) <= convert_to_value(REGISTER_CMP[1]).(uint32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case UINT64:
				if convert_to_value(REGISTER_CMP[0]).(uint64) <= convert_to_value(REGISTER_CMP[1]).(uint64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case FLOAT32:
				if convert_to_value(REGISTER_CMP[0]).(float32) <= convert_to_value(REGISTER_CMP[1]).(float32) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case FLOAT64:
				if convert_to_value(REGISTER_CMP[0]).(float64) <= convert_to_value(REGISTER_CMP[1]).(float64) {
					*index = function.Labels[arg1[1]]
				} else {
					*index = function.Labels[arg2[1]]
				}
			case REG:
				switch REGISTER_CMP[0][1] {
				case BI, RBI:
					if convert_to_value(REGISTER_CMP[0]).(int8) <= convert_to_value(REGISTER_CMP[1]).(int8) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case SI, RSI:
					if convert_to_value(REGISTER_CMP[0]).(int16) <= convert_to_value(REGISTER_CMP[1]).(int16) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LI, RLI:
					if convert_to_value(REGISTER_CMP[0]).(int32) <= convert_to_value(REGISTER_CMP[1]).(int32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLI, RLLI:
					if convert_to_value(REGISTER_CMP[0]).(int64) <= convert_to_value(REGISTER_CMP[1]).(int64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case BUI, RBUI:
					if convert_to_value(REGISTER_CMP[0]).(uint8) <= convert_to_value(REGISTER_CMP[1]).(uint8) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case SUI, RSUI:
					if convert_to_value(REGISTER_CMP[0]).(uint16) <= convert_to_value(REGISTER_CMP[1]).(uint16) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LUI, RLUI:
					if convert_to_value(REGISTER_CMP[0]).(uint32) <= convert_to_value(REGISTER_CMP[1]).(uint32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLUI, RLLUI:
					if convert_to_value(REGISTER_CMP[0]).(uint64) <= convert_to_value(REGISTER_CMP[1]).(uint64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LF, RLF:
					if convert_to_value(REGISTER_CMP[0]).(float32) <= convert_to_value(REGISTER_CMP[1]).(float32) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
					}
				case LLF, RLLF:
					if convert_to_value(REGISTER_CMP[0]).(float64) <= convert_to_value(REGISTER_CMP[1]).(float64) {
						*index = function.Labels[arg1[1]]
					} else {
						*index = function.Labels[arg2[1]]
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
