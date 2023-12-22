package models

type Author struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Books []Book `json:"books,omitempty" gorm:"many2many:books_authors"`
}

func (Author) TableName() string {
	return "author"
}
