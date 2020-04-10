package cmd

import (
	"fmt"
	"github.com/premsvmm/db/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"

)
var (
	conf model.Config
)
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "db",
	Short: "db cli helps to connect to the mock server",
	Long:
	`db cli helps to connect to the mock server
********************************************		
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(InitConfig)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(setDbCtxCmd)
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(addDbCmd)
	rootCmd.AddCommand(deleteDbCmd)
	rootCmd.AddCommand(listCmd)
}

// initConfig reads in config file and ENV variables if set.
func InitConfig(){
	//present working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Search config in home directory with name ".db" (without extension).
	viper.AddConfigPath(dir + "/config/")
	viper.SetConfigName(".db")
	viper.SetConfigType("json")
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error in loading the config")
	}
	viper.Unmarshal(&conf)
}
