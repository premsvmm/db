package cmd

import (
	"fmt"
	"github.com/premsvmm/db/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go/build"
	"os"
)

var (
	conf     model.Config
	filePath string
)

const (
	folderDir = "/src/github.com/premsvmm/db/config/"
	fileName  = ".db"
	extension = "json"
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
👉 db list 
    
👉 db exec "select * from payments limit 1"

👉 db get

👉 db set <name>

👉 db delete <name>

👉 db add  -1 <jdbcdriver> -2 <host> -3 <port> -4 <dbname> -5 <dbusername> -6 <dbpassword> -7 <name> -8 <dburl>

👉 db configure -U <url> -N <basic auth username> -P <basic auth password>
`
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	//present working directory
	dir := build.Default.GOPATH
	filePath = dir + folderDir + fileName + "." + extension
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModeDir)
	}
	// Search config in home directory with name ".db" (without extension).
	viper.AddConfigPath(dir + folderDir)
	viper.SetConfigName(fileName)
	viper.SetConfigType(extension)
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		fmt.Println("Error in loading the config")
	}
	viper.Unmarshal(&conf)
}
