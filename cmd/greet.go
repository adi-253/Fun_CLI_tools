package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var greetcmd = &cobra.Command{
	Use: "greet [name]",
	Short: "greets the user",
	RunE: func(cmd *cobra.Command,args []string) error {
			name := "user"
			if len(args) > 0 {
        		name = strings.Join(args, " ")
    		}
			fmt.Printf("Hello %s!\n",name)
			return nil
	},
}
func init() {
	rootCmd.AddCommand(greetcmd)
}