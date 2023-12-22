package dtos

type UserListDto struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserDto struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	Description string `json:"description"`
}
