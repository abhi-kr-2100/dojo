package cmd

import (
	"github.com/spf13/cobra"
)

func NewLsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List all managed workspaces",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte("ls called\n"))
		},
	}
}