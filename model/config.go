package model

type Config struct {
	Database         []Db   `json:"database"`
	CurrentDbContext string `json:"current_db_context" mapstructure:"current_db_context"`
	ReturnType       string `json:"return_type" mapstructure:"return_type"`
	URL              string `json:"url" mapstructure:"url"`
	Username         string `json:"username" mapstructure:"username"`
	Password         string `json:"password" mapstructure:"password"`
}

type Db struct {
	JdbcDrive  string `json:"jdbc_drive" mapstructure:"jdbc_drive" validate:"required"`
	Host       string `json:"host" mapstructure:"host" validate:"required"`
	Port       string `json:"port" mapstructure:"port" validate:"required"`
	DbName     string `json:"db_name" mapstructure:"db_name" validate:"required"`
	DbUserName string `json:"db_user_Name" mapstructure:"db_user_Name" validate:"required"`
	DbPassword string `json:"db_password" mapstructure:"db_password" validate:"required"`
	URL        string `json:"url" mapstructure:"url"`
	Name       string `json:"name" mapstructure:"name" validate:"required"`
}
