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
			fmt.Println("HAHA. To smart error in you query pls check "+ string(data))
			fmt.Println("*******************************************************")
		}
	}
}

func GenerateGoFile(file_path string,conf model.Config){
	jsonValue, _ := json.MarshalIndent(conf,"","")
	formatedBody := pretty.Pretty(jsonValue)
	ioutil.WriteFile(file_path,formatedBody,0644)
}
