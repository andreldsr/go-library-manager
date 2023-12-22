package models

import (
	"time"
)

type Book struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	Title            string    `json:"title"`
	Volume           string    `json:"volume"`
	Copy             string    `json:"copy"`
	Location         string    `json:"location"`
	PublicationYear  int       `json:"publicationYear"`
	AcquisitionForm  string    `json:"acquisitionForm"`
	Index            string    `json:"index"`
	CDD              string    `json:"cdd"`
	Observation      string    `json:"observation"`
	RegisterNumber   int       `json:"registerNumber"`
	RegistrationDate time.Time `json:"registrationDate"`
	Lending          *Lending  `json:"lending"`
	LendingId        int       `gorm:"default:null"`
	PublisherId      int
	Publisher        Publisher `json:"publisher,omitempty"`
	Authors          []Author  `json:"authors" gorm:"many2many:books_authors"`
}

func (Book) TableName() string {
	return "book"
}
