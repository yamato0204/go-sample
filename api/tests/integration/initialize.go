package integration

import (
	"context"
	"go-temp/go-sample/api/internal/app/config"
	"go-temp/go-sample/api/internal/app/container"
	"go-temp/go-sample/api/internal/pkg/database"
)

func Initialize(ctx context.Context) (*container.Container, InfraInstance) {

	// DB初期化
	dbConfig := config.NewDBConfig()
	databaseDBConfig := ConvertDBEnv(dbConfig)
	db := database.NewDB(databaseDBConfig)

	app, infraInstances := NewTestController(ctx, db)
	return app, infraInstances
}

func ConvertDBEnv(cfg *config.DBConfig) *database.DBConfig {
	return &database.DBConfig{
		User:     cfg.User,
		DBName:   cfg.DBName,
		RootPass: cfg.RootPass,
	}
}
