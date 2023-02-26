package postsmongo

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var incAction = bson.M{"$inc": bson.M{"views": 1}}

func (p *PostsMongo) GetAllPosts() ([]structs.Post, error) {
	items := make([]structs.Post, 0)

	_, err := p.posts.UpdateMany(nil, bson.M{}, incAction)
	if err != nil {
		return nil, err
	}
	cur, err := p.posts.Find(nil, bson.M{}, nil)
	if err != nil {
		return items, err
	}
	err = cur.All(nil, &items)
	return items, err
}

func (p *PostsMongo) PostsByCategory(category string) ([]structs.Post, error) {
	items := make([]structs.Post, 0)
	filter := bson.M{"category": category}

	_, err := p.posts.UpdateMany(nil, filter, incAction)
	if err != nil {
		logrus.Println(err)
	}
	cur, err := p.posts.Find(nil, filter, nil)
	if err != nil {
		return nil, err
	}
	err = cur.All(nil, &items)
	return items, err
}
func (p *PostsMongo) PostsByUsername(username string) ([]structs.Post, error) {
	items := make([]structs.Post, 0)
	filter := bson.M{"author.username": username}
	_, err := p.posts.UpdateMany(nil, filter, incAction)
	if err != nil {
		logrus.Println(err)
	}
	cur, err := p.posts.Find(nil, filter, nil)
	if err != nil {
		return nil, err
	}
	err = cur.All(nil, &items)
	return items, err
}

func (p *PostsMongo) CreateNewPost(data structs.NewPostData, author structs.Author) (structs.Post, error) {
	newPost := structs.Post{
		Author:           author,
		Category:         data.Category,
		Comments:         nil,
		Created:          time.Now().Format("2006-01-02T15:04:05.000"),
		CreatedUnix:      time.Now().Unix(),
		Score:            0,
		Text:             data.Text,
		URL:              data.URL,
		Title:            data.Title,
		Type:             data.Type,
		UpvotePercentage: 0,
		Views:            0,
		Votes:            nil,
	}

	one, err := p.posts.InsertOne(nil, newPost)
	newPost.ID = one.InsertedID.(primitive.ObjectID)

	return newPost, err
}
func (p *PostsMongo) findPostByID(id primitive.ObjectID) (structs.Post, error) {
	item := structs.Post{}

	cur := p.posts.FindOne(nil, bson.M{"_id": id})
	if cur.Err() != nil {
		return item, cur.Err()
	}

	err := cur.Decode(&item)
	return item, err
}

func (p *PostsMongo) GetPostByID(id primitive.ObjectID) (structs.Post, error) {
	item, err := p.findPostByID(id)
	if err != nil {
		return item, err
	}
	_, err = p.posts.UpdateByID(nil, id, incAction)
	if err != nil {
		logrus.Println(err)
	}
	return item, nil
}

func (p *PostsMongo) DeletePostByID(postID primitive.ObjectID, userID int) error {

	post, err := p.findPostByID(postID)
	if err != nil {
		return err
	}
	if post.Author.ID != userID {
		return errors.New("you can`t delete this post")
	}

	_, err = p.posts.DeleteOne(nil, bson.M{"_id": postID})
	return err
}
