package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure help to configure the mock url and credentials",
	Long: `Configure help to configure the mock url and credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
	},
}
