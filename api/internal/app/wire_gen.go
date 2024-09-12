// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"context"
	"go-temp/go-sample/api/internal/app/config"
	"go-temp/go-sample/api/internal/app/container"
	"go-temp/go-sample/api/internal/controller"
	"go-temp/go-sample/api/internal/infra/db"
	"go-temp/go-sample/api/internal/pkg/database"
	"go-temp/go-sample/api/internal/usecase"
)

// Injectors from wire.go:

func New(contextContext context.Context) (*container.App, error) {
	dbConfig := config.NewDBConfig()
	databaseDBConfig := ConvertDBEnv(dbConfig)
	databaseDB := database.NewDB(databaseDBConfig)
	userRepository := db.NewUserRepository(databaseDB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	containerContainer := container.NewCtrl(userController)
	app, err := container.NewApp(databaseDB, containerContainer)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// wire.go:

func ConvertDBEnv(cfg *config.DBConfig) *database.DBConfig {
	return &database.DBConfig{
		User:     cfg.User,
		DBName:   cfg.DBName,
		RootPass: cfg.RootPass,
	}
}
