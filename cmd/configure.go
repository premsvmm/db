package cmd

import (
	"fmt"
	"github.com/premsvmm/db/service"
	"github.com/spf13/cobra"
	"time"
)

var (
	service_url string
	return_type string
	username    string
	password    string
	reset       string
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure the mock url and credentials",
	Long:  `configure the mock url and credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(reset) > 0 {
			conf.Password = ""
			conf.URL = ""
			conf.ReturnType = ""
			conf.Username = ""
			time.Sleep(1000 * time.Millisecond)
			service.GenerateGoFile(file_path,conf)
			fmt.Println("✔ Values are reseted, Please set the value again")
		} else {
			if service_url != "" {
				conf.URL = service_url
				fmt.Println("✔ Service url is set")
				time.Sleep(1000 * time.Millisecond)
			}
			if return_type != "" {
				conf.ReturnType = return_type
				fmt.Println("✔ Return type is set")
				time.Sleep(1000 * time.Millisecond)
			}
			if username != "" {
				conf.Username = username
				fmt.Println("✔ Username is set")
				time.Sleep(1000 * time.Millisecond)
			}
			if password != "" {
				conf.Password = password
				fmt.Println("✔ Password is set")
				time.Sleep(1000 * time.Millisecond)
			}
			service.GenerateGoFile(file_path,conf)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.Flags().StringVarP(&service_url, "url", "U", "", "Service Url")
	configureCmd.Flags().StringVarP(&return_type, "rtype", "R", "", "Return type (json or map)")
	configureCmd.Flags().StringVarP(&username, "username", "N", "", "Username")
	configureCmd.Flags().StringVarP(&password, "password", "P", "", "Password")
	configureCmd.Flags().StringVarP(&reset, "reset", "r", "", "Reset")
}
