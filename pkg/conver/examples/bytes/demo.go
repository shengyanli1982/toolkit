package main

import (
	"fmt"

	"github.com/shengyanli1982/toolkit/pkg/conver"
)

func main() {
	// 调用 conver.StringToBytes 函数，将字符串 "hello" 转换为字节切片，并将结果存储在 s 中
	// Call the conver.StringToBytes function to convert the string "hello" to a byte slice, and store the result in s
	s := conver.StringToBytes("hello")

	// 打印 s 的值
	// Print the value of s
	fmt.Println(s)

	// 调用 conver.BytesToString 函数，将字节切片 s 转换回字符串，并将结果存储在 b 中
	// Call the conver.BytesToString function to convert the byte slice s back to a string, and store the result in b
	b := conver.BytesToString(s)
	
	// 打印 b 的值
	// Print the value of b
	fmt.Println(b)
}
