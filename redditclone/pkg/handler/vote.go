package handler

import (
	"github.com/gorilla/mux"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) upvote(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(userCTX).(structs.Author)

	vars := mux.Vars(request)
	if !primitive.IsValidObjectID(vars[PostID]) {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "bad post_id parameter"})
		return
	}
	id, _ := primitive.ObjectIDFromHex(vars[PostID]) //  nolint
	vote := structs.Vote{
		User: user.ID,
		Vote: 0,
	}
	post, err := h.services.Posts.Upvote(id, vote)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "couldn't vote for the post"})
		return
	}
	JSONResponse(writer, http.StatusOK, post)
}

func (h *Handler) downvote(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(userCTX).(structs.Author)

	vars := mux.Vars(request)
	if !primitive.IsValidObjectID(vars[PostID]) {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "bad post_id parameter"})
		return
	}
	id, _ := primitive.ObjectIDFromHex(vars[PostID]) //  nolint
	vote := structs.Vote{
		User: user.ID,
		Vote: 0,
	}
	post, err := h.services.Posts.DownVote(id, vote)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "couldn't vote for the post"})
		return
	}
	JSONResponse(writer, http.StatusOK, post)
}

func (h *Handler) unvote(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(userCTX).(structs.Author)

	vars := mux.Vars(request)
	if !primitive.IsValidObjectID(vars[PostID]) {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "bad id"})
		return
	}
	id, _ := primitive.ObjectIDFromHex(vars[PostID]) //  nolint

	post, err := h.services.UnVote(id, user.ID)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "couldn't unvote"})
		return
	}
	JSONResponse(writer, http.StatusOK, post)
}
