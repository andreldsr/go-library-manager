package models

import "time"

type Lending struct {
	ID         int  `json:"id,omitempty"`
	User       User `json:"user"`
	UserId     int
	Book       Book `json:"book"`
	BookId     int
	ReturnDate time.Time `json:"returnDate,omitempty" gorm:"default:null"`
	ReturnedAt time.Time `json:"returnedAt,omitempty" gorm:"default:null"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (Lending) TableName() string {
	return "lending"
}
