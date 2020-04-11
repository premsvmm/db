package model

type Mock struct {
	Database struct {
		JdbcDriver string `json:"jdbcDriver"`
		Host       string `json:"host"`
		Port       int    `json:"port"`
		DbName     string `json:"dbName"`
		DbUserName string `json:"dbUserName"`
		DbPassword string `json:"dbPassword"`
		URL        string `json:"url"`
	} `json:"database"`
	MySQLQuery string `json:"mySqlQuery"`
	SelectType string `json:"selectType"`
}