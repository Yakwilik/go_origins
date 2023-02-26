package handler

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/service"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_InitRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)

	router := handlers.InitRoutes()

	if router == nil {
		t.Errorf("error nil value")
	}
}
func TestHandle_JsonResponse(t *testing.T) {
	buf := bytes.NewBufferString(`afsd { : : ""}`)

	w := httptest.NewRecorder()
	JSONResponse(w, http.StatusOK, buf.Bytes())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewReader(buf.Bytes()))
	if err != nil {
		return
	}
	handlers := NewHandler(mockService)
	handlers.indexHTML(w, req)

}

func TestMiddlewareErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(userCTX)
		if user == nil {
			t.Errorf("user must not be nill")
		}
		_, ok := user.(structs.Author)
		if !ok {
			t.Error("not valid Author")
		}
		w.WriteHeader(http.StatusOK)
	})
	handlers := NewHandler(mockService)
	handlerWithMiddleware := handlers.authMiddleware(next)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)

	if err != nil {
		t.Fatalf("can`t create request")
	}

	handlerWithMiddleware.ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("bad status code")
	}
	token := "fasdfadsfga"
	req.Header.Set(AuthorizationHeader, "token "+token)
	author := structs.Author{
		Username: Username,
		ID:       0,
	}

	mockAuth.EXPECT().ParseToken(token).Return(author, errors.New("error"))
	handlerWithMiddleware.ServeHTTP(w, req)

	res = w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("bad status code")
	}
}

func TestMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(userCTX)
		if user == nil {
			t.Errorf("user must not be nill")
		}
		_, ok := user.(structs.Author)
		if !ok {
			t.Error("not valid Author")
		}
		w.WriteHeader(http.StatusOK)
	})
	handlers := NewHandler(mockService)
	handlerWithMiddleware := handlers.authMiddleware(next)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)

	if err != nil {
		t.Fatalf("can`t create request")
	}

	token := "fasdfadsfga"
	req.Header.Set(AuthorizationHeader, "token "+token)
	author := structs.Author{
		Username: Username,
		ID:       0,
	}

	mockAuth.EXPECT().ParseToken(token).Return(author, nil)
	handlerWithMiddleware.ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("bad status code")
	}

}
