package dtos

import "time"

type UserListDto struct {
	Id          int       `json:"id"`
	Login       string    `json:"login"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	BirthDate   time.Time `json:"birth_date,omitempty"`
	Class       string    `json:"class"`
	Shift       string    `json:"shift"`
}

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserDto struct {
	Login        string `json:"login"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	Document     string `json:"document"`
	Description  string `json:"description"`
	BirthDate    string `json:"birth_date"`
	Class        string `json:"class"`
	Shift        string `json:"shift"`
	Address      string `json:"address"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	Phone        string `json:"phone"`
}
