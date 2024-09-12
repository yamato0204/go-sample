package integration

import (
	"context"
	"encoding/json"
	"go-temp/go-sample/api/internal/domain/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	app, _ := Initialize(ctx)

	tests := []struct {
		name        string
		input       string
		expected    model.User
		expectedErr string
		preparation func()
	}{
		{
			name:  "success",
			input: "1",
			expected: model.User{
				ID:    1,
				Name:  "test",
				Email: "test@test.com",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のEchoコンテキストとレスポンスレコーダを作成
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/users/"+tt.input, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// 実際のハンドラを呼び出し
			err := app.UserCtrl.GetUserByID(c)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
				return
			}

			// レスポンスのステータスコードを確認
			assert.Equal(t, http.StatusOK, rec.Code)

			// レスポンスをデコード
			var actual model.User
			err = json.Unmarshal(rec.Body.Bytes(), &actual)
			assert.NoError(t, err)

			// 期待される結果と一致するか確認
			assert.Equal(t, tt.expected.ID, actual.ID)
			assert.Equal(t, tt.expected.Name, actual.Name)
			assert.Equal(t, tt.expected.Email, actual.Email)
		})
	}
}
