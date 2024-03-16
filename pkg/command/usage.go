package command

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// CustomCobraUsage 是一个自定义的 Cobra 命令行使用说明生成器
// CustomCobraUsage is a custom Cobra command line usage generator
func CustomCobraUsage() func(*cobra.Command) error {
	return func(cmd *cobra.Command) error {
		// 获取命令行输出对象
		// Get the command line output object
		output := cmd.OutOrStderr()

		// 创建一个缓冲区来存储使用说明
		// Create a buffer to store the usage
		var buf bytes.Buffer

		// 输出 "Usage:" 字样
		// Output the word "Usage:"
		fmt.Fprintln(&buf, "Usage:")

		// 输出命令行的使用方法
		// Output the usage of the command line
		fmt.Fprintf(&buf, "\t%s\n", cmd.UseLine())

		// 如果有可用的子命令，输出子命令列表
		// If there are available subcommands, output the subcommand list
		if cmd.HasAvailableSubCommands() {
			// 输出 "Available Commands:" 字样
			// Output the word "Available Commands:"
			fmt.Fprintln(&buf, "\nAvailable Commands:")

			// 遍历所有子命令
			// Iterate through all subcommands
			for _, subCmd := range cmd.Commands() {
				// 输出子命令的名称和简短描述
				// Output the name and short description of the subcommand
				fmt.Fprintf(&buf, "\t%-24s %s\n", subCmd.Name(), subCmd.Short)
			}
		}

		// 如果有可用的标志，输出标志列表
		// If there are available flags, output the flag list
		if cmd.HasAvailableFlags() {
			// 输出 "Flags:" 字样
			// Output the word "Flags:"
			fmt.Fprintln(&buf, "\nFlags:")

			// 遍历所有标志
			// Iterate through all flags
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				// 输出标志的短名称、长名称和使用说明
				// Output the short name, long name, and usage of the flag
				fmt.Fprintf(&buf, "\t%-24s %s\n", "-"+f.Shorthand+", --"+f.Name, f.Usage)
			})
		}

		// 如果有示例，输出示例
		// If there are examples, output the examples
		if cmd.HasExample() {
			fmt.Fprintln(&buf, "\nExamples:")
			fmt.Fprintf(&buf, "\t%s\n", cmd.Example)
		}

		// 输出缓冲区的内容，并返回可能的错误
		// Output the content of the buffer and return possible errors
		_, err := fmt.Fprint(output, buf.String())
		return err
	}
}

// CustomCobraHelp 是一个自定义的 Cobra 命令行帮助信息生成器
// CustomCobraHelp is a custom Cobra command line help message generator
func CustomCobraHelp() func(*cobra.Command, []string) {
	// 返回一个函数，该函数接受一个命令和一个字符串切片作为参数
	// Return a function that takes a command and a slice of strings as arguments
	return func(cmd *cobra.Command, args []string) {
		// 获取 "help" 标志的值
		// Get the value of the "help" flag
		ok, err := cmd.Flags().GetBool("help")

		// 如果获取失败，输出错误信息并退出程序
		// If the acquisition fails, output the error message and exit the program
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// 如果 "help" 标志的值为 true，输出命令的使用说明并退出程序
		// If the value of the "help" flag is true, output the usage of the command and exit the program
		if ok {
			_ = cmd.Usage()
			os.Exit(0)
		}
	}
}

// PrettyCobraHelpAndUsage 设置美化后的 Cobra 命令行帮助信息和使用说明
// PrettyCobraHelpAndUsage sets the beautified Cobra command line help message and usage
func PrettyCobraHelpAndUsage(cmd *cobra.Command) {
	// 设置使用说明生成器为自定义的生成器
	// Set the usage generator to the custom generator
	cmd.SetUsageFunc(CustomCobraUsage())

	// 设置帮助信息生成器为自定义的生成器
	// Set the help message generator to the custom generator
	cmd.SetHelpFunc(CustomCobraHelp())
}
