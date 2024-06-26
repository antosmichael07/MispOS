package main

const (
	INT8 byte = iota
	INT16
	INT32
	INT64
	UINT8
	UINT16
	UINT32
	UINT64
	FLOAT32
	FLOAT64
	BOOL
	STRING
	REG
)

var type_sizes = []byte{1, 2, 4, 8, 1, 2, 4, 8, 4, 8, 1, 0, 2}
