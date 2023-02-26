package service

import (
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
)

type PostService struct {
	repo repository.Posts
}

func NewPostService(repo repository.Posts) *PostService {
	return &PostService{repo: repo}
}

func sortPostsByScore(posts []structs.Post) []structs.Post { //  nolint
	sort.Slice(posts, func(i, j int) bool { return posts[i].Score > posts[j].Score })
	return posts
}

func sortPostsByCreatedUnix(posts []structs.Post) []structs.Post {
	sort.Slice(posts, func(i, j int) bool { return posts[i].CreatedUnix < posts[j].CreatedUnix })
	return posts
}
func (p *PostService) GetAllPosts() (posts []structs.Post, err error) {
	posts, err = p.repo.GetAllPosts()
	sortPostsByScore(posts)
	return posts, err
}

func (p *PostService) PostsByCategory(category string) (posts []structs.Post, err error) {
	posts, err = p.repo.PostsByCategory(category)
	sortPostsByScore(posts)
	return posts, err
}

func (p *PostService) PostsByUsername(username string) ([]structs.Post, error) {
	posts, err := p.repo.PostsByUsername(username)
	sortPostsByCreatedUnix(posts)
	return posts, err
}

func (p *PostService) NewPost(data structs.NewPostData, author structs.Author) (structs.Post, error) {
	return p.repo.CreateNewPost(data, author)
}

func (p *PostService) GetPostByID(id primitive.ObjectID) (structs.Post, error) {
	return p.repo.GetPostByID(id)
}

func (p *PostService) AddComment(comment structs.Comment, postID primitive.ObjectID) (structs.Post, error) {
	return p.repo.AddComment(comment, postID)
}

func (p *PostService) DeleteComment(commentID, postID primitive.ObjectID, userID int) (structs.Post, error) {
	return p.repo.DeleteComment(commentID, postID, userID)
}

func (p *PostService) DeletePostByID(postID primitive.ObjectID, userID int) error {
	return p.repo.DeletePostByID(postID, userID)
}

func (p *PostService) Upvote(postID primitive.ObjectID, vote structs.Vote) (structs.Post, error) {
	vote.Vote = 1
	return p.repo.Vote(postID, vote)
}

func (p *PostService) DownVote(postID primitive.ObjectID, vote structs.Vote) (structs.Post, error) {
	vote.Vote = -1
	return p.repo.Vote(postID, vote)
}
func (p *PostService) UnVote(postID primitive.ObjectID, userID int) (structs.Post, error) {
	return p.repo.UnVote(postID, userID)
}
