package database

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type collection struct {
	coll *mongo.Collection
}

func newCollection(name string, db mongo.Database) (*collection, error) {
	names, err := db.ListCollectionNames(nil, nil)
	if err != nil {
		return nil, err
	}
	exist := false
	for _, collName := range names {
		if collName == name {
			exist = true
			break
		}
	}
	if !exist {
		err := db.CreateCollection(nil, name)
		if err != nil {
			return nil, err
		}
	}
	coll := db.Collection(name, nil)
	return &collection{
		coll: coll,
	}, nil
}

func (c *collection) CreateUser(data User) (interface{}, error) {
	result, err := c.coll.InsertOne(nil, data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (c *collection) AddBookToMyCollection(loggedUserId int, data Book) (interface{}, error) {
	err := c.loggedUserbyID(loggedUserId)
	if err != nil {
		return nil, err
	}
	result, err := c.coll.InsertOne(nil, data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (c *collection) LendBook(loggedUserId int, data Book) (interface{}, error) {
	err := c.loggedUserbyID(loggedUserId)
	if err != nil {
		return nil, err
	}
	_, err = c.borrowedBook(data.Id)
	if err != nil {
		return nil, err
	}
	inserted, err := c.coll.InsertOne(nil, data)
	if err != nil {
		return nil, err
	}
	return inserted.InsertedID, nil
}

func (c *collection) ReturnBook(loggedUserId int, bookId int) (*BookLoan, error) {
	err := c.loggedUserbyID(loggedUserId)
	if err != nil {
		return nil, err
	}
	bookLoan, err := c.borrowedBook(bookId)
	if err != nil {
		return nil, err
	}
	result := c.coll.FindOneAndUpdate(nil, bookLoan.Id, bookLoan)
	if result.Err() != nil {
		return nil, result.Err()
	}
	err = result.Decode(&bookLoan)
	if err != nil {
		return nil, err
	}
	return bookLoan, nil
}

func (c *collection) loggedUserbyID(id int) error {
	result := c.coll.FindOne(nil, id)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (c *collection) borrowedBook(id int) (*BookLoan, error) {
	result := c.coll.FindOne(nil, id)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var bookLoan BookLoan
	err := result.Decode(&bookLoan)
	if err != nil {
		return nil, err
	}
	if bookLoan.Returned != true {
		return nil, errors.New("book's already loan.")
	}
	return &bookLoan, nil
}
