package postsmongo

import (
	mg "gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository/mongo"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var src = rand.NewSource(time.Now().UnixNano())

func RandString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

var config = mg.Config{
	Host:     "localhost",
	Port:     "27017",
	Username: `root`,
	Password: `example`,
}

var client *mongo.Client
var posts *PostsMongo

func init() {
	var err error
	client, err = mg.NewMongoDB(config)
	if err != nil {
		panic(err)
	}
	posts = NewPostsMongo(client, "coursera", "posts")
}
func TestNewPostsMongo(t *testing.T) {

	newMongoPosts := NewPostsMongo(nil, RandString(10), RandString(10))

	if newMongoPosts != nil {
		t.Errorf("nil expected")
	}
}
func TestPostsMongo_CreateNewPost(t *testing.T) {
	data := structs.NewPostData{
		Title:    "secondpost",
		Category: "life",
		Type:     "coding",
		Text:     "yes text",
		URL:      "yes url",
	}

	author := structs.Author{
		Username: "khasbulat",
		ID:       0,
	}
	post, err := posts.CreateNewPost(data, author)
	if err != nil {
		t.Errorf("error occured where not expected: %v", err)
	}
	idString := post.ID.Hex()
	if !primitive.IsValidObjectID(idString) {
		t.Errorf("invalid ObjectID didn't expected")
	}
	_, err = posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}

}

func TestPostsMongo_GetAllPosts(t *testing.T) {
	dataSlice := []structs.NewPostData{
		{
			Title:    "first post",
			Category: "life",
			Type:     "text",
			Text:     "yes text"},
		{
			Title:    "second post",
			Category: "work",
			Type:     "url",
			URL:      "yes url"},
	}
	author := structs.Author{
		Username: "khasbulat",
		ID:       0,
	}
	for _, data := range dataSlice {
		_, err := posts.CreateNewPost(data, author)
		if err != nil {
			t.Fatalf("Error occured while creating %v test envoronment, can`t proceed: %v", t.Name(), err)
		}
	}
	res, err := posts.GetAllPosts()
	if err != nil {
		t.Errorf("error occured where not expected: %v", err)
	}
	if len(res) != len(dataSlice) {
		t.Errorf("Slice length doesn`t coincide with expected\n"+
			"expected: %v\n"+
			"got: %v\n", len(dataSlice), len(res))
	}
	for i := range res {
		if res[i].Title != dataSlice[i].Title {
			t.Errorf("Titles doesn`t coincide\n"+
				"expected equality of these variables:\n"+
				"dataSlice[%d].Title = %s\n"+
				"res[%d].Title = %s\n", i, dataSlice[i].Title, i, res[i].Title)
		}
		if res[i].Category != dataSlice[i].Category {
			t.Errorf("Categories doesn`t coincide\n"+
				"expected equality of these variables:\n"+
				"dataSlice[%d].Category = %s\n"+
				"res[%d].Category = %s\n", i, dataSlice[i].Category, i, res[i].Category)
		}
		if res[i].Type != dataSlice[i].Type {
			t.Errorf("Types doesn`t coincide\n"+
				"expected equality of these variables:\n"+
				"dataSlice[%d].Type = %s\n"+
				"res[%d].Type = %s\n", i, dataSlice[i].Type, i, res[i].Type)
		}
	}
	_, err = posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}
}

func TestPostsMongo_PostsByCategory(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	maxStringSize := 20
	maxCategories := 20
	maxPosts := 20
	categories := make([]string, 0)
	for i := 0; i < rand.Intn(maxCategories)+1; i++ {
		categories = append(categories, RandString(maxStringSize))
	}
	postsDataByCategories := make([][]structs.NewPostData, 0)
	for _, category := range categories {
		oneCategoryPosts := make([]structs.NewPostData, 0)
		oneCategoryPostsSize := rand.Intn(maxPosts)
		for i := 0; i < oneCategoryPostsSize; i++ {
			postData := structs.NewPostData{
				Title:    RandString(maxStringSize),
				Category: category,
				Type:     "text",
				Text:     RandString(maxStringSize),
			}
			oneCategoryPosts = append(oneCategoryPosts, postData)
		}
		postsDataByCategories = append(postsDataByCategories, oneCategoryPosts)
	}
	for _, oneCatPosts := range postsDataByCategories {
		for _, postData := range oneCatPosts {
			author := structs.Author{
				Username: RandString(maxStringSize),
				ID:       rand.Int(),
			}
			_, err := posts.CreateNewPost(postData, author)
			if err != nil {
				t.Fatalf("Error occured while creating %v test envoronment, can`t proceed: %v", t.Name(), err)
			}
		}
	}
	results := make([][]structs.Post, 0)
	for _, category := range categories {
		result, err := posts.PostsByCategory(category)
		if err != nil {
			t.Errorf("Error didn`t expected, but: %v", err)
		}
		results = append(results, result)
	}
	for i := range postsDataByCategories {
		postsCount := len(results[i])
		dataCount := len(postsDataByCategories[i])
		if postsCount != dataCount {
			t.Errorf("expected posts count of %s category: %d\n"+
				"got : %d", categories[i], dataCount, postsCount)
		}
	}

	for i, posts := range results {
		for j, post := range posts {
			catExp := postsDataByCategories[i][j].Category
			catGot := post.Category
			if catExp != catGot {
				t.Errorf("categories mismatch:\n"+
					"results[%d][%d].Category = %s\n"+
					"postsDataByCategories[%d][%d].Category = %s\n",
					i, j, catGot, i, j, catExp)
			}
			titleExp := postsDataByCategories[i][j].Title
			titleGot := post.Title
			if titleExp != titleGot {
				t.Errorf("titles mismatch:\n"+
					"results[%d][%d].Title = %s\n"+
					"postsDataByCategories[%d][%d].Title = %s\n",
					i, j, titleGot, i, j, titleExp)
			}
			textExp := postsDataByCategories[i][j].Text
			textGot := post.Text
			if textExp != textGot {
				t.Errorf("titles mismatch:\n"+
					"results[%d][%d].Text = %s\n"+
					"postsDataByCategories[%d][%d].Text = %s\n",
					i, j, textGot, i, j, textExp)
			}
		}
	}
	_, err := posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}
}

