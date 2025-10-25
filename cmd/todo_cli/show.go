package todocli

import (

	"github.com/spf13/cobra"
)

var showcmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo tasks",
	Long:  "List all todo tasks",
	Run: func(cmd *cobra.Command, args []string) {
		loadTasks()

		if len(Todo_list) == 0 {
			println("Well done!! You have completed all your tasks")
		} else {
			for i, task := range Todo_list {
				println(i+1, "-", task)
			}
		}
	},
}


func init(){
	TodoCmd.AddCommand(showcmd)
}
