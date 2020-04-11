package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// getDbCtxCmd represents the getDbCtxCmd command
var getDbCtxCmd = &cobra.Command{
	Use:   "gctx",
	Short: "get the current db context to access",
	Long:  `Detail description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current Db Context : " + conf.CurrentDbContext)
	},
}

func init() {
	rootCmd.AddCommand(getDbCtxCmd)
}
