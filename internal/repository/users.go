package repository

import (
	"fmt"
	"go-library-manager/internal/database"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/models"
	"gorm.io/gorm"
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
	return
}

func ExistsUserById(id int) bool {
	var count int64
	database.DB.Model(&models.User{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func findUserListContent(query string, pageNumber, pageSize int, contentChan chan []dtos.UserListDto) (result []dtos.UserListDto) {
	database.DB.
		Model(&models.User{}).
		Joins("Profile").
		Where(`"user".id > 1 and "user".active is true and "user".name ilike ?`, "%"+query+"%").
		Limit(pageSize).
		Offset(pageNumber * pageSize).
		Select(`"user".id, "user".login, "user".name, "Profile".description, "Profile".birth_date, "Profile".class, "Profile".shift`).
		Scan(&result)
	contentChan <- result
	return
}

func findUserListCount(query string, countChan chan int) (count int64) {
	database.DB.
		Model(&models.User{}).
		Where(`"user".id > 1 and "user".active is true and "user".name ilike ?`, "%"+query+"%").
		Count(&count)
	countChan <- int(count)
	return
}

func FindUserById(id int) (user models.User) {
	database.DB.Model(&models.User{}).Preload("Roles").Preload("Profile").First(&user, id)
	user.Password = ""
	return
}

func UpdateUser(m *models.User) {
	database.DB.Model(&models.User{}).Where("id = ?", m.ID).Update("name", m.Name)
	database.DB.Model(&models.Profile{}).Where("id = ?", m.Profile.ID).
		Update("name", m.Profile.Name).
		Update("description", m.Profile.Description).
		Update("document", m.Profile.Document).
		Update("birth_date", m.Profile.BirthDate).
		Update("class", m.Profile.Class).
		Update("shift", m.Profile.Shift).
		Update("address", m.Profile.Address).
		Update("number", m.Profile.Number).
		Update("neighborhood", m.Profile.Neighborhood).
		Update("phone", m.Profile.Phone)
}

func UpdateUserTx(user *models.User) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Update user and profile in a single transaction
		if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(user).Error; err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		// Update roles association separately as it's a many-to-many relationship
		if err := tx.Model(user).Association("Roles").Replace(user.Roles); err != nil {
			return fmt.Errorf("failed to update roles: %w", err)
		}

		return nil
	})
}
