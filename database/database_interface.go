package database

import (
	. "challenge-SmartMEI/database/models"
)

type Database interface {
	Collection(name string) (Collection, error)
}

type Collection interface {
	CreateUser(data User) (*User, error)
	AddBookToMyCollection(user User, data Book) (*Book, error)
	LendBook(fromUser User, toUser User, data BookLoan) (*BookLoan, error)
	ReturnBook(user User, fromUser User) error
	GetUserDetails(loggedUserId int) (*User, error)
}
