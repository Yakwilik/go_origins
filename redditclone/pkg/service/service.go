package service

import (
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	CreateUser(user structs.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (structs.Author, error)
	CheckIdentity(username string, id int) error
}

type Posts interface {
	GetAllPosts() ([]structs.Post, error)
	PostsByCategory(category string) ([]structs.Post, error)
	NewPost(data structs.NewPostData, author structs.Author) (structs.Post, error)
	GetPostByID(id primitive.ObjectID) (structs.Post, error)
	AddComment(comment structs.Comment, postID primitive.ObjectID) (structs.Post, error)
	DeleteComment(commentID primitive.ObjectID, postID primitive.ObjectID, userID int) (structs.Post, error)
	PostsByUsername(username string) ([]structs.Post, error)
	DeletePostByID(postID primitive.ObjectID, userID int) error
	Upvote(postID primitive.ObjectID, vote structs.Vote) (structs.Post, error)
	DownVote(postID primitive.ObjectID, vote structs.Vote) (structs.Post, error)
	UnVote(postID primitive.ObjectID, userID int) (structs.Post, error)
}

type Service struct {
	Authorization
	Posts
}

func NewService(repos *repository.Repository) *Service {

	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Posts:         NewPostService(repos.Posts),
	}
}
