package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure the mock url and credentials",
	Long: `configure the mock url and credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
	},
}
