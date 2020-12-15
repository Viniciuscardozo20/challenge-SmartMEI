package dto

type LendBookInput struct {
	BookId   string `json:"bookId"`
	ToUserId string `json:"toUserId"`
}
