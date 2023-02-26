package postsmongo

import (
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (p *PostsMongo) AddComment(comment structs.Comment, postID primitive.ObjectID) (structs.Post, error) {
	post, err := p.findPostByID(postID)
	if err != nil {
		return post, err
	}
	comment.Created = time.Now().Format("2006-01-02T15:04:05.000")
	comment.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	if len(post.Comments) == 0 {
		comments := make([]structs.Comment, 0)
		comments = append(comments, comment)
		post.Comments = comments
		_, err = p.posts.UpdateByID(nil, postID, bson.M{"$set": bson.M{"comments": comments}})
		return post, err
	}
	post.Comments = append(post.Comments, comment)
	_, err = p.posts.UpdateByID(nil, postID, bson.M{"$addToSet": bson.M{"comments": comment}})
	return post, err
}

func (p *PostsMongo) DeleteComment(commentID, postID primitive.ObjectID, userID int) (structs.Post, error) {

	post, err := p.findPostByID(postID)
	if err != nil {
		return post, nil
	}
	filter := bson.M{"_id": postID}

	action := bson.M{"$pull": bson.M{"comments": bson.M{"_id": commentID, "author.id": userID}}}
	up, err := p.posts.UpdateOne(nil, filter, action)
	if err != nil {
		return post, err
	}
	if up.ModifiedCount == 0 && post.Author.ID == userID {
		up, err = p.posts.UpdateOne(nil, filter, bson.M{"$pull": bson.M{"comments": bson.M{"_id": commentID}}})
		if err != nil {
			return post, err
		}
	}
	post, err = p.findPostByID(postID)
	return post, err
}
