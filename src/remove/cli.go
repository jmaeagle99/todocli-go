/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package remove

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
func NewCommand() *cobra.Command {
	var removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove a TODO item.",
		Long:  `Remove an existing TODO item with a given ID.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remove called")
		},
	}
	return removeCmd
}
