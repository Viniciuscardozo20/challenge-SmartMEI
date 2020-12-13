package database

import "time"

type User struct {
	Id            int        `json:"id"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	CreatedAt     time.Time  `json:"pages"`
	Collection    []Book     `json:"collection"`
	LentBooks     []BookLoan `json:"lentBooks"`
	BorrowedBooks []BookLoan `json:"borrowedBooks"`
}
