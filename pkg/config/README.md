# Config

**config** is a package that provides a simple and flexible way to manage configuration settings for your application. It is designed to be easy to use and compatible with various configuration sources.

`config` supports two modes: `file` and `stream`. You can use `config` to easily read configuration from a file or a stream.

-   **File** : Read configuration from a file.
-   **Stream** : Read configuration from a stream.

## Installation

```bash
go get github.com/shengyanli1982/toolkit/pkg/config
```

## Quick Start

### Config

The `config` package provides a config object that allows you to configure the behavior of the batch process. The config object supports the following methods:

-   `SetSearchPaths`: Set the search paths for the configuration file.
-   `SetFileName`: Set the file name for the configuration file.
-   `SetFileFormat`: Set the file format for the configuration file.
-   `SetReader`: Set the reader for the configuration file. This method is only supported in `stream` mode.

### Components

#### 1. **File** : Read configuration from a file.

The most commonly used component in `config` is `File`. The `File` component provides the following methods for configuration:

**Methods**

-   `LoadFromFile`: Load configuration from a file.
-   `SaveToFile`: Save configuration to a file.
-   `SaveToFileWithName`: Save configuration to a file with a specific name.
-   `GetViper`: Get the viper object.

**Example**

```go
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
```

**Result**

```bash
$ go run demo.go
Key1: value1 Key2: value2
```

#### 2. **Stream** : Read configuration from IO stream.

The `config` package allows you to read configuration from a stream and save it to a file. The `Stream` component provides the following methods for configuration:

**Methods**

-   `LoadFromStream`: Load configuration from a stream.
-   `SaveToFile`: Save configuration to a stream.
-   `SaveToFileWithName`: Save configuration to a file with a specific name.
-   `GetViper`: Get the viper object.

**Example**

```go
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
```

**Result**

```bash
$ go run demo.go
Key1: value1 Key2: value2
```
