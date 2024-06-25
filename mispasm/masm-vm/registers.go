package main

const (
	BI byte = iota
	SI
	LI
	LLI
	BUI
	SUI
	LUI
	LLUI
	LF
	LLF
	S
	RBI
	RSI
	RLI
	RLLI
	RBUI
	RSUI
	RLUI
	RLLUI
	RLF
	RLLF
	RS
)

var REGISTER_BI = [256]int8{}
var REGISTER_SI = [256]int16{}
var REGISTER_LI = [256]int32{}
var REGISTER_LLI = [256]int64{}
var REGISTER_BUI = [256]uint8{}
var REGISTER_SUI = [256]uint16{}
var REGISTER_LUI = [256]uint32{}
var REGISTER_LLUI = [256]uint64{}
var REGISTER_LF = [256]float32{}
var REGISTER_LLF = [256]float64{}
var REGISTER_S = [256]string{}
var REGISTER_RBI = [256]int8{}
var REGISTER_RSI = [256]int16{}
var REGISTER_RLI = [256]int32{}
var REGISTER_RLLI = [256]int64{}
var REGISTER_RBUI = [256]uint8{}
var REGISTER_RSUI = [256]uint16{}
var REGISTER_RLUI = [256]uint32{}
var REGISTER_RLLUI = [256]uint64{}
var REGISTER_RLF = [256]float32{}
var REGISTER_RLLF = [256]float64{}
var REGISTER_RS = [256]string{}

func register_set(reg byte, index int, data any) {
	switch reg {
	case BI:
		REGISTER_BI[index] = data.(int8)
	case SI:
		REGISTER_SI[index] = data.(int16)
	case LI:
		REGISTER_LI[index] = data.(int32)
	case LLI:
		REGISTER_LLI[index] = data.(int64)
	case BUI:
		REGISTER_BUI[index] = data.(uint8)
	case SUI:
		REGISTER_SUI[index] = data.(uint16)
	case LUI:
		REGISTER_LUI[index] = data.(uint32)
	case LLUI:
		REGISTER_LLUI[index] = data.(uint64)
	case LF:
		REGISTER_LF[index] = data.(float32)
	case LLF:
		REGISTER_LLF[index] = data.(float64)
	case S:
		REGISTER_S[index] = data.(string)
	case RBI:
		REGISTER_RBI[index] = data.(int8)
	case RSI:
		REGISTER_RSI[index] = data.(int16)
	case RLI:
		REGISTER_RLI[index] = data.(int32)
	case RLLI:
		REGISTER_RLLI[index] = data.(int64)
	case RBUI:
		REGISTER_RBUI[index] = data.(uint8)
	case RSUI:
		REGISTER_RSUI[index] = data.(uint16)
	case RLUI:
		REGISTER_RLUI[index] = data.(uint32)
	case RLLUI:
		REGISTER_RLLUI[index] = data.(uint64)
	case RLF:
		REGISTER_RLF[index] = data.(float32)
	case RLLF:
		REGISTER_RLLF[index] = data.(float64)
	case RS:
		REGISTER_RS[index] = data.(string)
	}
}

func register_get(reg byte, index int) any {
	switch reg {
	case BI:
		return REGISTER_BI[index]
	case SI:
		return REGISTER_SI[index]
	case LI:
		return REGISTER_LI[index]
	case LLI:
		return REGISTER_LLI[index]
	case BUI:
		return REGISTER_BUI[index]
	case SUI:
		return REGISTER_SUI[index]
	case LUI:
		return REGISTER_LUI[index]
	case LLUI:
		return REGISTER_LLUI[index]
	case LF:
		return REGISTER_LF[index]
	case LLF:
		return REGISTER_LLF[index]
	case S:
		return REGISTER_S[index]
	case RBI:
		return REGISTER_RBI[index]
	case RSI:
		return REGISTER_RSI[index]
	case RLI:
		return REGISTER_RLI[index]
	case RLLI:
		return REGISTER_RLLI[index]
	case RBUI:
		return REGISTER_RBUI[index]
	case RSUI:
		return REGISTER_RSUI[index]
	case RLUI:
		return REGISTER_RLUI[index]
	case RLLUI:
		return REGISTER_RLLUI[index]
	case RLF:
		return REGISTER_RLF[index]
	case RLLF:
		return REGISTER_RLLF[index]
	case RS:
		return REGISTER_RS[index]
	}
	return nil
}
