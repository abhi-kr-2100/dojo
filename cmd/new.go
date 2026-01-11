package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewNewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new <name>",
		Short: "Create a new workspace",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "new called with name: %s\n", args[0])
		},
	}
}