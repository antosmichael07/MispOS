package main

type Type int

const (
	INT8 Type = iota
	INT16
	INT32
	INT64
	UINT8
	UINT16
	UINT32
	UINT64
	FLOAT32
	FLOAT64
	STRING
	BOOL
)

func convert_to_value(data []byte) any {
	switch data[0] {
	case byte(INT8):
		return ByteToInt8(data[1])
	case byte(INT16):
		return BytesToInt16(data[1:])
	case byte(INT32):
		return BytesToInt32(data[1:])
	case byte(INT64):
		return BytesToInt64(data[1:])
	case byte(UINT8):
		return ByteToUint8(data[1])
	case byte(UINT16):
		return BytesToUint16(data[1:])
	case byte(UINT32):
		return BytesToUint32(data[1:])
	case byte(UINT64):
		return BytesToUint64(data[1:])
	case byte(FLOAT32):
		return BytesToFloat32(data[1:])
	case byte(FLOAT64):
		return BytesToFloat64(data[1:])
	case byte(STRING):
		return BytesToString(data[1:])
	case byte(BOOL):
		return ByteToBool(data[1])
	}
	return nil
}

func convert_to_bytes(t Type, data any) []byte {
	bytes := []byte{}
	switch t {
	case INT8:
		bytes = []byte{Int8ToByte(data.(int8))}
	case INT16:
		bytes = Int16ToBytes(data.(int16))
	case INT32:
		bytes = Int32ToBytes(data.(int32))
	case INT64:
		bytes = Int64ToBytes(data.(int64))
	case UINT8:
		bytes = []byte{Uint8ToByte(data.(uint8))}
	case UINT16:
		bytes = Uint16ToBytes(data.(uint16))
	case UINT32:
		bytes = Uint32ToBytes(data.(uint32))
	case UINT64:
		bytes = Uint64ToBytes(data.(uint64))
	case FLOAT32:
		bytes = Float32ToBytes(data.(float32))
	case FLOAT64:
		bytes = Float64ToBytes(data.(float64))
	case STRING:
		bytes = StringToBytes(data.(string))
	case BOOL:
		bytes = []byte{BoolToByte(data.(bool))}
	}
	return append([]byte{byte(t)}, bytes...)
}

func Int8ToByte(i int8) byte {
	return byte(i)
}

func ByteToInt8(b byte) int8 {
	return int8(b)
}

func Int16ToBytes(i int16) []byte {
	return []byte{byte(i >> 8), byte(i)}
}

func BytesToInt16(b []byte) int16 {
	return int16(b[0])<<8 | int16(b[1])
}

func Int32ToBytes(i int32) []byte {
	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func BytesToInt32(b []byte) int32 {
	return int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])
}

func Int64ToBytes(i int64) []byte {
	return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func BytesToInt64(b []byte) int64 {
	return int64(b[0])<<56 | int64(b[1])<<48 | int64(b[2])<<40 | int64(b[3])<<32 | int64(b[4])<<24 | int64(b[5])<<16 | int64(b[6])<<8 | int64(b[7])
}

func Uint8ToByte(i uint8) byte {
	return byte(i)
}

func ByteToUint8(b byte) uint8 {
	return uint8(b)
}

func Uint16ToBytes(i uint16) []byte {
	return []byte{byte(i >> 8), byte(i)}
}

func BytesToUint16(b []byte) uint16 {
	return uint16(b[0])<<8 | uint16(b[1])
}

func Uint32ToBytes(i uint32) []byte {
	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func BytesToUint32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func Uint64ToBytes(i uint64) []byte {
	return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func BytesToUint64(b []byte) uint64 {
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}

func Float32ToBytes(f float32) []byte {
	return Int32ToBytes(int32(f))
}

func BytesToFloat32(b []byte) float32 {
	return float32(BytesToInt32(b))
}

func Float64ToBytes(f float64) []byte {
	return Int64ToBytes(int64(f))
}

func BytesToFloat64(b []byte) float64 {
	return float64(BytesToInt64(b))
}

func StringToBytes(s string) []byte {
	return []byte(s)
}

func BytesToString(b []byte) string {
	return string(b)
}

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func ByteToBool(b byte) bool {
	return b == 1
}
