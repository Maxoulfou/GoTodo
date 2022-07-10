/*
Copyright © 2022 Brochier Maxence maxence@brochier.xyz

*/
package cmd

import (
	"github.com/Maxoulfou/GoTodo/execution"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add a new todo
To add an entry,
Do gotodo add <text to describe your todo>`,
	Run: func(cmd *cobra.Command, args []string) {
		execution.AddTodo()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
