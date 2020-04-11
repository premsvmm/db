package cmd

import (
	"fmt"
	"github.com/premsvmm/db/service"
	"github.com/spf13/cobra"
	"strings"
)

// addDbCmd represents the addDb command
var deleteDbCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete db value to properties",
	Long:  `Delete db value to properties.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value := strings.Join(args, "")
		result, key, err := service.ValidateContextIsPresent(value, conf)
		if err != nil {
			fmt.Println("Enter value not present in db")
		}
		if result {
			conf.Database = append(conf.Database[:key], conf.Database[key+1:]...)
			service.GenerateGoFile(file_path, conf)
			fmt.Println("DB is removed")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteDbCmd)
}
