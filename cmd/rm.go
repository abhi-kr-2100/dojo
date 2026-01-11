package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRmCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rm [<name>]",
		Short: "Remove a workspace",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "rm called with name: %s\n", args[0])
			} else {
				fmt.Fprintln(cmd.OutOrStdout(), "rm called for current workspace")
			}
		},
	}
}