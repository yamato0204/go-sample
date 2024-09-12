package config

import (
	"os"
)

type DBConfig struct {
	User     string
	DBName   string
	RootPass string
}

func NewDBConfig() *DBConfig {

	user := os.Getenv("MYSQL_USER_NAME")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DB_NAME")

	return &DBConfig{
		User:     user,
		RootPass: pass,
		DBName:   database,
	}
}
