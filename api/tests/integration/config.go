package integration

import (
	"os"
)

type DBConfig struct {
	User     string
	DBName   string
	RootPass string
	Host     string
	Port     string
}

func NewDBConfig() *DBConfig {

	user := os.Getenv("TEST_MYSQL_USER_NAME")
	pass := os.Getenv("TEST_MYSQL_ROOT_PASS")
	database := os.Getenv("MYSQL_DB_NAME")
	host := os.Getenv("TEST_MYSQL_HOST")
	port := os.Getenv("PORT")

	return &DBConfig{
		User:     user,
		RootPass: pass,
		DBName:   database,
		Host:     host,
		Port:     port,
	}
}
