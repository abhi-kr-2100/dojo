package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dojo",
		Short: "Dojo is a workspace manager for the Jujutsu (jj) version control system.",
		Long: `Dojo is a workspace manager for the Jujutsu (jj) version control system.
It aims to eliminate every pain point associated with working on multiple workspaces.`,
	}

	cmd.AddCommand(NewLsCmd())
	cmd.AddCommand(NewNewCmd())
	cmd.AddCommand(NewCdCmd())
	cmd.AddCommand(NewRmCmd())

	return cmd
}

func Execute() error {
	return NewRootCmd().Execute()
}