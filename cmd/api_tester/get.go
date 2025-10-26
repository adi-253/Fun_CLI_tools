package apitester

import (
	"github.com/spf13/cobra"
	"strings"
	"fmt"
	"io"
	"net/http"
) 


var getcmd = &cobra.Command{
	Use: "GET [URL]",
	Short: "Get method",
	RunE: func(cmd *cobra.Command, args[] string) error {
		url:=strings.Join(args," ")
		resp,err := http.Get(url)
		if err!=nil{
			return err
		}
		defer resp.Body.Close()
		body,err:= io.ReadAll(resp.Body)
		if err!=nil{
			return err
		}
		fmt.Println(string(body))
		return nil
	},

}

func init(){
	Apicmd.AddCommand(getcmd)
}