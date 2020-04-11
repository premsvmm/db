package cmd

import (
	"github.com/premsvmm/db/service"
	"github.com/spf13/cobra"
	"strings"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute the sql commands",
	Long: `execute the sql commands`,
	Run: func(cmd *cobra.Command, args []string) {
		str:=strings.Join(args," ")
		service.ExecuteMock(str,conf)
	},
}
