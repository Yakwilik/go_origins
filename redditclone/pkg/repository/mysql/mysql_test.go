package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository/mysql/authmysql"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"testing"
)

func TestNewMySqlDB(t *testing.T) {
	cfgToken := Config{
		Host:     "localhost",
		Port:     "3301",
		Username: "service_admin",
		Password: "qwerty",
		DBName:   "redditclone?&charset=utf8&interpolateParams=true",
	}
	dbToken, err := NewMySQLDB(cfgToken)

	if err != nil {
		fmt.Println(err)
	}
	err = dbToken.Ping()

	if err != nil {
		fmt.Println(err)
	}

	cfgUsers := Config{
		Host:     "localhost",
		Port:     "3302",
		Username: "service_admin",
		Password: "qwerty",
		DBName:   "redditclone?&charset=utf8&interpolateParams=true",
	}
	dbUsers, err := NewMySQLDB(cfgUsers)
	if err != nil {
		fmt.Println(err)
	}

	err = dbUsers.Ping()
	if err != nil {
		fmt.Println(err)
	}
	auth := authmysql.NewAuthMySQL(dbToken, dbUsers)

	id, err := auth.CreateUser(structs.User{
		ID:       0,
		Username: "khasbulat",
		Password: "sdfasdf",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	user, err := auth.GetUser("khasbulat", "sdfasdf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

}
