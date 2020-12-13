package database

import "time"

type BookLoan struct {
	Id         int       `json:"id"`
	Book       Book      `json:"book"`
	FromUser   int       `json:"fromUser"`
	ToUser     int       `json:"toUser"`
	LentAt     time.Time `json:"lentAt"`
	ReturnedAt time.Time `json:"returnedAt"`
	Returned   bool      `json:"returned"`
}
