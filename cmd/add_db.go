package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addDbCmd represents the addDb command
var addDbCmd = &cobra.Command{
	Use:   "add-db",
	Short: "Add db value to properties",
	Long:  `Add db value to properties.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addDb called")
	},
}

func init() {
	rootCmd.AddCommand(addDbCmd)
}
