package database

import "time"

type Book struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Pages     string    `json:"pages"`
	CreatedAt time.Time `json:"createdAt"`
}
