package handler

import (
	"context"
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

func TestHandler_Upvote(t *testing.T) {
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
	post := structs.Post{}
	mockPosts.EXPECT().Upvote(ID, structs.Vote{}).Return(post, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.upvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_UpvoteServiceError(t *testing.T) {
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
	post := structs.Post{}
	mockPosts.EXPECT().Upvote(ID, structs.Vote{}).Return(post, errors.New("error"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.upvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}

func TestHandler_UpvoteBadID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.upvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_Downvote(t *testing.T) {
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
	post := structs.Post{}
	mockPosts.EXPECT().DownVote(ID, structs.Vote{}).Return(post, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.downvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_DownvoteServiceError(t *testing.T) {
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
	post := structs.Post{}
	mockPosts.EXPECT().DownVote(ID, structs.Vote{}).Return(post, errors.New("error"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.downvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}

func TestHandler_DownvoteBadID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.downvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}

func TestHandler_UnVote(t *testing.T) {
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
	post := structs.Post{}
	mockPosts.EXPECT().UnVote(ID, 0).Return(post, nil)
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.unvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error status code")
	}
}

func TestHandler_UnVoteServiceError(t *testing.T) {
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
	post := structs.Post{}
	mockPosts.EXPECT().UnVote(ID, 0).Return(post, errors.New("errors"))
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	req = mux.SetURLVars(req, muxVars)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.unvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Errorf("error status code")
	}
}

func TestHandler_UnVoteBadID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := service.NewMockAuthorization(ctrl)
	mockPosts := service.NewMockPosts(ctrl)
	mockService := &service.Service{
		Authorization: mockAuth,
		Posts:         mockPosts,
	}
	author := structs.Author{
		Username: Username,
		ID:       0,
	}
	handlers := NewHandler(mockService)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/posts/", nil)
	if err != nil {
		t.Fatalf("can`t create new request")
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, userCTX, author)
	handlers.unvote(w, req.WithContext(ctx))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error status code")
	}
}
