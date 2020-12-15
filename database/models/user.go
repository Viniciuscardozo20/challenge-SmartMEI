package models

import "time"

type User struct {
	Id            string     `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string     `json:"name" bson:"name"`
	Email         string     `json:"email" bson:"email"`
	CreatedAt     time.Time  `json:"createdAt" bson:"createdAt"`
	Collection    []Book     `json:"collection" bson:"collection"`
	LentBooks     []BookLoan `json:"lentBooks" bson:"lentBooks"`
	BorrowedBooks []BookLoan `json:"borrowedBooks" bson:"borrowedBooks"`
}
