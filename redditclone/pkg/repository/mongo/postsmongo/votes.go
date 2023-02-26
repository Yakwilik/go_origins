package postsmongo

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p *PostsMongo) FindVote(postID primitive.ObjectID, userID int) (bool, error) {
	posts := make([]structs.Post, 0)
	filter := bson.M{"_id": postID, "votes.user": userID}
	find, err := p.posts.Find(nil, filter)
	if err != nil {
		return false, err
	}
	err = find.All(nil, &posts)
	if err != nil {
		return false, err
	}
	return !(len(posts) == 0), nil
}

func (p *PostsMongo) UnVote(postID primitive.ObjectID, userID int) (structs.Post, error) {
	post, err := p.findPostByID(postID)
	if err != nil {
		return post, err
	}
	if len(post.Votes) == 0 {
		return post, nil
	}
	action := bson.M{"$pull": bson.M{"votes": bson.M{"user": userID}}}
	_, err = p.posts.UpdateByID(nil, postID, action)
	if err != nil {
		return post, err
	}
	err = p.SetScoreAndPercentage(postID)
	if err != nil {
		logrus.Println(err)
	}
	post, err = p.findPostByID(postID)
	return post, err
}
func (p *PostsMongo) Vote(postID primitive.ObjectID, vote structs.Vote) (structs.Post, error) {
	post, err := p.findPostByID(postID)
	if err != nil {
		return post, err
	}
	if len(post.Votes) == 0 {
		votes := make([]structs.Vote, 0)
		votes = append(votes, vote)
		_, err = p.posts.UpdateByID(nil, postID, bson.M{"$set": bson.M{"votes": votes}})
		if err != nil {
			return post, err
		}
		err = p.SetScoreAndPercentage(postID)
		if err != nil {
			logrus.Println(err)
		}
		post.Score = vote.Vote
		post.Votes = votes
		return post, nil
	}
	found, err := p.FindVote(postID, vote.User)
	if err != nil {
		return post, err
	}
	if found {
		filter := bson.M{"_id": postID}
		arrayFilters := options.ArrayFilters{
			Filters: []interface{}{bson.M{"vote.user": vote.User}},
		}
		updateOpts := options.UpdateOptions{
			ArrayFilters: &arrayFilters,
		}
		updateOpts.SetUpsert(true)
		_, err = p.posts.UpdateOne(nil, filter, bson.M{"$set": bson.M{"votes.$[vote]": vote}}, &updateOpts)
		if err != nil {
			return post, err
		}
	} else {
		_, err = p.posts.UpdateByID(nil, postID, bson.M{"$addToSet": bson.M{"votes": vote}})
		if err != nil {
			return post, err
		}
	}
	err = p.SetScoreAndPercentage(postID)
	if err != nil {
		logrus.Println(err)
	}
	post, err = p.findPostByID(postID)
	return post, err
}

func (p *PostsMongo) count(postID primitive.ObjectID, vote int) (int, error) {
	pipeline := mongo.Pipeline{}
	matchStage := bson.D{{Key: "$match", Value: bson.M{"_id": postID}}}
	unwindStage := bson.D{{Key: "$unwind", Value: "$votes"}}
	projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "votes", Value: 1}, {Key: "_id", Value: 0}}}}
	insideMatchStage := bson.D{{Key: "$match", Value: bson.M{"votes.vote": vote}}}
	groupStage := bson.D{{Key: "$count", Value: "score"}}
	pipeline = append(pipeline, matchStage)
	pipeline = append(pipeline, unwindStage)
	pipeline = append(pipeline, projectStage)
	pipeline = append(pipeline, insideMatchStage)
	pipeline = append(pipeline, groupStage)
	aggregate, err := p.posts.Aggregate(nil, pipeline)
	if err != nil {
		return 0, err
	}
	result := make([]bson.M, 0)
	err = aggregate.All(nil, &result)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, nil
	}
	count := int(result[0]["score"].(int32))
	return count, nil
}

func (p *PostsMongo) getUpvotesCount(postID primitive.ObjectID) (int, error) {
	res, err := p.count(postID, 1)
	return res, err
}

func (p *PostsMongo) getDownvotesCount(postID primitive.ObjectID) (int, error) {
	res, err := p.count(postID, -1)
	return res, err
}

func (p *PostsMongo) getUpvotesDownVotesCount(postID primitive.ObjectID) (upvote, downvote int, err error) {
	upvote, err = p.getUpvotesCount(postID)
	if err != nil {
		return upvote, downvote, err
	}
	downvote, err = p.getDownvotesCount(postID)
	return upvote, downvote, err
}

func (p *PostsMongo) SetScore(postID primitive.ObjectID) error {
	upvotes, downvotes, err := p.getUpvotesDownVotesCount(postID)
	if err != nil {
		return err
	}
	_, err = p.posts.UpdateByID(nil, postID, bson.M{"$set": bson.M{"score": upvotes - downvotes}})
	return err
}

func (p *PostsMongo) SetUpvotePercentage(postID primitive.ObjectID) error {
	upvotes, downvotes, err := p.getUpvotesDownVotesCount(postID)
	if err != nil {
		return err
	}
	totalVotes := upvotes + downvotes
	var percentage float32
	if totalVotes == 0 {
		percentage = 0
	} else {
		percentage = float32(100 * upvotes / totalVotes)
	}
	_, err = p.posts.UpdateByID(nil, postID, bson.M{"$set": bson.M{"upvotePercentage": percentage}})
	return err
}

func (p *PostsMongo) SetScoreAndPercentage(postID primitive.ObjectID) error {
	err := p.SetScore(postID)
	if err != nil {
		return err
	}
	err = p.SetUpvotePercentage(postID)
	return err
}
