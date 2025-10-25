package todocli

import (
	"strings"

	"github.com/spf13/cobra"
)

var addcmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a task to the todo tasks",
	Long:  "Add a task to the todo tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			println("Please provide a task to add")
			return
		}

		// Load existing tasks
		loadTasks()

		task := strings.Join(args, " ") // handle spaces in task
		Todo_list = append(Todo_list, task)

		// Save updated tasks
		if err := saveTasks(); err != nil {
			println("Error saving tasks:", err.Error())
		} else {
			println("Task added:", task)
		}
	},
}

func init(){
	// attach this subcommand to the exported TodoCmd
	TodoCmd.AddCommand(addcmd)
}
