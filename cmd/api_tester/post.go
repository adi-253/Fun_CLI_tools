package apitester

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var (
    postData string
    contentType string
)

var postcmd = &cobra.Command{
    Use: "POST [URL]",
    Short: "POST method",
    Long: "Send a POST request to the specified URL with optional data and content type",
    RunE: func(cmd *cobra.Command, args []string) error {
        if len(args) == 0 {
            return fmt.Errorf("please provide a URL")
        }
        
        url := args[0]
        
        // Create request body
        var body io.Reader
        if postData != "" {
            body = bytes.NewBufferString(postData)
        }
        
        // Create POST request
        req, err := http.NewRequest("POST", url, body)
        if err != nil {
            return err
        }
        
        // Set content type if provided
        if contentType != "" {
            req.Header.Set("Content-Type", contentType)
        }
        
        // Send request
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        
        // Read response
        respBody, err := io.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        
        fmt.Printf("Status: %s\n", resp.Status)
        fmt.Println(string(respBody))
        return nil
    },
}

func init() {
    postcmd.Flags().StringVarP(&postData, "data", "d", "", "Data to send in POST request body")
    postcmd.Flags().StringVarP(&contentType, "content-type", "c", "application/json", "Content-Type header")
    Apicmd.AddCommand(postcmd)
}