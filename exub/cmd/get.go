package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-resty/resty"
	"github.com/spf13/cobra"
)

// getCmd represents get to print the air message
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send get request to url",
	Long:  `Send get request to url and print the request/response message`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("you need to provide url to send request")
			os.Exit(1)
		}
		getUrl(args[0])

	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
func getUrl(url string) {
	response, clientErr := resty.GetClient().Get(url)
	if clientErr != nil {
		fmt.Println(clientErr.Error())
		return
	}
	defer response.Body.Close()
	fmt.Println("------------STATUS------------------------------")
	fmt.Println(response.Status)
	fmt.Println("------------HEADER------------------------")
	response.Header.Write(os.Stdout)
	fmt.Println("------------BODY--------------------------")
	buffer := make([]byte, 1024, 1024)
	//fmt.Println(len(buffer))
	var readErr error
	var readed int
	for readed, readErr = response.Body.Read(buffer); readed > 0 && readErr == nil; {
		fmt.Println(string(buffer))
		ioutil.Discard(buffer)
	}

	fmt.Println(readErr)

	//var readed int

}
