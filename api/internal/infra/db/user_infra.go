package db

import (
	"context"
	"database/sql"
	"go-temp/go-sample/api/internal/domain/model"
	"go-temp/go-sample/api/internal/domain/repo"
	"go-temp/go-sample/api/internal/pkg/database"

	"github.com/pkg/errors"
)

type userRepository struct {
	db *database.DB
}

// NewUserRepository は UserRepository を作成します
func NewUserRepository(db *database.DB) repo.UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetUserByID はIDからユーザーを取得します
func (r *userRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Get(user, "SELECT * FROM users WHERE id = ?", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return user, nil
}

// UpdateUser はユーザー情報を更新します

// CreateUser はユーザーを作成します
