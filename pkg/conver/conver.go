package conver

import (
	"strconv"
	"unsafe"
)

// BytesToString 将字节切片转换为字符串
// BytesToString converts a byte slice to a string
func BytesToString(b []byte) string {
	// 使用 unsafe.Pointer 将字节切片的地址转换为字符串的地址，然后解引用得到字符串
	// Use unsafe.Pointer to convert the address of the byte slice to the address of the string, then dereference to get the string
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes 将字符串转换为字节切片
// StringToBytes converts a string to a byte slice
func StringToBytes(s string) []byte {
	// 使用 unsafe.Pointer 将字符串的地址转换为 [2]uintptr 的地址，然后解引用得到 [2]uintptr
	// Use unsafe.Pointer to convert the address of the string to the address of [2]uintptr, then dereference to get [2]uintptr
	x := (*[2]uintptr)(unsafe.Pointer(&s))

	// 创建一个 [3]uintptr，其中前两个元素是从字符串得到的，第三个元素是字符串的末尾地址
	// Create a [3]uintptr, where the first two elements are obtained from the string, and the third element is the end address of the string
	h := [3]uintptr{x[0], x[1], x[1]}

	// 使用 unsafe.Pointer 将 [3]uintptr 的地址转换为字节切片的地址，然后解引用得到字节切片
	// Use unsafe.Pointer to convert the address of [3]uintptr to the address of the byte slice, then dereference to get the byte slice
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Uint64ToString 将无符号整数转换为字符串
// Uint64ToString converts a uint to a string
func Uint64ToString(n uint64) string {
	return strconv.FormatUint(n, 10)
}

// Int64ToString 将整数转换为字符串
// Int64ToString converts an int to a string
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

// StringToInt64 将字符串转换为整数
// StringToInt64 converts a string to an int
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// StringToUint64 将字符串转换为无符号整数
// StringToUint64 converts a string to a uint
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// BoolToString 将布尔值转换为字符串
// BoolToString converts a bool to a string
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// Float64ToString 将浮点数转换为字符串
// Float64ToString converts a float64 to a string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToFloat64 将字符串转换为浮点数
// StringToFloat64 converts a string to a float64
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
