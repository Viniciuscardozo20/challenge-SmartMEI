package database

type Database interface {
	Collection(name string) (Collection, error)
}

type Collection interface {
	CreateUser(data User) (interface{}, error)
	AddBookToMyCollection(loggedUserID int, data Book) (interface{}, error)
	LendBook(loggedUserId int, data Book) (interface{}, error)
	ReturnBook(loggedUserId int, bookId int) (*BookLoan, error)
}
