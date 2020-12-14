package database

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type collection struct {
	coll *mongo.Collection
}

func newCollection(name string, db mongo.Database) (*collection, error) {
	coll := db.Collection(name, nil)
	_, err := coll.InsertOne(nil, bson.M{"Test": "init"})
	if err != nil {
		return nil, err
	}
	return &collection{
		coll: coll,
	}, nil
}

func (c *collection) CreateUser(data User) (*User, error) {
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

func (c *collection) AddBookToMyCollection(loggedUserId int, data Book) (*Book, error) {
	user, err := c.findUser(loggedUserId, "id")
	if err != nil {
		return nil, err
	}
	user.Collection = append(user.Collection, data)
	result := c.coll.FindOneAndUpdate(nil, bson.M{"id": loggedUserId}, bson.M{"$set": user})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var userR User
	err = result.Decode(&userR)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *collection) LendBook(loggedUserId int, data BookLoan) (*BookLoan, error) {
	user, err := c.findUser(loggedUserId, "id")
	if err != nil {
		return nil, err
	}
	toUser, err := c.findUser(data.ToUser, "id")
	if err != nil {
		return nil, err
	}
	exists := false
	for _, book := range user.Collection {
		if book.Id == data.Book.Id {
			exists = true
		}
	}
	if exists != true {
		return nil, errors.New("You don't have this book")
	}
	for _, lbook := range user.LentBooks {
		if lbook.Book.Id == data.Book.Id {
			if lbook.Returned != true {
				return nil, errors.New("Book's already loan.")
			}
		}
	}
	user.LentBooks = append(user.LentBooks, data)
	toUser.BorrowedBooks = append(toUser.BorrowedBooks, data)

	result := c.coll.FindOneAndUpdate(nil, bson.M{"id": loggedUserId}, bson.M{"$set": user})
	if result.Err() != nil {
		return nil, result.Err()
	}
	result = c.coll.FindOneAndUpdate(nil, bson.M{"id": toUser.Id}, bson.M{"$set": toUser})
	if result.Err() != nil {
		return nil, result.Err()
	}

	return &data, nil
}

func (c *collection) ReturnBook(loggedUserId int, bookId int) (*BookLoan, error) {
	user, err := c.findUser(loggedUserId, "id")
	if err != nil {
		return nil, err
	}
	exists := false
	for _, book := range user.BorrowedBooks {
		if book.Book.Id == bookId {
			exists = true
		}
	}
	if exists != true {
		return nil, errors.New("You don't have this book")
	}
	var bookLoan BookLoan
	for i, lbook := range user.BorrowedBooks {
		if lbook.Book.Id == bookId {
			if lbook.Returned == true {
				return nil, errors.New("Book's already returned.")
			}
			user.BorrowedBooks[i].ReturnedAt = time.Now()
			user.BorrowedBooks[i].Returned = true
			bookLoan = user.BorrowedBooks[i]
		}
	}
	fromUser, err := c.findUser(bookLoan.FromUser, "id")
	if err != nil {
		return nil, err
	}

	for i, lbook := range fromUser.LentBooks {
		if lbook.Book.Id == bookId {
			fromUser.LentBooks[i] = bookLoan
		}
	}
	result := c.coll.FindOneAndUpdate(nil, bson.M{"id": user.Id}, bson.M{"$set": user})
	if result.Err() != nil {
		return nil, result.Err()
	}
	result = c.coll.FindOneAndUpdate(nil, bson.M{"id": fromUser.Id}, bson.M{"$set": fromUser})
	if result.Err() != nil {
		return nil, result.Err()
	}

	return &bookLoan, nil
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
