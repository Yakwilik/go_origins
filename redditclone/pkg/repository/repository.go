package repository

import (
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	CreateUser(user structs.User) (int, error)
	GetUser(username, password string) (structs.User, error)
	CheckIdentity(username string, id int) error
	GenerateToken(username string, id int, iat, exp int64) error
}

type Posts interface {
	GetAllPosts() ([]structs.Post, error)
	CreateNewPost(data structs.NewPostData, author structs.Author) (structs.Post, error)
	PostsByCategory(category string) ([]structs.Post, error)
	GetPostByID(id primitive.ObjectID) (structs.Post, error)
	AddComment(comment structs.Comment, postID primitive.ObjectID) (structs.Post, error)
	DeleteComment(commentID primitive.ObjectID, postID primitive.ObjectID, userID int) (structs.Post, error)
	PostsByUsername(username string) ([]structs.Post, error)
	DeletePostByID(postID primitive.ObjectID, userID int) error
	Vote(postID primitive.ObjectID, vote structs.Vote) (structs.Post, error)
	UnVote(postID primitive.ObjectID, userID int) (structs.Post, error)
}

type Repository struct {
	Authorization
	Posts
}
