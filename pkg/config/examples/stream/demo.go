package main

import (
	"fmt"
	"strings"

	"github.com/shengyanli1982/toolkit/pkg/config"
)

func main() {
	// 定义测试数据
	// Define test data
	testData := `
	{
		"key1": "value1",
		"key2": "value2"
	}
	`

	// 创建一个新的 Reader 实例，用于读取字符串
	// Create a new Reader instance for reading strings
	buff := strings.NewReader(testData)

	// 创建一个新的 Content 实例
	// Create a new Content instance
	cfg := config.NewConfig().SetReader(buff).SetFileFormat(config.JSONType)
	content := config.NewStreamContent(cfg)

	// 定义预期的数据结构用于反序列化
	// Define the expected data structure for unmarshaling
	var data struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
	}

	// 调用 LoadFromStream 方法
	// Call the LoadFromStream method
	err := content.LoadFromStream(&data)
	if err != nil {
		// 如果加载文件时出现错误，打印错误并返回
		// If an error occurs while loading the file, print the error and return
		fmt.Println("Error:", err)
		return
	}

	// 打印加载的数据
	// Print the loaded data
	fmt.Println("Key1:", data.Key1, "Key2:", data.Key2)

	// 保存文件路径以供下一次测试使用
	// Save the file path for the next test
	err = content.SaveToFile()
	if err != nil {
		// 如果保存文件时出现错误，打印错误并返回
		// If an error occurs while saving the file, print the error and return
		fmt.Println("Error:", err)
		return
	}
}
