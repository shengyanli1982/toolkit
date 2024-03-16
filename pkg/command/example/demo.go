package main

import (
	"fmt"

	"github.com/shengyanli1982/toolkit/pkg/command"
	"github.com/spf13/cobra"
)

func main() {
	// 创建 cobra 命令
	// Create a cobra command
	cmd := cobra.Command{Use: "demo"}

	// 设置命令行使用说明生成器
	// Set the command line usage instruction generator
	command.PrettyCobraHelpAndUsage(&cmd)

	// 执行 cobra 命令
	// Execute the cobra command
	if err := cmd.Execute(); err != nil {
		// 如果执行命令时出现错误，打印错误并返回
		// If an error occurs while executing the command, print the error and return
		fmt.Println("Error:", err)
		return
	}
}
