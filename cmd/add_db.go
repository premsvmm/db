package cmd

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"strings"
)

type AddDb struct {
	JdbcDriver string `validate:"required"`
	Host       string `validate:"required"`
	Port       string `validate:"required"`
	DbName     string `validate:"required"`
	DbUsername string `validate:"required"`
	DbPassword string `validate:"required"`
	Url        string
	Name       string `validate:"required"`
}

var (
	add_db   AddDb
	validate *validator.Validate
)

// addDbCmd represents the addDb command
var addDbCmd = &cobra.Command{
	Use:   "add-db",
	Short: "Add db value to properties",
	Long:  `Add db value to properties.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ValidateAllValuePresent() {
			fmt.Print("The above mentioned fields are required")
		} else {
			fmt.Print("false")
		}
	},
}

func init() {
	rootCmd.AddCommand(addDbCmd)
	addDbCmd.Flags().StringVarP(&add_db.JdbcDriver, "jdbcdriver", "1", "", "JDBC criver")
	addDbCmd.Flags().StringVarP(&add_db.Host, "host", "2", "", "Host")
	addDbCmd.Flags().StringVarP(&add_db.Port, "port", "3", "", "Port")
	addDbCmd.Flags().StringVarP(&add_db.DbName, "dbname", "4", "", "Database name")
	addDbCmd.Flags().StringVarP(&add_db.DbUsername, "dbusername", "5", "", "Database username")
	addDbCmd.Flags().StringVarP(&add_db.DbPassword, "dbpassword", "6", "", "Database password")
	addDbCmd.Flags().StringVarP(&add_db.Name, "name", "7", "", "Name")
	addDbCmd.Flags().StringVarP(&add_db.Url, "dburl", "8", "", "Datbase Url")
}

func ValidateAllValuePresent() bool {
	validate = validator.New()
	err := validate.Struct(add_db)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Print("❌ ")
			fmt.Println(fieldError{e}.String())

		}
		return false
	}
	return true
}

type fieldError struct {
	err validator.FieldError
}

func (q fieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}
