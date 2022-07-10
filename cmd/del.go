/*
Copyright © 2022 Brochier Maxence maxence@brochier.xyz

*/
package cmd

import (
	"github.com/Maxoulfou/GoTodo/execution"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete todo",
	Long: `Delete todo
To delete an entry, do the list command just before, 
the # designates the ID of the task (the index here 
starts at 1).
Do gotodo del <id>`,
	Run: func(cmd *cobra.Command, args []string) {
		execution.DeleteTodo()
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
