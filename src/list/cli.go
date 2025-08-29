/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
func NewCommand() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List the TODO items.",
		Long:  `List the TODO items.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list called")
		},
	}
	return listCmd
}
