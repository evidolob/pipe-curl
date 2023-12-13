package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	pipePath        string
	httpMethod      string
	includeProtocol bool
)

var rootCmd = &cobra.Command{
	Use:   "pipe-curl",
	Short: "pipe-curl is a CLI to send HTTP Request thru windows named-pipe",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		handleCommand(args)

	},
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&pipePath, "pipe", "p", "", "Path to Windows named pipe to interact with")
	rootCmd.MarkPersistentFlagRequired("pipe")

	rootCmd.PersistentFlags().StringVarP(&httpMethod, "request", "X", "", "Change the method to use when starting the transfer")
	rootCmd.PersistentFlags().BoolVarP(&includeProtocol, "include", "i", false, "Include protocol response headers in the output")
}

func handleCommand(args []string) {
	if len(args) != 1 {
		err := fmt.Errorf("only one argument is supported, got %s", args)
		printErrorAndExit(err)
	}

	host := args[0]

	err := makeRequest(httpMethod, host, pipePath, includeProtocol)
	if err != nil {
		printErrorAndExit(err)
	}
}

func printErrorAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
