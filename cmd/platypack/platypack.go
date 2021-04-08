package main

import "github.com/spf13/cobra"

func NewPlatypackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "platypack",
		Short: "create and manage pulumi packages",
		Long:  `platypack is a CLI for creating and managing pulumi packages.`,
	}
	cmd.AddCommand(newNewCommand())
	return cmd
}
