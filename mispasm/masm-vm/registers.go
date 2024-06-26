package main

const (
	bi byte = iota
	si
	li
	lli
	bui
	sui
	lui
	llui
	lf
	llf
	s
	rbi
	rsi
	rli
	rlli
	rbui
	rsui
	rlui
	rllui
	rlf
	rllf
	rs
)

var register_bi = [256]int8{}
var register_si = [256]int16{}
var register_li = [256]int32{}
var register_lli = [256]int64{}
var register_bui = [256]uint8{}
var register_sui = [256]uint16{}
var register_lui = [256]uint32{}
var register_llui = [256]uint64{}
var register_lf = [256]float32{}
var register_llf = [256]float64{}
var register_s = [256]string{}
var register_rbi = [256]int8{}
var register_rsi = [256]int16{}
var register_rli = [256]int32{}
var register_rlli = [256]int64{}
var register_rbui = [256]uint8{}
var register_rsui = [256]uint16{}
var register_rlui = [256]uint32{}
var register_rllui = [256]uint64{}
var register_rlf = [256]float32{}
var register_rllf = [256]float64{}
var register_rs = [256]string{}

var register_set = [22]func(byte, any){
	func(index byte, data any) { register_bi[index] = data.(int8) },
	func(index byte, data any) { register_si[index] = data.(int16) },
	func(index byte, data any) { register_li[index] = data.(int32) },
	func(index byte, data any) { register_lli[index] = data.(int64) },
	func(index byte, data any) { register_bui[index] = data.(uint8) },
	func(index byte, data any) { register_sui[index] = data.(uint16) },
	func(index byte, data any) { register_lui[index] = data.(uint32) },
	func(index byte, data any) { register_llui[index] = data.(uint64) },
	func(index byte, data any) { register_lf[index] = data.(float32) },
	func(index byte, data any) { register_llf[index] = data.(float64) },
	func(index byte, data any) { register_s[index] = data.(string) },
	func(index byte, data any) { register_rbi[index] = data.(int8) },
	func(index byte, data any) { register_rsi[index] = data.(int16) },
	func(index byte, data any) { register_rli[index] = data.(int32) },
	func(index byte, data any) { register_rlli[index] = data.(int64) },
	func(index byte, data any) { register_rbui[index] = data.(uint8) },
	func(index byte, data any) { register_rsui[index] = data.(uint16) },
	func(index byte, data any) { register_rlui[index] = data.(uint32) },
	func(index byte, data any) { register_rllui[index] = data.(uint64) },
	func(index byte, data any) { register_rlf[index] = data.(float32) },
	func(index byte, data any) { register_rllf[index] = data.(float64) },
	func(index byte, data any) { register_rs[index] = data.(string) },
}

var register_get = [22]func(int) any{
	func(index int) any { return register_bi[index] },
	func(index int) any { return register_si[index] },
	func(index int) any { return register_li[index] },
	func(index int) any { return register_lli[index] },
	func(index int) any { return register_bui[index] },
	func(index int) any { return register_sui[index] },
	func(index int) any { return register_lui[index] },
	func(index int) any { return register_llui[index] },
	func(index int) any { return register_lf[index] },
	func(index int) any { return register_llf[index] },
	func(index int) any { return register_s[index] },
	func(index int) any { return register_rbi[index] },
	func(index int) any { return register_rsi[index] },
	func(index int) any { return register_rli[index] },
	func(index int) any { return register_rlli[index] },
	func(index int) any { return register_rbui[index] },
	func(index int) any { return register_rsui[index] },
	func(index int) any { return register_rlui[index] },
	func(index int) any { return register_rllui[index] },
	func(index int) any { return register_rlf[index] },
	func(index int) any { return register_rllf[index] },
	func(index int) any { return register_rs[index] },
}
