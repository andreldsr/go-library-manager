package models

import "time"

type User struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name,omitempty"`
	Login     string  `json:"login,omitempty"`
	Password  string  `json:"password,omitempty"`
	Active    bool    `json:"active,omitempty"`
	Profile   Profile `json:"profile"`
	Roles     []Role  `json:"roles" gorm:"many2many:user_roles"`
	ProfileId int
}

func (User) TableName() string {
	return "user"
}

type Profile struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Document     string    `json:"document,omitempty"`
	BirthDate    time.Time `json:"birth_date,omitempty"`
	Class        string    `json:"class,omitempty"`
	Shift        string    `json:"shift,omitempty"`
	Address      string    `json:"address,omitempty"`
	Number       string    `json:"number,omitempty"`
	Neighborhood string    `json:"neighborhood,omitempty"`
	Phone        string    `json:"phone,omitempty"`
}

func (Profile) TableName() string {
	return "profile"
}
