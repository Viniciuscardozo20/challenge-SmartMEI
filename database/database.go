package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type driver struct {
	db *mongo.Database
}

func NewDatabase(config Config) (Database, error) {
	client, err := mongo.Connect(nil, options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)).
		SetAuth(options.Credential{
			Username: config.User,
			Password: config.Password,
		}),
	)
	if err != nil {
		return nil, err
	}
	return &driver{db: client.Database(config.Database)}, nil
}

func (d *driver) Collection(name string) (Collection, error) {
	return newCollection(name, *d.db)
}
