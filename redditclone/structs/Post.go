package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author           Author             `json:"author" bson:"author"`
	Category         string             `json:"category" bson:"category"`
	Comments         []Comment          `json:"comments" bson:"comments"`
	Created          string             `json:"created" bson:"created"`
	CreatedUnix      int64              `json:"-" bson:"createdUnix"`
	Score            int                `json:"score" bson:"score"`
	Upvote           int                `json:"-" bson:"-"`
	Text             string             `json:"text,omitempty" bson:"text,omitempty"`
	URL              string             `json:"url,omitempty" bson:"url,omitempty"`
	Title            string             `json:"title" bson:"title"`
	Type             string             `json:"type" bson:"type"`
	UpvotePercentage int                `json:"upvotePercentage" bson:"upvotePercentage"`
	Views            int                `json:"views" bson:"views"`
	Votes            []Vote             `json:"votes" bson:"votes"`
}

type NewPostData struct {
	Title    string `json:"title" bson:"title"`
	Category string `json:"category" bson:"category"`
	Type     string `json:"type" bson:"type"`
	Text     string `json:"text,omitempty" bson:"text,omitempty"`
	URL      string `json:"url,omitempty" bson:"url,omitempty"`
}

type Author struct {
	Username string `json:"username" bson:"username"`
	ID       int    `json:"id,string" bson:"id,omitempty"`
}

type Comment struct {
	Author  Author             `json:"author" bson:"author"`
	Body    string             `json:"body" bson:"body"`
	Created string             `json:"created" bson:"created"`
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type NewCommentData struct {
	Comment string `json:"comment" bson:"comment"`
}

type Vote struct {
	User int `json:"user,,string" bson:"user"`
	Vote int `json:"vote" bson:"vote"`
}
