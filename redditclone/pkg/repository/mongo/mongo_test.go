package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestNewMongoDB(t *testing.T) {

	cfg := Config{
		Host:     "localhost",
		Port:     "27017",
		Username: `"root"`,
		Password: `"example"`,
	}

	client, err := NewMongoDB(cfg)

	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(nil, nil)

	fmt.Println(err)
	collection := client.Database("coursera").Collection("some")
	newItem := bson.M{
		"_id":         bson.NewObjectId(),
		"title":       "hi",
		"description": "testing",
		"some_filed":  123,
	}
	one, err := collection.InsertOne(nil, newItem)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(one.InsertedID)
}
