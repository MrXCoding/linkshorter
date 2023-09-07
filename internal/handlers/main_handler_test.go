package handlers

import (
	"bytes"
	"github.com/MrXCoding/linkshorter/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type HasherForTest struct {
}

func (h *HasherForTest) Encode(str string, seed string) string {
	return str
}

func TestHandle(t *testing.T) {
	db := storage.NewInMemory(&HasherForTest{})

	type want struct {
		url        string
		statusCode int
		hash       string
	}
	tests := []struct {
		request string
		name    string
		body    string
		method  string
		handler http.HandlerFunc
		want    want
	}{
		{
			request: "/",
			name:    "POST test",
			body:    "ya.ru",
			method:  http.MethodPost,
			handler: Save(db),
			want: want{
				statusCode: http.StatusCreated,
			},
		},
		{
			name:    "GET test",
			request: "/ya.ru",
			method:  http.MethodGet,
			body:    "ya.ru",
			handler: Get(db),
			want: want{
				statusCode: http.StatusTemporaryRedirect,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.request, bytes.NewBufferString(tt.body))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(tt.handler)
			h(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

			if tt.method == http.MethodGet {
				header := result.Header.Get("Location")
				assert.Equal(t, header, tt.body)
			}

			if tt.method == http.MethodPost {
				userResult, err := io.ReadAll(result.Body)
				require.NoError(t, err)
				err = result.Body.Close()
				require.NoError(t, err)

				assert.True(t, strings.Contains(string(userResult), baseURL))
			}
		})
	}
}
