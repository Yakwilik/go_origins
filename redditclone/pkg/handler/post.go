package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	AuthorizationHeader = "Authorization"
)

func (h *Handler) getAllPosts(writer http.ResponseWriter, request *http.Request) {
	posts, err := h.services.GetAllPosts()
	if err != nil {
		JSONResponse(writer, http.StatusInternalServerError, interfaceMap{"message": err})
		return
	}
	JSONResponse(writer, http.StatusOK, posts)
}

func (h *Handler) newPost(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(userCTX).(structs.Author)
	post := structs.NewPostData{}
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "bad request body"})
		return
	}
	newPost, err := h.services.Posts.NewPost(post, user)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "error occurred while creating a post"})
		return
	}
	JSONResponse(writer, http.StatusOK, newPost)
}

func (h *Handler) getPostsByCategory(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	category, ok := vars[CategoryName]
	if !ok {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "empty category"})
		return
	}
	posts, err := h.services.Posts.PostsByCategory(category)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "storage is not doing well"})
		return
	}
	JSONResponse(writer, http.StatusOK, posts)
}

func (h *Handler) getPost(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	if !primitive.IsValidObjectID(vars[PostID]) {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "post id bad format"})
		return
	}
	//  nolint
	id, _ := primitive.ObjectIDFromHex(vars[PostID])
	post, err := h.services.Posts.GetPostByID(id)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "couldn't get post from storage"})
		return
	}
	JSONResponse(writer, http.StatusOK, post)
}

func (h *Handler) deletePost(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(userCTX).(structs.Author)

	vars := mux.Vars(request)
	if !primitive.IsValidObjectID(vars[PostID]) {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "bad id parameter"})
		return
	}
	//  nolint
	id, _ := primitive.ObjectIDFromHex(vars[PostID])
	err := h.services.DeletePostByID(id, user.ID)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "coulnd't delete post"})
		return
	}
	JSONResponse(writer, http.StatusOK, interfaceMap{"message": "success"})
}

func (h *Handler) getPostsByUserLogin(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username, ok := vars[UserLogin]
	if !ok {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "empty username"})
		return
	}
	posts, err := h.services.Posts.PostsByUsername(username)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "storage is not doing well"})
		return
	}
	JSONResponse(writer, http.StatusOK, posts)
}
