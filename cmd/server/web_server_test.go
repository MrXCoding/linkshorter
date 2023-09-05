package server

//import (
//	"fmt"
//	"github.com/MrXCoding/linkshorter/internal/storage"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestHandle(t *testing.T) {
//	db := storage.NewInMemory()
//
//	type want struct {
//		url        string
//		statusCode int
//		hasher       string
//	}
//	tests := []struct {
//		request string
//		method  string
//		name    string
//		url     string
//		storage storage.Repository
//		want    want
//	}{
//		{
//			request: "/",
//			name:    "simple test #1",
//			url:     "http:/ya.ru",
//			method:  http.MethodPost,
//			storage: db,
//			want: want{
//				statusCode: http.StatusOK,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			request := httptest.NewRequest(http.MethodPost, tt.request, nil)
//			w := httptest.NewRecorder()
//			h := http.HandlerFunc(handle(db))
//			h(w, request)
//
//			result := w.Result()
//
//			fmt.Print(result)
//		})
//	}
//}
