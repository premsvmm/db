package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// getDbCtxCmd represents the getDbCtxCmd command
var getDbCtxCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Current Db Context",
	Long:  "Get the current Db Context \nUsage : \n-> db get",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current Db Context : " + conf.CurrentDbContext)
	},
}

func init() {
	rootCmd.AddCommand(getDbCtxCmd)
}
