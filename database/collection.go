package database

import (
	. "challenge-SmartMEI/database/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collection struct {
	coll *mongo.Collection
}

func newCollection(name string, db mongo.Database) (*collection, error) {
	names, err := db.ListCollectionNames(nil, bson.D{{}})
	if err != nil {
		return nil, err
	}
	var exists = false
	for _, n := range names {
		if n == name {
			exists = true
		}
	}
	var coll *mongo.Collection
	if exists != true {
		coll = db.Collection(name, nil)
		_, err = coll.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: "email", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)
		if err != nil {
			return nil, err
		}
		_, err = coll.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: "id", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)
		if err != nil {
			return nil, err
		}
	} else {
		coll = db.Collection(name, nil)
	}
	return &collection{
		coll: coll,
	}, nil
}

func (c *collection) CreateUser(data User) (*User, error) {
	data.CreatedAt = time.Now()
	result, err := c.coll.InsertOne(nil, data)
	if err != nil {
		return nil, err
	}
	user, err := c.findUser(result.InsertedID, "_id")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *collection) AddBookToMyCollection(user User, data Book) (*Book, error) {
	user.Collection = append(user.Collection, data)
	result := c.coll.FindOneAndUpdate(nil, bson.M{"email": user.Email}, bson.M{"$set": user})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var userR User
	err := result.Decode(&userR)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *collection) LendBook(fromUser User, toUser User, data BookLoan) (*BookLoan, error) {
	fromUser.LentBooks = append(fromUser.LentBooks, data)
	toUser.BorrowedBooks = append(toUser.BorrowedBooks, data)

	result := c.coll.FindOneAndUpdate(nil, bson.M{"id": fromUser.Id}, bson.M{"$set": fromUser})
	if result.Err() != nil {
		return nil, result.Err()
	}
	result = c.coll.FindOneAndUpdate(nil, bson.M{"id": toUser.Id}, bson.M{"$set": toUser})
	if result.Err() != nil {
		return nil, result.Err()
	}

	return &data, nil
}

func (c *collection) ReturnBook(user User, fromUser User) error {
	result := c.coll.FindOneAndUpdate(nil, bson.M{"id": user.Id}, bson.M{"$set": user})
	if result.Err() != nil {
		return result.Err()
	}
	result = c.coll.FindOneAndUpdate(nil, bson.M{"id": fromUser.Id}, bson.M{"$set": fromUser})
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (c *collection) GetUserDetails(loggedUserId int) (*User, error) {
	user, err := c.findUser(loggedUserId, "id")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *collection) findUser(input interface{}, field string) (*User, error) {
	result := c.coll.FindOne(nil, bson.M{field: input})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var user User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