func TestPostsMongo_PostsByUsername(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	maxStringSize := 20
	maxUsernames := 20
	maxPosts := 20
	usernames := make([]string, 0)
	for i := 0; i < rand.Intn(maxUsernames)+1; i++ {
		usernames = append(usernames, RandString(maxStringSize))
	}

	postsDataByUsers := make([][]structs.NewPostData, 0)
	for range usernames {
		oneUserPosts := make([]structs.NewPostData, 0)
		oneUserPostsSize := rand.Intn(maxPosts)
		for i := 0; i < oneUserPostsSize; i++ {
			postData := structs.NewPostData{
				Title:    RandString(maxStringSize),
				Category: RandString(maxStringSize),
				Type:     "text",
				Text:     RandString(maxStringSize),
			}
			oneUserPosts = append(oneUserPosts, postData)
		}
		postsDataByUsers = append(postsDataByUsers, oneUserPosts)
	}
	for i, oneUserPosts := range postsDataByUsers {
		for _, postData := range oneUserPosts {
			author := structs.Author{
				Username: usernames[i],
				ID:       rand.Int(),
			}
			_, err := posts.CreateNewPost(postData, author)
			if err != nil {
				t.Fatalf("Error occured while creating %v test envoronment, can`t proceed: %v", t.Name(), err)
			}
		}
	}
	results := make([][]structs.Post, 0)
	for _, username := range usernames {
		result, err := posts.PostsByUsername(username)
		if err != nil {
			t.Errorf("Error didn`t expected, but: %v", err)
		}
		results = append(results, result)
	}
	for i := range postsDataByUsers {
		postsCount := len(results[i])
		dataCount := len(postsDataByUsers[i])
		if postsCount != dataCount {
			t.Errorf("expected posts count of %s user: %d\n"+
				"got : %d", usernames[i], dataCount, postsCount)
		}
	}

	for i, posts := range results {
		for j, post := range posts {
			authorExp := usernames[i]
			authorGot := post.Author.Username
			if authorExp != authorGot {
				t.Errorf("usernames mismatch:\n"+
					"results[%d][%d].Author.Username = %s\n"+
					"usernames[%d] = %s\n",
					i, j, authorGot, i, authorExp)
			}
		}
	}
	_, err := posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}
}

func TestPostsMongo_GetPostByID(t *testing.T) {

	maxStringSize := 20
	postData := structs.NewPostData{
		Title:    RandString(maxStringSize),
		Category: RandString(maxStringSize),
		Type:     "text",
		Text:     RandString(maxStringSize),
	}
	author := structs.Author{
		Username: RandString(maxStringSize),
		ID:       0,
	}

	post, err := posts.CreateNewPost(postData, author)

	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}

	getPost, err := posts.GetPostByID(post.ID)

	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}

	if !reflect.DeepEqual(post, getPost) {
		t.Errorf("structs mismatch:"+
			"post: %v"+
			"get post: %v", post, getPost)
	}
	_, err = posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}
}

