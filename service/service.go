package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/premsvmm/db/model"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ExecuteMock(sql string, conf model.Config) {
	var mock model.Mock
	mock.SelectType = conf.ReturnType
	mock.MySQLQuery = sql
	for key, _ := range conf.Database {
		if conf.Database[key].Name == conf.CurrentDbContext {
			mock.Database.URL = conf.Database[key].URL
			mock.Database.DbName = conf.Database[key].DbName
			mock.Database.DbUserName = conf.Database[key].DbUserName
			mock.Database.DbPassword = conf.Database[key].DbPassword
			mock.Database.Host = conf.Database[key].Host
			mock.Database.JdbcDriver = conf.Database[key].JdbcDrive
			mock.Database.Port, _ = strconv.Atoi(conf.Database[key].Port)
		}
	}
	jsonValue, _ := json.Marshal(mock)
	request, _ := http.NewRequest("POST", conf.URL, bytes.NewBuffer(jsonValue))
	request.SetBasicAuth(conf.Username, conf.Password)
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode == 200 {
			formatedBody := pretty.Pretty(data)
			fmt.Println(string(pretty.Color(formatedBody, nil)))
		} else {
			fmt.Println("*******************************************************")
			fmt.Println("HAHA. To smart error in you query pls check " + string(data))
			fmt.Println("*******************************************************")
		}
	}
}

func GenerateGoFile(file_path string, conf model.Config) {
	jsonValue, _ := json.MarshalIndent(conf, "", "")
	formatedBody := pretty.Pretty(jsonValue)
	ioutil.WriteFile(file_path, formatedBody, 0644)
}

func ListDownTheDBName(conf model.Config) {
	fmt.Println(`List of Database added in config file:
**************************************`)
	for key, _ := range conf.Database {
		fmt.Println("👻 " + conf.Database[key].Name)
	}
	fmt.Println("**************************************")
}

func PrintTheConfig(conf model.Config) {
	jsonValue, _ := json.MarshalIndent(conf, "", "")
	formatedBody := pretty.Pretty(jsonValue)
	fmt.Println(string(pretty.Color(formatedBody, nil)))
}

func PrintTheDatabase(conf model.Db) {
	jsonValue, _ := json.MarshalIndent(conf, "", "")
	formatedBody := pretty.Pretty(jsonValue)
	fmt.Println(string(pretty.Color(formatedBody, nil)))
}

var (
	count = 1
)

func AskForConfirmation() bool {
	if count == 1 {
		fmt.Print("Please type (y)) or (n) and then press enter:")
	}
	var response string
	_, err := fmt.Scanln(&response)
	count++
	if err != nil {
		fmt.Println(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		count = 1
		return true
	} else if containsString(nokayResponses, response) {
		count = 1
		return false
	} else {
		fmt.Print("Seems wrong Input , Please type (y)) or (n) and then press enter:")
		return AskForConfirmation()
	}
}

func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}
