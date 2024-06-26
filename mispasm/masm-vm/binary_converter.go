package main

func convert_to_value(data []byte) any {
	switch data[0] {
	case INT8:
		return byte_to_int8(data[1])
	case INT16:
		return bytes_to_int16(data[1:])
	case INT32:
		return bytes_to_int32(data[1:])
	case INT64:
		return bytes_to_int64(data[1:])
	case UINT8:
		return byte_to_uint8(data[1])
	case UINT16:
		return bytes_to_uint16(data[1:])
	case UINT32:
		return bytes_to_uint32(data[1:])
	case UINT64:
		return bytes_to_uint64(data[1:])
	case FLOAT32:
		return bytes_to_float32(data[1:])
	case FLOAT64:
		return bytes_to_float64(data[1:])
	case BOOL:
		return byte_to_bool(data[1])
	case STRING:
		return bytes_to_string(data[1:])
	case REG:
		return register_get[data[1]](int(data[2]))
	}
	return nil
}

func convert_to_bytes(t byte, data any) []byte {
	bytes := []byte{}
	switch t {
	case INT8:
		bytes = []byte{int8_to_byte(data.(int8))}
	case INT16:
		bytes = int16_to_bytes(data.(int16))
	case INT32:
		bytes = int32_to_bytes(data.(int32))
	case INT64:
		bytes = int64_to_bytes(data.(int64))
	case UINT8:
		bytes = []byte{uint8_to_byte(data.(uint8))}
	case UINT16:
		bytes = uint16_to_bytes(data.(uint16))
	case UINT32:
		bytes = uint32_to_bytes(data.(uint32))
	case UINT64:
		bytes = uint64_to_bytes(data.(uint64))
	case FLOAT32:
		bytes = float32_to_bytes(data.(float32))
	case FLOAT64:
		bytes = float64_to_bytes(data.(float64))
	case BOOL:
		bytes = []byte{bool_to_byte(data.(bool))}
	case STRING:
		bytes = string_to_bytes(data.(string))
	}
	return append([]byte{t}, bytes...)
}

func int8_to_byte(i int8) byte {
	return byte(i)
}

func byte_to_int8(b byte) int8 {
	return int8(b)
}

func int16_to_bytes(i int16) []byte {
	return []byte{byte(i >> 8), byte(i)}
}

func bytes_to_int16(b []byte) int16 {
	return int16(b[0])<<8 | int16(b[1])
}

func int32_to_bytes(i int32) []byte {
	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func bytes_to_int32(b []byte) int32 {
	return int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])
}

func int64_to_bytes(i int64) []byte {
	return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func bytes_to_int64(b []byte) int64 {
	return int64(b[0])<<56 | int64(b[1])<<48 | int64(b[2])<<40 | int64(b[3])<<32 | int64(b[4])<<24 | int64(b[5])<<16 | int64(b[6])<<8 | int64(b[7])
}

func uint8_to_byte(i uint8) byte {
	return byte(i)
}

func byte_to_uint8(b byte) uint8 {
	return uint8(b)
}

func uint16_to_bytes(i uint16) []byte {
	return []byte{byte(i >> 8), byte(i)}
}

func bytes_to_uint16(b []byte) uint16 {
	return uint16(b[0])<<8 | uint16(b[1])
}

func uint32_to_bytes(i uint32) []byte {
	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func bytes_to_uint32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func uint64_to_bytes(i uint64) []byte {
	return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func bytes_to_uint64(b []byte) uint64 {
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}

func float32_to_bytes(f float32) []byte {
	return int32_to_bytes(int32(f))
}

func bytes_to_float32(b []byte) float32 {
	return float32(bytes_to_int32(b))
}

func float64_to_bytes(f float64) []byte {
	return int64_to_bytes(int64(f))
}

func bytes_to_float64(b []byte) float64 {
	return float64(bytes_to_int64(b))
}

func bool_to_byte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func byte_to_bool(b byte) bool {
	return b == 1
}

func string_to_bytes(s string) []byte {
	return []byte(s)
}

func bytes_to_string(b []byte) string {
	return string(b)
}
