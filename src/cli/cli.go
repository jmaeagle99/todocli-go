/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"github.com/jmaeagle99/todocli/add"
	"github.com/jmaeagle99/todocli/edit"
	"github.com/jmaeagle99/todocli/list"
	"github.com/jmaeagle99/todocli/remove"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "todocli",
		Short: "A CLI for managing TODO items.",
		Long:  `A CLI for managing TODO items.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	rootCmd.AddCommand(add.NewCommand())
	rootCmd.AddCommand(edit.NewCommand())
	rootCmd.AddCommand(list.NewCommand())
	rootCmd.AddCommand(remove.NewCommand())

	return rootCmd
}
