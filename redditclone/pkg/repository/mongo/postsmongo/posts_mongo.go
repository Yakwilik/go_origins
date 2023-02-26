package postsmongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsMongo struct {
	client *mongo.Client
	posts  *mongo.Collection
}

func NewPostsMongo(client *mongo.Client, db, postsCollection string) *PostsMongo {
	if client == nil {
		return nil
	}
	collection := client.Database(db).Collection(postsCollection)
	if collection == nil {
		return nil
	}
	return &PostsMongo{client: client, posts: collection}

}
