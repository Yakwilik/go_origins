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

func TestHandler_Comment(t *testing.T) {
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
	npd := structs.NewCommentData{
		Comment: "data",
	}
	body, err := json.Marshal(npd)
	if err != nil {
		t.Fatalf("can`t marshal json")
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	post := structs.Post{}
	newComment := structs.Comment{
		Author:  author,
		Body:    "data",
		Created: "",
		ID:      primitive.ObjectID{},
	}
	mockPosts.EXPECT().AddComment(newComment, ID).Return(post, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", bytes.NewReader(body))
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.addComment(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_CommentBadID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	npd := structs.NewCommentData{
		Comment: "data",
	}
	body, err := json.Marshal(npd)
	if err != nil {
		t.Fatalf("can`t marshal json")
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.addComment(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_CommentBadPayload(t *testing.T) {
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
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	buf := bytes.NewBufferString(`afsd { : : ""}`)
	req, err := http.NewRequest("GET", "/api/posts/", bytes.NewReader(buf.Bytes()))
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.addComment(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_CommentServiceError(t *testing.T) {
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
	npd := structs.NewCommentData{
		Comment: "data",
	}
	body, err := json.Marshal(npd)
	if err != nil {
		t.Fatalf("can`t marshal json")
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	post := structs.Post{}
	newComment := structs.Comment{
		Author:  author,
		Body:    "data",
		Created: "",
		ID:      primitive.ObjectID{},
	}
	mockPosts.EXPECT().AddComment(newComment, ID).Return(post, errors.New("error"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", bytes.NewReader(body))
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.addComment(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}

func TestHandler_DeleteComment(t *testing.T) {
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
		PostID:    ID.Hex(),
		CommentID: ID.Hex(),
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	post := structs.Post{}
	mockPosts.EXPECT().DeleteComment(ID, ID, author.ID).Return(post, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.deleteComment(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_DeleteCommentBadPostID(t *testing.T) {
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
		CommentID: ID.Hex(),
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.deleteComment(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_DeleteCommentBadCommentID(t *testing.T) {
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
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.deleteComment(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_DeleteCommentServiceError(t *testing.T) {
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
		PostID:    ID.Hex(),
		CommentID: ID.Hex(),
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	post := structs.Post{}
	mockPosts.EXPECT().DeleteComment(ID, ID, author.ID).Return(post, errors.New("error"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.deleteComment(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}
