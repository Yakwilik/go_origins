package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/handler"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository/mongo"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository/mysql"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/service"
	"time"
)

var cfgToken = mysql.Config{
	Host:     "localhost",
	Port:     "3301",
	Username: "service_admin",
	Password: "qwerty",
	DBName:   "redditclone?&charset=utf8&interpolateParams=true",
}

var cfgUsers = mysql.Config{
	Host:     "localhost",
	Port:     "3302",
	Username: "service_admin",
	Password: "qwerty",
	DBName:   "redditclone?&charset=utf8&interpolateParams=true",
}

var cfgPosts = mongo.Config{
	Host:     "localhost",
	Port:     "27017",
	Username: `root`,
	Password: `example`,
}

var configs = repository.Configs{
	UsersDBConfig:       cfgUsers,
	SessionDBConfig:     cfgToken,
	PostsDBConfig:       cfgPosts,
	MongoDBName:         "coursera",
	MongoCollectionName: "posts",
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("Ошибка в прочтении файла конфигурации: %s", err.Error())
	}

	repos := repository.NewMySQLMongoRepository(configs)
	timer := time.NewTicker(1 * time.Second)
	fmt.Println("begin")
	for range timer.C {
		if repos != nil {
			fmt.Println("repos initialized")
			break
		}
		fmt.Println("waiting for DBs initialization")
		repos = repository.NewMySQLMongoRepository(configs)
	}
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(redditclone.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("произошла ошибка во время запуска http-сервера: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
