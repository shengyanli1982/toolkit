package main

import (
	"fmt"

	"github.com/shengyanli1982/toolkit/pkg/conver"
)

func main() {
	// 调用 conver.Int64ToString 函数，将 int64 类型的 1 转换为字符串，并打印结果
	// Call the conver.Int64ToString function to convert the int64 type 1 to a string, and print the result
	fmt.Println(conver.Int64ToString(1))

	// 调用 conver.StringToInt64 函数，将字符串 "1" 转换为 int64 类型，并打印结果
	// Call the conver.StringToInt64 function to convert the string "1" to int64 type, and print the result
	fmt.Println(conver.StringToInt64("1"))

	// 调用 conver.Float64ToString 函数，将 float64 类型的 1.0 转换为字符串，并打印结果
	// Call the conver.Float64ToString function to convert the float64 type 1.0 to a string, and print the result
	fmt.Println(conver.Float64ToString(1.0))

	// 调用 conver.StringToFloat64 函数，将字符串 "1.0" 转换为 float64 类型，并打印结果
	// Call the conver.StringToFloat64 function to convert the string "1.0" to float64 type, and print the result
	fmt.Println(conver.StringToFloat64("1.0"))
}
