package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "agent check --service <<name>>",
	Long: `*************************************************************************
	This command check enviroment variables for your MongoDB service.
*************************************************************************`,
	Run: func(cmd *cobra.Command, args []string) {

		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Println("Key:", pair[0])
			fmt.Println("Value:", pair[1])
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
