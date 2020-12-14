package database

type Database interface {
	Collection(name string) (Collection, error)
}

type Collection interface {
	CreateUser(data User) (*User, error)
	AddBookToMyCollection(loggedUserID int, data Book) (*Book, error)
	LendBook(loggedUserId int, data BookLoan) (*BookLoan, error)
	ReturnBook(loggedUserId int, bookId int) (*BookLoan, error)
}
