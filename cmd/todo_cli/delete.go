package todocli

import (
	"strings"
	"fmt"
	"github.com/spf13/cobra"
)

var deletecmd = &cobra.Command{
	Use:   "delete [task]",
	Short: "Add a task to the todo tasks",
	Long:  "Add a task to the todo tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			println("Please provide a task to delete")
			return
		}

		// Load existing tasks
		loadTasks()

		target := strings.Join(args, " ")
		if len(Todo_list)==0{
			fmt.Println("Todo Task is empty")
		}else{
			found:=false
			for i,task:= range Todo_list{
				if task==target{
					found=true
					Todo_list = append(Todo_list[:i],Todo_list[i+1:]...)
					break
				}
			}
			if !found{
				fmt.Println("task is not present in the list")
			}
			if err:=saveTasks(); err!=nil{
				fmt.Println("error saving tasks",err.Error())
			}
		}
	},
}

func init(){
	TodoCmd.AddCommand(deletecmd)
}
