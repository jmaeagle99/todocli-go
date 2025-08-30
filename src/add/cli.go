/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"

	"github.com/jmaeagle99/todocli/store"
	"github.com/jmaeagle99/todocli/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
func NewCommand() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a TODO item.",
		Long:  `Add a TODO item with a name.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("add called")
			store, err := store.CreateStore[todo.Todo]()
			if nil != err {
				return err
			}
			return store.AddItem(todo.Todo{Name: args[0]})
		},
	}
	return addCmd
}
