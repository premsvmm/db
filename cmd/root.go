package cmd

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/premsvmm/db/model"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go/build"
	"log"
	"os"
)

var (
	conf      model.Config
	file_path string
)

const (
	folder_dir = "/src/github.com/premsvmm/db/config/"
	file_name  = ".db"
	extension  = "json"
	version    = "v1.0.2"
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
1 db list 
    
2 db exec "select * from payments limit 1"

3 db get

4 db set <name>

5 db delete <name>

6 db add  -1 <jdbcdriver> -2 <host> -3 <port> -4 <dbname> -5 <dbusername> -6 <dbpassword> -7 <name> -8 <dburl>

7 db configure -U <url> -N <basic auth username> -P <basic auth password>
`
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	doSelfUpdate()
	//present working directory
	dir := build.Default.GOPATH
	file_path = dir + folder_dir + file_name + "." + extension
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		os.Mkdir(file_path, os.ModeDir)
	}
	// Search config in home directory with name ".db" (without extension).
	viper.AddConfigPath(dir + folder_dir)
	viper.SetConfigName(file_name)
	viper.SetConfigType(extension)
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		fmt.Println("Error in loading the config")
	}
	viper.Unmarshal(&conf)
}

func doSelfUpdate() {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "premsvmm/db")
	if err != nil {
		log.Println("Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		log.Println("Current binary is the latest version", version)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
}
