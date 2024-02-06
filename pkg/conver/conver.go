package conver

import (
	"strconv"
	"unsafe"
)

// BytesToString converts a byte slice to a string.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts a string to a byte slice.
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Uint64ToString converts a uint to a string.
func Uint64ToString(n uint64) string {
	return strconv.FormatUint(n, 10)
}

// Int64ToString converts an int to a string.
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

// StringToInt64 converts a string to an int.
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// StringToUint64 converts a string to a uint.
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// BoolToString converts a bool to a string.
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// Float64ToString converts a float64 to a string.
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToFloat64 converts a string to a float64.
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
