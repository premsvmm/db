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
				fmt.Println("🔥 DB : " + conf.CurrentDbContext + " is set")
			} else {
				fmt.Println("❌ db is not present in properties")
				fmt.Println()
				service.ListDownTheDBName(conf)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setDbCtxCmd)
	setDbCtxCmd.Example=
`
db set <dbname> - (set the db which the query need to execute)
`
}
