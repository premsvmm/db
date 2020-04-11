package cmd

import (
	"github.com/premsvmm/db/service"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the db",
	Long:  `List the db`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ListDownTheDBName(conf)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Example= `👉 db list - (display all db added to properties file)`
}
