package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute the sql commands",
	Long: `execute the sql commands`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing the SQL Query")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(args)
	},
}
