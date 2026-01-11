package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cd <name>",
		Short: "Switch to a different workspace",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			shouldCreate, err := cmd.Flags().GetBool("new")
			if err != nil {
				cmd.PrintErrf("Error: %v\n", err)
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), "cd called with name: %s, new: %v\n", args[0], shouldCreate)
		},
	}
	cmd.Flags().Bool("new", false, "create the workspace if it doesn't exist")
	return cmd
}