package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "agent check --service <<name>>",
	Long: `*************************************************************************
This command check enviroment variables for your MongoDB service.
*************************************************************************`,
	Run: func(cmd *cobra.Command, args []string) {
		checkMongoDB()
	},
}

// MONGODB_ADDRESS: << driver://hostname:port >>
// MONGODB_DATABASE: << database >>
// MONGODB_USER: << username >>
// MONGODB_PASSWORD: << password >>

func checkMongoDB() {

	for _, e := range viper.AllKeys() {
		keyvar := strings.ToUpper(e)
		resp := checkEnv(keyvar)
		fmt.Println(resp)
	}

}

func checkEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		os.Setenv(key, value)
		fmt.Println(key, value)
	}
	return value
}

func init() {

	checkCmd.Flags().BoolP("mongodb", "m", true, "Check enviroment variables for MongoDB")
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
