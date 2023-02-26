package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/service"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	Username = "yakwilik"
	Password = "password"
)

func TestHandlerRegistration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	user := structs.User{
		ID:       0,
		Username: Username,
		Password: Password,
	}
	token := "gqretwertert"
	mockAuth.EXPECT().CreateUser(user).Return(token, nil)
	handlers := NewHandler(mockService)
	data := signInInput{}
	data.Username = Username
	data.Password = Password
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("error didn`t expected")
	}
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewReader(body))
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	handlers.register(w, req)

	resp := w.Result()
	respMap := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&respMap)
	if err != nil {
		return
	}
	if respMap["token"] != token {
		t.Errorf("error token")
	}
}

func TestHandlerRegistrationBadJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	buf := bytes.NewBufferString(`afsd { : : ""}`)
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewReader(buf.Bytes()))
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	handlers.register(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("bad status code")
	}
}

func TestHandlerRegistrationErrorService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	user := structs.User{
		ID:       0,
		Username: Username,
		Password: Password,
	}
	token := "gqretwertert"
	mockAuth.EXPECT().CreateUser(user).Return(token, errors.New("error"))

	data := signInInput{}
	data.Username = Username
	data.Password = Password
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("error didn`t expected")
	}
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewReader(body))
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	handlers.register(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("bad status code")
	}
}

func TestHandlerLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	data := signInInput{}
	data.Username = Username
	data.Password = Password
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("error didn`t expected")
	}
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewReader(body))
	if err != nil {
		return
	}
	mockAuth.EXPECT().GenerateToken("yakwilik", "password").Return("token", nil)
	w := httptest.NewRecorder()
	handlers.login(w, req)
	resp := w.Result()
	respMap := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&respMap)
	if err != nil {
		return
	}
	if respMap["token"] != "token" {
		t.Errorf("error token")
	}
}

func TestHandlerLoginBadJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	buf := bytes.NewBufferString(`afsd { : : ""}`)
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewReader(buf.Bytes()))
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	handlers.login(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandlerLoginServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	data := signInInput{}
	data.Username = Username
	data.Password = Password
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("error didn`t expected")
	}
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewReader(body))
	if err != nil {
		return
	}
	mockAuth.EXPECT().GenerateToken("yakwilik", "password").Return("token", errors.New("error"))
	w := httptest.NewRecorder()
	handlers.login(w, req)
	resp := w.Result()
	if http.StatusBadRequest != resp.StatusCode {
		t.Errorf("bad status code")
	}
}
