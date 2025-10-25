package todocli

import (
	"github.com/spf13/cobra"
)

var Todo_list []string


var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Base Command for todo",
	Long:  "So basically this is the base command for todo; subcommands will be added to this command",
}