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

var register_set = [22]func(byte, any){
	func(index byte, data any) { REGISTER_BI[index] = data.(int8) },
	func(index byte, data any) { REGISTER_SI[index] = data.(int16) },
	func(index byte, data any) { REGISTER_LI[index] = data.(int32) },
	func(index byte, data any) { REGISTER_LLI[index] = data.(int64) },
	func(index byte, data any) { REGISTER_BUI[index] = data.(uint8) },
	func(index byte, data any) { REGISTER_SUI[index] = data.(uint16) },
	func(index byte, data any) { REGISTER_LUI[index] = data.(uint32) },
	func(index byte, data any) { REGISTER_LLUI[index] = data.(uint64) },
	func(index byte, data any) { REGISTER_LF[index] = data.(float32) },
	func(index byte, data any) { REGISTER_LLF[index] = data.(float64) },
	func(index byte, data any) { REGISTER_S[index] = data.(string) },
	func(index byte, data any) { REGISTER_RBI[index] = data.(int8) },
	func(index byte, data any) { REGISTER_RSI[index] = data.(int16) },
	func(index byte, data any) { REGISTER_RLI[index] = data.(int32) },
	func(index byte, data any) { REGISTER_RLLI[index] = data.(int64) },
	func(index byte, data any) { REGISTER_RBUI[index] = data.(uint8) },
	func(index byte, data any) { REGISTER_RSUI[index] = data.(uint16) },
	func(index byte, data any) { REGISTER_RLUI[index] = data.(uint32) },
	func(index byte, data any) { REGISTER_RLLUI[index] = data.(uint64) },
	func(index byte, data any) { REGISTER_RLF[index] = data.(float32) },
	func(index byte, data any) { REGISTER_RLLF[index] = data.(float64) },
	func(index byte, data any) { REGISTER_RS[index] = data.(string) },
}

var register_get = [22]func(int) any{
	func(index int) any { return REGISTER_BI[index] },
	func(index int) any { return REGISTER_SI[index] },
	func(index int) any { return REGISTER_LI[index] },
	func(index int) any { return REGISTER_LLI[index] },
	func(index int) any { return REGISTER_BUI[index] },
	func(index int) any { return REGISTER_SUI[index] },
	func(index int) any { return REGISTER_LUI[index] },
	func(index int) any { return REGISTER_LLUI[index] },
	func(index int) any { return REGISTER_LF[index] },
	func(index int) any { return REGISTER_LLF[index] },
	func(index int) any { return REGISTER_S[index] },
	func(index int) any { return REGISTER_RBI[index] },
	func(index int) any { return REGISTER_RSI[index] },
	func(index int) any { return REGISTER_RLI[index] },
	func(index int) any { return REGISTER_RLLI[index] },
	func(index int) any { return REGISTER_RBUI[index] },
	func(index int) any { return REGISTER_RSUI[index] },
	func(index int) any { return REGISTER_RLUI[index] },
	func(index int) any { return REGISTER_RLLUI[index] },
	func(index int) any { return REGISTER_RLF[index] },
	func(index int) any { return REGISTER_RLLF[index] },
	func(index int) any { return REGISTER_RS[index] },
}
