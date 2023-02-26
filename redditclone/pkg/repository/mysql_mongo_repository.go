package repository

import (
	"fmt"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository/mongo"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository/mongo/postsmongo"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository/mysql"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository/mysql/authmysql"
	"log"
)

type Configs struct {
	UsersDBConfig       mysql.Config
	SessionDBConfig     mysql.Config
	PostsDBConfig       mongo.Config
	MongoDBName         string
	MongoCollectionName string
}

func NewMySQLMongoRepository(cfg Configs) *Repository {
	usersDB, err := mysql.NewMySQLDB(cfg.UsersDBConfig)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = usersDB.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}
	sessionsDB, err := mysql.NewMySQLDB(cfg.SessionDBConfig)
	if err != nil {
		return nil
	}
	err = sessionsDB.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}
	postsDB, err := mongo.NewMongoDB(cfg.PostsDBConfig)
	if err != nil {
		return nil
	}
	err = postsDB.Ping(nil, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Repository{
		Authorization: authmysql.NewAuthMySQL(sessionsDB, usersDB),
		Posts:         postsmongo.NewPostsMongo(postsDB, cfg.MongoDBName, cfg.MongoCollectionName),
	}
}
