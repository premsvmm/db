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

