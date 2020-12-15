package controller

import (
	"challenge-SmartMEI/controller/dto"
	"challenge-SmartMEI/database"
	"challenge-SmartMEI/database/models"
	"errors"
	"time"
)

type Controller struct {
	coll database.Collection
}

func NewController(coll database.Collection) *Controller {
	return &Controller{
		coll: coll,
	}
}

func (c *Controller) CreateUser(input dto.CreateUserInput) (*models.User, error) {
	var userInput = models.User{
		Name:          input.Name,
		Email:         input.Email,
		Collection:    make([]models.Book, 0),
		LentBooks:     make([]models.BookLoan, 0),
		BorrowedBooks: make([]models.BookLoan, 0),
	}
	user, err := c.coll.CreateUser(userInput)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Controller) GetUserDetails(loggedUserId string) (*models.User, error) {
	user, err := c.coll.GetUserDetails(loggedUserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Controller) AddBookToMyCollection(loggedUserId string, input dto.AddBookInput) (*models.Book, error) {
	user, err := c.coll.GetUserDetails(loggedUserId)
	if err != nil {
		return nil, err
	}
	var bookInput = models.Book{
		Title:     input.Title,
		Pages:     input.Pages,
		CreatedAt: time.Now(),
	}
	book, err := c.coll.AddBookToMyCollection(*user, bookInput)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (c *Controller) LendBook(loggedUserId string, input dto.LendBookInput) (*models.BookLoan, error) {
	user, err := c.coll.GetUserDetails(loggedUserId)
	if err != nil {
		return nil, err
	}
	toUser, err := c.coll.GetUserDetails(input.ToUserId)
	if err != nil {
		return nil, err
	}
	exists := false
	var book models.Book
	for _, bookr := range user.Collection {
		if bookr.Id == input.BookId {
			exists = true
			book = bookr
		}
	}
	if exists != true {
		return nil, errors.New("You don't have this book")
	}
	for _, lbook := range user.LentBooks {
		if lbook.Book.Id == input.BookId {
			if lbook.Returned != true {
				return nil, errors.New("Book's already loan.")
			}
		}
	}
	var lendInput = models.BookLoan{
		Book:     book,
		FromUser: user.Id,
		ToUser:   toUser.Id,
		LentAt:   time.Now(),
		Returned: false,
	}
	bookLoan, err := c.coll.LendBook(*user, *toUser, lendInput)
	if err != nil {
		return nil, err
	}
	return bookLoan, nil
}

func (c *Controller) ReturnBook(loggedUserId string, bookId string) (*models.BookLoan, error) {
	user, err := c.coll.GetUserDetails(loggedUserId)
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
	var bookLoan models.BookLoan
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
	fromUser, err := c.coll.GetUserDetails(bookLoan.FromUser)
	if err != nil {
		return nil, err
	}
	for i, lbook := range fromUser.LentBooks {
		if lbook.Book.Id == bookId {
			fromUser.LentBooks[i] = bookLoan
		}
	}

	err = c.coll.ReturnBook(*user, *fromUser)
	if err != nil {
		return nil, err
	}
	return &bookLoan, nil
}
