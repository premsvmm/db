package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addDbCmd represents the addDb command
var deleteDbCmd = &cobra.Command{
	Use:   "delete-db",
	Short: "Delete db value to properties",
	Long: `Delete db value to properties.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addDb called")
	},
}
