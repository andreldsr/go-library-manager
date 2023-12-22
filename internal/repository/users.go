package repository

import (
	"go-library-manager/internal/database"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/models"
)

func FindUserList(query string, pageNumber, pageSize int) dtos.Page[dtos.UserListDto] {
	contentChan := make(chan []dtos.UserListDto)
	countChan := make(chan int)
	go findUserListContent(query, pageNumber, pageSize, contentChan)
	go findUserListCount(query, countChan)
	return dtos.BuildPage(<-contentChan, <-countChan, pageNumber, pageSize)
}

func FindUserByIdAsync(id int, c chan models.User) {
	user := models.User{ID: id}
	database.DB.Find(&user)
	c <- user
}

func CreateUser(user *models.User) {
	database.DB.Create(user)
}

func FindUserByLogin(login string) (user models.User) {
	database.DB.
		Preload("Roles").
		Joins("Profile").
		Find(&user, "login = ?", login)
	if err := database.DB.Preload("Roles").Error; err != nil {
		panic(err.Error())
	}
	return
}

func ExistsUserByLogin(login string) bool {
	var count int64
	database.DB.Model(&models.User{}).Where("login = ?", login).Count(&count)
	return count > 0
}

func findUserListContent(query string, pageNumber, pageSize int, contentChan chan []dtos.UserListDto) (result []dtos.UserListDto) {
	database.DB.
		Model(&models.User{}).
		Joins("Profile").
		Where(`"user".name ilike ?`, "%"+query+"%").
		Limit(pageSize).
		Offset(pageNumber * pageSize).
		Select(`"user".id, "user".login, "user".name, "Profile".description`).
		Scan(&result)
	contentChan <- result
	return
}

func findUserListCount(query string, countChan chan int) (count int64) {
	database.DB.
		Model(&models.User{}).
		Where(`"user".name ilike ?`, "%"+query+"%").
		Count(&count)
	countChan <- int(count)
	return
}
