package cmd

import (
	"fmt"
	"github.com/premsvmm/db/model"
	"github.com/premsvmm/db/service"
	"github.com/spf13/cobra"
)

var (
	add_db model.Db
)

var addDbCmd = &cobra.Command{
	Use:   "add",
	Short: "Add db value to properties",
	Long:  `Add db value to properties.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Ref :https://stackoverflow.com/questions/33892599/how-to-initialize-values-for-nested-struct-array-in-golang
		if add_db.ValidateAllValuePresent() {
			newDb := model.Db{
				JdbcDrive:  add_db.JdbcDrive,
				Host:       add_db.Host,
				Port:       add_db.Port,
				DbName:     add_db.DbName,
				DbUserName: add_db.DbUserName,
				DbPassword: add_db.DbPassword,
				URL:        add_db.URL,
				Name:       add_db.Name,
			}
			conf.Database = append(conf.Database, newDb)
			service.PrintTheDatabase(newDb)
			if service.AskForConfirmation() {
				service.GenerateGoFile(filePath, conf)
			}
		} else {
			fmt.Println("****************************************")
			fmt.Println("The above mentioned fields are required")
			fmt.Println("****************************************")
		}
	},
}

func init() {
	rootCmd.AddCommand(addDbCmd)
	addDbCmd.Flags().StringVarP(&add_db.JdbcDrive, "jdbcdriver", "1", "", "JDBC criver")
	addDbCmd.Flags().StringVarP(&add_db.Host, "host", "2", "", "Host")
	addDbCmd.Flags().StringVarP(&add_db.Port, "port", "3", "", "Port")
	addDbCmd.Flags().StringVarP(&add_db.DbName, "dbname", "4", "", "Database name")
	addDbCmd.Flags().StringVarP(&add_db.DbUserName, "dbusername", "5", "", "Database username")
	addDbCmd.Flags().StringVarP(&add_db.DbPassword, "dbpassword", "6", "", "Database password")
	addDbCmd.Flags().StringVarP(&add_db.Name, "name", "7", "", "Name")
	addDbCmd.Flags().StringVarP(&add_db.URL, "dburl", "8", "", "Datbase Url")
	addDbCmd.Example =
`
👉 db add -1 <jdbcdriver> -2 <host> -3 <port> -4 <dbname> -5 <dbusername> -6 <dbpassword> -7 <name> -8 <dburl>

Default Values:
-----------------
jdbcdriver = MYSQL_DRIVER,POSTGRESS_DRIVER

`
}
