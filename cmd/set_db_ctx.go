package cmd

import (
	"fmt"
	"github.com/premsvmm/db/service"
	"github.com/spf13/cobra"
	"strings"
)

// setDbCtxCmd represents the setDbCtx command
var setDbCtxCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the current db context to access",
	Long:  `Detail description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value := strings.Join(args, "")
		if value != "" {
			conf.CurrentDbContext = value
			result, _, _ := service.ValidateContextIsPresent(value, conf)
			if result {
				service.GenerateGoFile(file_path, conf)
				fmt.Println("🔥 Db context is set")
			} else {
				fmt.Println("❌ db context is not present")
				service.ListDownTheDBName(conf)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setDbCtxCmd)
}
