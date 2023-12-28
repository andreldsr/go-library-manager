package models

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
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Document    string `json:"document,omitempty"`
}

func (Profile) TableName() string {
	return "profile"
}
