/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
func NewCommand() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a TODO item.",
		Long:  `Add a TODO item with a name.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add called")
		},
	}
	return addCmd
}
