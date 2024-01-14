package command

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func CustomCobraUsage() func(*cobra.Command) error {
	return func(cmd *cobra.Command) error {
		output := cmd.OutOrStderr()

		var buf bytes.Buffer
		fmt.Fprintln(&buf, "Usage:")
		fmt.Fprintf(&buf, "\t%s\n", cmd.UseLine())

		if cmd.HasAvailableSubCommands() {
			fmt.Fprintln(&buf, "\nAvailable Commands:")
			for _, subCmd := range cmd.Commands() {
				fmt.Fprintf(&buf, "\t%-24s %s\n", subCmd.Name(), subCmd.Short)
			}
		}

		if cmd.HasAvailableFlags() {
			fmt.Fprintln(&buf, "\nFlags:")
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				fmt.Fprintf(&buf, "\t%-24s %s\n", "-"+f.Shorthand+", --"+f.Name, f.Usage)
			})
		}

		if cmd.HasExample() {
			fmt.Fprintln(&buf, "\nExamples:")
			fmt.Fprintf(&buf, "\t%s\n", cmd.Example)
		}

		_, err := fmt.Fprint(output, buf.String())
		return err
	}
}

func CustomCobraHelp() func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		ok, err := cmd.Flags().GetBool("help")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if ok {
			_ = cmd.Usage()
			os.Exit(0)
		}
	}
}

func PrettyCobraHelpAndUsage(cmd *cobra.Command) {
	cmd.SetUsageFunc(CustomCobraUsage())
	cmd.SetHelpFunc(CustomCobraHelp())
}
