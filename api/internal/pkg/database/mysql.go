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
}

func NewDB(cfg *DBConfig) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(db)/%s?charset=utf8&parseTime=true&loc=Asia%%2FTokyo", cfg.User, cfg.RootPass, cfg.DBName)

	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	return &DB{db}
}
