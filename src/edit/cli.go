/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package edit

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
func NewCommand() *cobra.Command {
	var editCmd = &cobra.Command{
		Use:   "edit",
		Short: "Update an existing TODO item.",
		Long:  `Update an existing TODO item with a given ID.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("edit called")
		},
	}
	return editCmd
}
