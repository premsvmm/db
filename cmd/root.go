package cmd

import (
	"fmt"
	"github.com/premsvmm/db/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	conf      model.Config
	file_path string
)

const (
	folder_dir = "/config/"
	file_name  = ".db"
	extension  = "json"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "db",
	Short: "db cli helps to connect to the stage database",
	Long: `db cli helps to connect to the stage database
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
	rootCmd.Example = `
-> db list 
    
-> db exec "select * from payments limit 1"
`
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	//present working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Search config in home directory with name ".db" (without extension).
	viper.AddConfigPath(dir + folder_dir)
	viper.SetConfigName(file_name)
	viper.SetConfigType(extension)
	file_path = dir + folder_dir + file_name + "." + extension
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error in loading the config")
	}
	viper.Unmarshal(&conf)
}
