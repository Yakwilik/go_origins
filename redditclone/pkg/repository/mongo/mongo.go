package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

func NewMongoDB(cfg Config) (*mongo.Client, error) {
	client, err := mongo.Connect(nil,
		options.Client().
			ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)).
			SetAuth(options.
				Credential{
				AuthMechanism:           "",
				AuthMechanismProperties: nil,
				AuthSource:              "",
				Username:                cfg.Username,
				Password:                cfg.Password,
				PasswordSet:             false,
			}))

	return client, err
}
