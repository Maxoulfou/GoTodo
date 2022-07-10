/*
Copyright © 2022 Brochier Maxence maxence@brochier.xyz

*/
package cmd

import (
	"github.com/Maxoulfou/GoTodo/execution"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long: `List all todos
To list all todos, 
Do gotodo list`,
	Run: func(cmd *cobra.Command, args []string) {
		execution.ListTodo()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
