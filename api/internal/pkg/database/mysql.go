package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

type DBConfig struct {
	User     string
	DBName   string
	RootPass string
	Host     string
	Port     string
}

func NewDB(cfg *DBConfig) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.User, cfg.RootPass, cfg.Host, cfg.DBName)
	log.Printf("dsn: %s", dsn)

	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	log.Println("db connection success")

	return &DB{db}
}
