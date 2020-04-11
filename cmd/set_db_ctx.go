package cmd

import (
	"fmt"
	"github.com/premsvmm/db/model"
	"github.com/premsvmm/db/service"
	"github.com/spf13/cobra"
	"strings"
)

// setDbCtxCmd represents the setDbCtx command
var setDbCtxCmd = &cobra.Command{
	Use:   "ctx",
	Short: "Set the current db context to access",
	Long:  `Detail description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value := strings.Join(args, "")
		if value != "" {
			conf.CurrentDbContext = value
			if ValidateContextIsPresent(value, conf) {
				service.GenerateGoFile(file_path, conf)
				fmt.Println("🔥 Db context is set")
			} else {
				fmt.Println("❌ db context is not present")
				service.ListDownTheDBName(conf)
			}
		}
	},
}

func ValidateContextIsPresent(value string, conf model.Config) bool {
	for key, _ := range conf.Database {
		if conf.Database[key].Name == value {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(setDbCtxCmd)
}
