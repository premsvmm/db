package model


type Config struct {
	Database []struct {
		JdbcDrive  string `json:"jdbc_drive" mapstructure:"jdbc_drive"`
		Host       string `json:"host" mapstructure:"host"`
		Port       string `json:"port" mapstructure:"port"`
		DbName     string `json:"db_name" mapstructure:"db_name"`
		DbUserName string `json:"db_user_Name" mapstructure:"db_user_Name"`
		DbPassword string `json:"db_password" mapstructure:"db_password"`
		URL        string `json:"url" mapstructure:"url"`
		Name       string `json:"name" mapstructure:"name"`
	} `json:"database"`
	CurrentDbContext string `json:"current_db_context" mapstructure:"current_db_context"`
	ReturnType       string `json:"return_type" mapstructure:"return_type"`
	URL              string `json:"url" mapstructure:"url"`
	Username         string `json:"username" mapstructure:"username"`
	Password         string `json:"password" mapstructure:"password"`
}



