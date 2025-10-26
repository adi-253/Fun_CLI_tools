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
		if len(args)==0{
			return fmt.Errorf("Please Provide a Valid url")
		}
		url:=strings.Join(args," ")
		resp,err := http.Get(url)
		if err!=nil{
			return err
		}
		defer resp.Body.Close() // close the connection
		body,err:= io.ReadAll(resp.Body)  // see you may have learnt that the data comes in chunks and not all together so its like stream and hence wait to readall 
		if err!=nil{
			return err
		}
		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Println(string(body))
		return nil
	},

}

func init(){
	Apicmd.AddCommand(getcmd)
}