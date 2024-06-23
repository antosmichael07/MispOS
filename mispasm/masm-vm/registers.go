package main

const (
	BI = iota
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
	case byte(BI):
		REGISTER_BI[index] = data.(int8)
	case byte(SI):
		REGISTER_SI[index] = data.(int16)
	case byte(LI):
		REGISTER_LI[index] = data.(int32)
	case byte(LLI):
		REGISTER_LLI[index] = data.(int64)
	case byte(BUI):
		REGISTER_BUI[index] = data.(uint8)
	case byte(SUI):
		REGISTER_SUI[index] = data.(uint16)
	case byte(LUI):
		REGISTER_LUI[index] = data.(uint32)
	case byte(LLUI):
		REGISTER_LLUI[index] = data.(uint64)
	case byte(LF):
		REGISTER_LF[index] = data.(float32)
	case byte(LLF):
		REGISTER_LLF[index] = data.(float64)
	case byte(S):
		REGISTER_S[index] = data.(string)
	case byte(RBI):
		REGISTER_RBI[index] = data.(int8)
	case byte(RSI):
		REGISTER_RSI[index] = data.(int16)
	case byte(RLI):
		REGISTER_RLI[index] = data.(int32)
	case byte(RLLI):
		REGISTER_RLLI[index] = data.(int64)
	case byte(RBUI):
		REGISTER_RBUI[index] = data.(uint8)
	case byte(RSUI):
		REGISTER_RSUI[index] = data.(uint16)
	case byte(RLUI):
		REGISTER_RLUI[index] = data.(uint32)
	case byte(RLLUI):
		REGISTER_RLLUI[index] = data.(uint64)
	case byte(RLF):
		REGISTER_RLF[index] = data.(float32)
	case byte(RLLF):
		REGISTER_RLLF[index] = data.(float64)
	case byte(RS):
		REGISTER_RS[index] = data.(string)
	}
}

func register_get(reg byte, index int) any {
	switch reg {
	case byte(BI):
		return REGISTER_BI[index]
	case byte(SI):
		return REGISTER_SI[index]
	case byte(LI):
		return REGISTER_LI[index]
	case byte(LLI):
		return REGISTER_LLI[index]
	case byte(BUI):
		return REGISTER_BUI[index]
	case byte(SUI):
		return REGISTER_SUI[index]
	case byte(LUI):
		return REGISTER_LUI[index]
	case byte(LLUI):
		return REGISTER_LLUI[index]
	case byte(LF):
		return REGISTER_LF[index]
	case byte(LLF):
		return REGISTER_LLF[index]
	case byte(S):
		return REGISTER_S[index]
	case byte(RBI):
		return REGISTER_RBI[index]
	case byte(RSI):
		return REGISTER_RSI[index]
	case byte(RLI):
		return REGISTER_RLI[index]
	case byte(RLLI):
		return REGISTER_RLLI[index]
	case byte(RBUI):
		return REGISTER_RBUI[index]
	case byte(RSUI):
		return REGISTER_RSUI[index]
	case byte(RLUI):
		return REGISTER_RLUI[index]
	case byte(RLLUI):
		return REGISTER_RLLUI[index]
	case byte(RLF):
		return REGISTER_RLF[index]
	case byte(RLLF):
		return REGISTER_RLLF[index]
	case byte(RS):
		return REGISTER_RS[index]
	}
	return nil
}
