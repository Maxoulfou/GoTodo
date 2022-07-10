/*
Copyright © 2022 Brochier Maxence maxence@brochier.xyz

*/
package cmd

import (
	"github.com/Maxoulfou/GoTodo/execution"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark todo as completed",
	Long: `Mark todo as completed
To mark todo as complete , do the list command just 
before, the # designates the ID of the task (the index
here starts at 1).
Do $_ gotodo complete <id>`,
	Run: func(cmd *cobra.Command, args []string) {
		execution.CompleteTodo()
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
