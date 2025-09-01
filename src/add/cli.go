/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"errors"

	"github.com/jmaeagle99/todocli/store"
	"github.com/jmaeagle99/todocli/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
func NewCommand() *cobra.Command {
	var name string

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a TODO item.",
		Long:  `Add a TODO item with a name.`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			store, err := store.CreateStore[todo.Todo]()
			if nil != err {
				return err
			}
			defer func() {
				err = errors.Join(err, store.Close())
			}()

			_, err = store.Add(todo.Todo{Name: name})

			return
		},
	}

	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of TODO item.")
	addCmd.MarkFlagRequired("name")

	return addCmd
}
