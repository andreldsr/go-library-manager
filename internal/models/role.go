package models

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name,omitempty"`
}

func (Role) TableName() string {
	return "role"
}
