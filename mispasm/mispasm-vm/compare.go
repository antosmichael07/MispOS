package main

var register_cmp = [2][]byte{}

func compare(function *Function, index *int) func([]byte, []byte) {
	return func(arg1 []byte, arg2 []byte) {
		switch function.instructions[*index] {
		case je:
			if convert_to_value[register_cmp[0][0]](register_cmp[0]) == convert_to_value[register_cmp[1][0]](register_cmp[1]) {
				*index = function.labels[arg1[1]]
			} else {
				*index = function.labels[arg2[1]]
			}
		case jne:
			if convert_to_value[register_cmp[0][0]](register_cmp[0]) != convert_to_value[register_cmp[1][0]](register_cmp[1]) {
				*index = function.labels[arg1[1]]
			} else {
				*index = function.labels[arg2[1]]
			}
		case jg:
			switch register_cmp[0][0] {
			case t_int8:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int16:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint8:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint16:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_float32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_float64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_reg:
				switch register_cmp[0][1] {
				case bi, rbi:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case si, rsi:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case li, rli:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lli, rlli:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case bui, rbui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case sui, rsui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lui, rlui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case llui, rllui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lf, rlf:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case llf, rllf:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case jge:
			switch register_cmp[0][0] {
			case t_int8:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int16:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint8:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint16:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_float32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_float64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_reg:
				switch register_cmp[0][1] {
				case bi, rbi:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case si, rsi:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case li, rli:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lli, rlli:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case bui, rbui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case sui, rsui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lui, rlui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case llui, rllui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lf, rlf:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case llf, rllf:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case jl:
			switch register_cmp[0][0] {
			case t_int8:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int16:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint8:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint16:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_float32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_float64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_reg:
				switch register_cmp[0][1] {
				case bi, rbi:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case si, rsi:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case li, rli:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lli, rlli:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case bui, rbui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case sui, rsui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lui, rlui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case llui, rllui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lf, rlf:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case llf, rllf:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				default:
					panic("Invalid argument type for CMP\n")
				}
			default:
				panic("Invalid argument type for CMP\n")
			}
		case jle:
			switch register_cmp[0][0] {
			case t_int8:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int16:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_int64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint8:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint16:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_uint64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_float32:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_float64:
				if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
					*index = function.labels[arg1[1]]
				} else {
					*index = function.labels[arg2[1]]
				}
			case t_reg:
				switch register_cmp[0][1] {
				case bi, rbi:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case si, rsi:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case li, rli:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lli, rlli:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case bui, rbui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case sui, rsui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lui, rlui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case llui, rllui:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case lf, rlf:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
					}
				case llf, rllf:
					if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
						*index = function.labels[arg1[1]]
					} else {
						*index = function.labels[arg2[1]]
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
