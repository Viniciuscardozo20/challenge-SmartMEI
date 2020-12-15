package models

import "time"

type Book struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string    `json:"title"`
	Pages     string    `json:"pages"`
	CreatedAt time.Time `json:"createdAt"`
}
