package models

import "time"

type BookLoan struct {
	Book       Book      `json:"book"`
	FromUser   string    `json:"fromUser"`
	ToUser     string    `json:"toUser"`
	LentAt     time.Time `json:"lentAt"`
	ReturnedAt time.Time `json:"returnedAt"`
	Returned   bool      `json:"returned"`
}
