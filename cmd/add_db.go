package cmd

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/premsvmm/db/model"
	"github.com/premsvmm/db/service"
	"github.com/spf13/cobra"
	"strings"
)

var (
	add_db   model.Db
	validate *validator.Validate
)

var addDbCmd = &cobra.Command{
	Use:   "add-db",
	Short: "Add db value to properties",
	Long:  `Add db value to properties.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Ref :https://stackoverflow.com/questions/33892599/how-to-initialize-values-for-nested-struct-array-in-golang
		if ValidateAllValuePresent() {
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
				service.GenerateGoFile(file_path, conf)
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
}

func ValidateAllValuePresent() bool {
	validate = validator.New()
	err := validate.Struct(add_db)
	if err != nil {
		fmt.Println("****************************************")
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Print("❌ ")
			fmt.Println(FiledError{e}.String())

		}
		return false
	}
	return true
}

type FiledError struct {
	err validator.FieldError
}

func (q FiledError) String() string {
	var sb strings.Builder

	sb.WriteString("Filed is required '" + q.err.Field() + "'")

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}
