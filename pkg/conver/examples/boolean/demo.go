package main

import (
	"fmt"

	"github.com/shengyanli1982/toolkit/pkg/conver"
)

func main() {
	// 调用 conver.BoolToString 函数，将布尔值 true 转换为字符串，并将结果存储在 b 中
	// Call the conver.BoolToString function to convert the boolean value true to a string, and store the result in b
	b := conver.BoolToString(true)

	// 打印 b 的值
	// Print the value of b
	fmt.Println(b)
}
