package models

type Publisher struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Books []Book `json:"books,omitempty"`
}

func (Publisher) TableName() string {
	return "publisher"
}