func TestPostsMongo_DeletePostByID(t *testing.T) {
	maxStringSize := 20
	postData := structs.NewPostData{
		Title:    RandString(maxStringSize),
		Category: RandString(maxStringSize),
		Type:     "text",
		Text:     RandString(maxStringSize),
	}
	author := structs.Author{
		Username: RandString(maxStringSize),
		ID:       0,
	}

	post, err := posts.CreateNewPost(postData, author)
	if err != nil {
		t.Errorf("can`t procceed, %v", err)
	}
	err = posts.DeletePostByID(post.ID, author.ID+rand.Intn(10)+1)
	if err == nil {
		t.Errorf("error expected")
	}

	err = posts.DeletePostByID(post.ID, author.ID)
	if err != nil {
		t.Errorf("can`t procceed, %v", err)
	}
	post, err = posts.GetPostByID(post.ID)
	if err == nil {
		t.Errorf("error expected")
	}

	err = posts.DeletePostByID(post.ID, author.ID)
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}

}
func TestPostsMongo_AddComment(t *testing.T) {
	maxStringSize := 20
	postData := structs.NewPostData{
		Title:    RandString(maxStringSize),
		Category: RandString(maxStringSize),
		Type:     "text",
		Text:     RandString(maxStringSize),
	}
	author := structs.Author{
		Username: RandString(maxStringSize),
		ID:       0,
	}

	post, err := posts.CreateNewPost(postData, author)
	if err != nil {
		t.Errorf("can`t procceed, %v", err)
	}
	comment := structs.Comment{
		ID:      primitive.NewObjectID(),
		Author:  structs.Author{Username: "Jasbi", ID: 2},
		Body:    "some  COMMENT",
		Created: time.Now().Format("2006-01-02T15:04:05.000"),
	}
	res, err := posts.AddComment(comment, post.ID)
	if err != nil {
		t.Errorf("can`t procceed, %v", err)
	}
	if comment.Body != res.Comments[0].Body {
		t.Errorf("expected equality of these values:\n"+
			"1)%v\n"+
			"2)%v", comment.Body, res.Comments[0].Body)
	}
	res, err = posts.AddComment(comment, post.ID)
	if err != nil {
		t.Errorf("can`t procceed, %v", err)
	}
	if comment.Body != res.Comments[0].Body {
		t.Errorf("expected equality of these values:\n"+
			"1)%v\n"+
			"2)%v", comment.Body, res.Comments[0].Body)
	}
	_, err = posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}

}
func TestPostsMongo_DeleteComment(t *testing.T) {
	maxStringSize := 20
	postData := structs.NewPostData{
		Title:    RandString(maxStringSize),
		Category: RandString(maxStringSize),
		Type:     "text",
		Text:     RandString(maxStringSize),
	}
	author := structs.Author{
		Username: RandString(maxStringSize),
		ID:       0,
	}

	post, err := posts.CreateNewPost(postData, author)
	if err != nil {
		t.Errorf("can`t procceed, %v", err)
	}
	comment := structs.Comment{
		ID:      primitive.NewObjectID(),
		Author:  structs.Author{Username: "Jasbi", ID: 2},
		Body:    "some  COMMENT",
		Created: time.Now().Format("2006-01-02T15:04:05.000"),
	}
	res, err := posts.AddComment(comment, post.ID)
	if err != nil {
		t.Errorf("can`t procceed, %v", err)
	}

	updatedPost, err := posts.DeleteComment(res.Comments[0].ID, post.ID, author.ID)
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	if len(updatedPost.Comments) != 0 {
		t.Errorf("expected zero comments")
	}

	_, err = posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}

}

func TestPostsMongo_Vote(t *testing.T) {
	maxStringSize := 20
	postData := structs.NewPostData{
		Title:    RandString(maxStringSize),
		Category: RandString(maxStringSize),
		Type:     "text",
		Text:     RandString(maxStringSize),
	}
	author := structs.Author{
		Username: RandString(maxStringSize),
		ID:       0,
	}
	post, err := posts.CreateNewPost(postData, author)
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	newVotes := []structs.Vote{
		{
			User: 1,
			Vote: 1},
		{
			User: 2,
			Vote: 1,
		},
	}
	post, err = posts.Vote(post.ID, newVotes[0])
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	if post.Votes[0].User != newVotes[0].User {
		t.Errorf("expected equality of users")
	}
	_, err = posts.Vote(post.ID, newVotes[1])
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	_, err = posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}
}

func TestPostsMongo_UnVote(t *testing.T) {
	maxStringSize := 20
	postData := structs.NewPostData{
		Title:    RandString(maxStringSize),
		Category: RandString(maxStringSize),
		Type:     "text",
		Text:     RandString(maxStringSize),
	}
	author := structs.Author{
		Username: RandString(maxStringSize),
		ID:       0,
	}
	post, err := posts.CreateNewPost(postData, author)
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	newVotes := []structs.Vote{
		{
			User: 1,
			Vote: 1},
		{
			User: 2,
			Vote: 1,
		},
	}
	post, err = posts.Vote(post.ID, newVotes[0])
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	if post.Score != 1 {
		t.Errorf("expected score = 1")
	}
	post, err = posts.Vote(post.ID, newVotes[1])
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	if post.Score != 2 {
		t.Errorf("expected score = 2")
	}
	if post.UpvotePercentage != 100 {
		t.Errorf("expected UpvotePercentage = 100")
	}

	post, err = posts.UnVote(post.ID, newVotes[0].User)
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	if len(post.Votes) != 1 {
		t.Errorf("expected one vote")
	}
	newVotes[1].Vote = -1
	post, err = posts.Vote(post.ID, newVotes[1])
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	post, err = posts.UnVote(post.ID, newVotes[1].User)
	if err != nil {
		t.Fatalf("can`t procceed, %v", err)
	}
	if len(post.Votes) != 0 {
		t.Errorf("expected 0 votes")
	}
	_, err = posts.posts.DeleteMany(nil, bson.M{})
	if err != nil {
		t.Fatalf("exit point of test %v ended with error %v", t.Name(), err)
	}
}
