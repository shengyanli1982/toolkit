package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shengyanli1982/toolkit/pkg/config"
)

func main() {
	// 创建一个临时的配置文件用于测试
	// Create a temporary config file for testing
	tmpFile, err := os.CreateTemp("", "config_test")
	if err != nil {
		// 如果创建临时文件时出现错误，打印错误并返回
		// If an error occurs while creating the temporary file, print the error and return
		fmt.Println("Error:", err)
		return
	}

	// 在函数返回时删除临时文件
	// Remove the temporary file when the function returns
	defer os.Remove(tmpFile.Name())

	// 将测试数据写入临时配置文件
	// Write test data to the temporary config file
	testData := `
	{
		"key1": "value1",
		"key2": "value2"
	}
	`
	_, err = tmpFile.Write([]byte(testData))
	if err != nil {
		// 如果写入测试数据时出现错误，打印错误并返回
		// If an error occurs while writing the test data, print the error and return
		fmt.Println("Error:", err)
		return
	}

	// 创建一个新的 Content 实例
	// Create a new Content instance
	cfg := config.NewConfig().SetSearchPaths([]string{filepath.Dir(tmpFile.Name())}).SetFileName(tmpFile.Name()).SetFileFormat(config.JSONType)
	content := config.NewContent(cfg)

	// 定义预期的数据结构用于反序列化
	// Define the expected data structure for unmarshaling
	var data struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
	}

	// 调用 LoadFromFile 方法
	// Call the LoadFromFile method
	err = content.LoadFromFile(&data)
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
