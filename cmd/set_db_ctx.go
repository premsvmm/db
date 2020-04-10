package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setDbCtxCmd represents the setDbCtx command
var setDbCtxCmd = &cobra.Command{
	Use:   "set-db-ctx",
	Short: "Set the current db context to access",
	Long: `Detail description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setDbCtx called")
	},
}

func init() {
	rootCmd.AddCommand(setDbCtxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setDbCtxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setDbCtxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
