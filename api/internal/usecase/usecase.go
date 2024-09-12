package usecase

import (
	"context"
	"go-temp/go-sample/api/internal/domain/model"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}
