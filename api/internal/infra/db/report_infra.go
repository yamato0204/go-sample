package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-temp/go-sample/api/internal/domain/model"
	"go-temp/go-sample/api/internal/domain/repo"
	"go-temp/go-sample/api/internal/pkg/database"

	"github.com/pkg/errors"
)

type reportRepository struct {
	db *database.DB
}

func NewReportRepository(db *database.DB) repo.ReportRepository {
	return &reportRepository{
		db: db,
	}
}

func (rr *reportRepository) GetReportByUserID(ctx context.Context, userID string) ([]*model.Report, error) {
	var report []*model.Report
	err := rr.db.GetContext(ctx, &report, `
	SELECT 
		r.id,
		r.comment,
		r.thumbnail_url,
		r.author_id,
		r.recipe_id,
		r.created_at,
		r.updated_at,

		rc.id AS recipe_id,
		rc.title AS recipe_title,
		rc.thumbnail_url AS recipe_thumbnail,
		rc.recipe,
		rc.category_id,
		rc.ingredient,
		rc.created_at AS recipe_created_at,
		rc.updated_at AS recipe_updated_at,

		c.id AS category_id,
		c.name AS category_name

	FROM 
		reports AS r
	INNER JOIN
		users AS u
	ON r.author_id = u.id
	INNER JOIN
		recipes AS rc
	ON r.recipe_id = rc.id
	INNER JOIN
		categories AS c
	ON rc.category_id = c.id
	WHERE
		r.author_id = ?
	`, userID)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows in infra")
		return nil, err

	}

	if err != nil {
		return nil, err
	}

	return report, nil
}
