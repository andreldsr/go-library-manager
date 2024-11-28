package dtos

import "time"

type LendingListDto struct {
	ID         int       `json:"id,omitempty"`
	BookId     int       `json:"bookId,omitempty"`
	BookTitle  string    `json:"bookTitle,omitempty"`
	UserName   string    `json:"userName,omitempty"`
	ReturnDate time.Time `json:"returnDate"`
	ReturnedAt time.Time `json:"returnedAt"`
	CreatedAt  time.Time `json:"createdAt"`
}
type LendingDetailDto struct {
	Id              int       `json:"id"`
	BookTitle       string    `json:"bookTitle"`
	BookCopy        string    `json:"bookCopy"`
	BookLocation    string    `json:"bookLocation"`
	BookObservation string    `json:"bookObservation"`
	ReturnedAt      string    `json:"returnedAt"`
	BookId          int       `json:"bookId"`
	ReturnDate      string    `json:"returnDate"`
	UserName        string    `json:"userName"`
	CreatedAt       time.Time `json:"createdAt"`
}

type CreateLendingDto struct {
	UserId     int    `json:"userId"`
	BookId     int    `json:"bookId"`
	ReturnDate string `json:"returnDate"`
}
