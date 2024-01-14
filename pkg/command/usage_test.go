package command

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCustomCobraUsage(t *testing.T) {
	// Create a new Cobra command
	cmd := &cobra.Command{
		Use:     "mycommand",
		Short:   "My command",
		Example: "mycommand --flag1 value1",
	}

	// Add subcommands
	subCmd1 := &cobra.Command{
		Use:   "subcommand1",
		Short: "Subcommand 1",
	}
	subCmd2 := &cobra.Command{
		Use:   "subcommand2",
		Short: "Subcommand 2",
	}
	cmd.AddCommand(subCmd1, subCmd2)

	// Add flags
	cmd.Flags().StringP("flag1", "f", "", "Flag 1")
	cmd.Flags().StringP("flag2", "g", "", "Flag 2")

	// Set the output buffer
	buf := bytes.Buffer{}
	cmd.SetOut(&buf)

	// Call the CustomCobraUsage function
	err := CustomCobraUsage()(cmd)
	assert.NoError(t, err)

	// Verify the output
	expectedOutput := `Usage:
	mycommand [flags]

Flags:
	-f, --flag1              Flag 1
	-g, --flag2              Flag 2

Examples:
	mycommand --flag1 value1
`
	assert.Equal(t, expectedOutput, buf.String())
}
