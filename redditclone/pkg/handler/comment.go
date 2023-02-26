package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) addComment(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(userCTX).(structs.Author)
	vars := mux.Vars(request)
	if !primitive.IsValidObjectID(vars[PostID]) {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "empty post id"})
		return
	}
	//  nolint
	id, _ := primitive.ObjectIDFromHex(vars[PostID])

	commentBody := structs.NewCommentData{}
	err := json.NewDecoder(request.Body).Decode(&commentBody)
	if err != nil {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "bad request body"})
		return
	}
	newComment := structs.Comment{
		Author: structs.Author{
			Username: user.Username,
			ID:       user.ID,
		},
		Body:    commentBody.Comment,
		Created: "",
	}
	comment, err := h.services.Posts.AddComment(newComment, id)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "couldn't create comment"})
		return
	}
	JSONResponse(writer, http.StatusOK, comment)
}

func (h *Handler) deleteComment(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(userCTX).(structs.Author)
	vars := mux.Vars(request)
	if !primitive.IsValidObjectID(vars[PostID]) {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "bad post_id parameter"})
		return
	}
	//  nolint
	postID, _ := primitive.ObjectIDFromHex(vars[PostID])
	if !primitive.IsValidObjectID(vars[CommentID]) {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "bad comment_id parameter"})
		return
	}
	//  nolint
	commentID, _ := primitive.ObjectIDFromHex(vars[CommentID])
	post, err := h.services.DeleteComment(commentID, postID, user.ID)
	if err != nil {
		JSONResponse(writer, http.StatusInsufficientStorage, interfaceMap{"message": "couldn't delete comment"})
		return
	}
	JSONResponse(writer, http.StatusOK, post)

}
