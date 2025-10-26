package apitester


import (
	"github.com/spf13/cobra"
)

var Apicmd = &cobra.Command{
	Use: "apitest",
	Short: "Use for testing api calls",
	Long: "Basically like a curl command",
}