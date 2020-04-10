package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the db",
	Long:  `List the db`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`List of Database added in config file:
**************************************`)
		for key, _ := range conf.Database {
			fmt.Println(conf.Database[key].Name)
		}
		fmt.Println("**************************************")
	},
}
