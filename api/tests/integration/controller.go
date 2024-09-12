package integration

import (
	"context"
	"go-temp/go-sample/api/internal/app/container"
	"go-temp/go-sample/api/internal/controller"
	"go-temp/go-sample/api/internal/domain/repo"
	"go-temp/go-sample/api/internal/infra/db"
	"go-temp/go-sample/api/internal/pkg/database"
	"go-temp/go-sample/api/internal/usecase"
)

type InfraInstance struct {
	UserRepository repo.UserRepository
}

func NewTestController(ctx context.Context, database *database.DB) (*container.Container, InfraInstance) {

	ur := db.NewUserRepository(database)

	uu := usecase.NewUserUsecase(ur)

	userController := controller.NewUserController(uu)
	controller := container.NewCtrl(userController)

	InfraInstance := InfraInstance{
		UserRepository: ur,
	}

	return controller, InfraInstance

}
