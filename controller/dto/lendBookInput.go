package dto

type LendBookInput struct {
	BookId   int `json:"bookId"`
	ToUserId int `json:"toUserId"`
}
