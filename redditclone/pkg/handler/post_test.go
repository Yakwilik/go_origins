package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/service"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetAllPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	posts := make([]structs.Post, 10)
	mockPosts.EXPECT().GetAllPosts().Return(posts, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Fatalf("can`t create request")
	}
	handlers.getAllPosts(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_GetAllPostsBadService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	posts := make([]structs.Post, 10)
	mockPosts.EXPECT().GetAllPosts().Return(posts, errors.New("bad"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Fatalf("can`t create new request")
	}

	handlers.getAllPosts(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("error status code")
	}
}

func TestHandler_NewPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	npd := structs.NewPostData{
		Title:    "title",
		Category: "cat",
		Type:     "text",
		Text:     "afdsfa",
	}
	body, err := json.Marshal(npd)
	if err != nil {
		t.Fatalf("can`t marshal json")
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:8080", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("error creating request")
	}
	ctx := req.Context()
	author := structs.Author{
		Username: "yakwilik",
		ID:       0,
	}
	mockPosts.EXPECT().NewPost(npd, author).Return(structs.Post{}, nil)
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.newPost(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_NewPostBadJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}

	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	buf := bytes.NewBufferString(`afsd { : : ""}`)
	req, err := http.NewRequest("GET", "http://localhost:8080", bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("error creating request")
	}
	ctx := req.Context()
	author := structs.Author{
		Username: "yakwilik",
		ID:       0,
	}
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.newPost(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}
func TestHandler_NewPostServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	npd := structs.NewPostData{
		Title:    "title",
		Category: "cat",
		Type:     "text",
		Text:     "afdsfa",
	}
	body, err := json.Marshal(npd)
	if err != nil {
		t.Fatalf("can`t marshal json")
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:8080", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("error creating request")
	}
	ctx := req.Context()
	author := structs.Author{
		Username: "yakwilik",
		ID:       0,
	}
	mockPosts.EXPECT().NewPost(npd, author).Return(structs.Post{}, errors.New("error"))
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.newPost(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostByCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	category := "coding"
	muxVars := map[string]string{
		CategoryName: category,
	}
	posts := make([]structs.Post, 10)
	mockPosts.EXPECT().PostsByCategory(category).Return(posts, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPostsByCategory(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostByCategoryEmptyCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPostsByCategory(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostByCategoryServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	category := "coding"
	muxVars := map[string]string{
		CategoryName: category,
	}
	posts := make([]structs.Post, 10)
	mockPosts.EXPECT().PostsByCategory(category).Return(posts, errors.New("error"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPostsByCategory(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	ID := primitive.NewObjectID()
	muxVars := map[string]string{
		PostID: ID.Hex(),
	}
	post := structs.Post{}
	mockPosts.EXPECT().GetPostByID(ID).Return(post, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPost(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostByIDBadID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPost(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostByIDServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	ID := primitive.NewObjectID()
	muxVars := map[string]string{
		PostID: ID.Hex(),
	}
	post := structs.Post{}
	mockPosts.EXPECT().GetPostByID(ID).Return(post, errors.New("error"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPost(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}

func TestHandler_DeletePostByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	ID := primitive.NewObjectID()
	muxVars := map[string]string{
		PostID: ID.Hex(),
	}
	mockPosts.EXPECT().DeletePostByID(ID, 0).Return(nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	author := structs.Author{
		Username: "yakwilik",
		ID:       0,
	}
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.deletePost(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_DeletePostByIDBadID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	author := structs.Author{
		Username: "yakwilik",
		ID:       0,
	}
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.deletePost(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}
func TestHandler_DeletePostByIDServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	ID := primitive.NewObjectID()
	muxVars := map[string]string{
		PostID: ID.Hex(),
	}
	mockPosts.EXPECT().DeletePostByID(ID, 0).Return(errors.New("error"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	author := structs.Author{
		Username: "yakwilik",
		ID:       0,
	}
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.deletePost(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	username := Username
	muxVars := map[string]string{
		UserLogin: username,
	}
	posts := make([]structs.Post, 10)
	mockPosts.EXPECT().PostsByUsername(username).Return(posts, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPostsByUserLogin(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostUserLoginBadLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPostsByUserLogin(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_GetPostServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	username := Username
	muxVars := map[string]string{
		UserLogin: username,
	}
	posts := make([]structs.Post, 10)
	mockPosts.EXPECT().PostsByUsername(username).Return(posts, errors.New("error"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	handlers.getPostsByUserLogin(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}
