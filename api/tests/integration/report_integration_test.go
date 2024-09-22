package integration

import (
	"context"
	"encoding/json"
	"go-temp/go-sample/api/internal/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetReportByUserID(t *testing.T) {
	ctx := context.Background()
	app, _ := Initialize(ctx)

	tests := []struct {
		name        string
		expected    controller.Report
		expectedErr string
		preparation func()
	}{
		{
			name: "success",
			expected: controller.Report{
				ID:           "1e8b6d70-9c2f-11eb-a8b3-0242ac130003",
				Comment:      "トマトサラダのレビュー。簡単で美味しかったです！",
				ThumbnailUrl: "https://video.com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg",
				UserID:       "a4e8b5d4-9c1f-11eb-a8b3-0242ac130003",
				Recipe: controller.Recipe{
					ID:           "01J6CXNF4GYJFTNTTH9TZ3MSJG",
					Title:        "トマトサラダ",
					ThumbnailUrl: "https://video.com/production/videos/b331637e-1067-4471-acf3-/.jpg",
					Recipe:       "トマトとバジルのシンプルなサラダ。",
					Category: controller.Category{
						ID:   "1",
						Name: "フレンチ",
					},
					Ingredient: `[{"ingredient": "トマト", "amount": "100g"}, {"ingredient": "バジル", "amount": "10枚"}]`,
					CreatedAt:  "2024-08-01 10:00:00",
					UpdatedAt:  "2024-08-01 10:00:00",
				},
				CreatedAt: "2024-09-01 09:00:00",
				UpdatedAt: "2024-09-01 09:00:00",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のEchoコンテキストとレスポンスレコーダを作成
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/reports/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// 実際のハンドラを呼び出し
			err := app.ReportCtrl.GetReportByUserID(c)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
				return
			}

			// レスポンスのステータスコードを確認
			assert.Equal(t, http.StatusOK, rec.Code)

			// レスポンスをデコード
			var actual controller.Report
			err = json.Unmarshal(rec.Body.Bytes(), &actual)
			assert.NoError(t, err)

			// 期待される結果と一致するか確認
			assert.Equal(t, tt.expected.ID, actual.ID)
			assert.Equal(t, tt.expected.Comment, actual.Comment)
			assert.Equal(t, tt.expected.ThumbnailUrl, actual.ThumbnailUrl)
			assert.Equal(t, tt.expected.UserID, actual.UserID)
			assert.Equal(t, tt.expected.Recipe.ID, actual.Recipe.ID)
			assert.Equal(t, tt.expected.Recipe.Title, actual.Recipe.Title)
			assert.Equal(t, tt.expected.Recipe.ThumbnailUrl, actual.Recipe.ThumbnailUrl)
			assert.Equal(t, tt.expected.Recipe.Recipe, actual.Recipe.Recipe)
			assert.Equal(t, tt.expected.Recipe.Category.ID, actual.Recipe.Category.ID)
			assert.Equal(t, tt.expected.Recipe.Category.Name, actual.Recipe.Category.Name)
			assert.Equal(t, tt.expected.Recipe.Ingredient, actual.Recipe.Ingredient)
			assert.Equal(t, tt.expected.Recipe.CreatedAt, actual.Recipe.CreatedAt)
			assert.Equal(t, tt.expected.Recipe.UpdatedAt, actual.Recipe.UpdatedAt)
			assert.Equal(t, tt.expected.CreatedAt, actual.CreatedAt)
			assert.Equal(t, tt.expected.UpdatedAt, actual.UpdatedAt)
		})
	}
}
