package services

import (
	"errors"
	"fmt"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/models"
	"go-library-manager/internal/repository"
	"go-library-manager/internal/util"
	"time"
)

func FindUsersByName(query string, pageNumber, pageSize int) dtos.Page[dtos.UserListDto] {
	return repository.FindUserList(query, pageNumber, pageSize)
}

func FindUserById(id int) models.User {
	return repository.FindUserById(id)
}

func Login(dto dtos.UserLoginDto) (string, error) {
	user := repository.FindUserByLogin(dto.Username)
	if user.ID == 0 {
		return "", errors.New("error trying to log in")
	}
	checkPassword := util.CheckPassword(dto.Password, user.Password)
	if !checkPassword {
		return "", errors.New("error trying to log in")
	}
	token, err := util.CreateJwt(&user)
	return token, err
}

func CreateUser(dto dtos.CreateUserDto) (err error) {
	birthDate, err := time.Parse("2006-01-02", dto.BirthDate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var user models.User
	roles := repository.FindRolesByName([]string{dto.Role, "ROLE_USER"})
	user = buildUser(dto, roles, birthDate)
	repository.CreateUser(&user)
	if user.ID == 0 {
		err = errors.New("error creating new user")
	}
	return
}

func buildUser(dto dtos.CreateUserDto, roles []models.Role, birthDate time.Time) models.User {
	password, err := util.EncryptPassword(dto.Password)
	if err != nil {
		return models.User{}
	}
	return models.User{
		Roles:    roles,
		Login:    dto.Login,
		Password: password,
		Profile: models.Profile{
			Description:  dto.Description,
			Document:     dto.Document,
			Name:         dto.Name,
			BirthDate:    birthDate,
			Class:        dto.Class,
			Shift:        dto.Shift,
			Address:      dto.Address,
			Number:       dto.Number,
			Neighborhood: dto.Neighborhood,
			Phone:        dto.Phone,
		},
		Active: true,
		Name:   dto.Name,
	}
}

// Custom error types
var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidDate  = errors.New("invalid birth date")
)

func UpdateUser(id int, dto dtos.CreateUserDto) error {
	if id <= 0 {
		return fmt.Errorf("%w: invalid id", ErrInvalidInput)
	}

	existingUser := repository.FindUserById(id)
	if existingUser.ID == 0 {
		return ErrUserNotFound
	}

	birthDate, err := time.Parse("2006-01-02", dto.BirthDate)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidDate, err)
	}

	roles := repository.FindRolesByName([]string{dto.Role, "ROLE_USER"})
	if len(roles) == 0 {
		return fmt.Errorf("%w: invalid role", ErrInvalidInput)
	}

	user := buildUser(dto, roles, birthDate)
	user.ID = id
	user.Profile.ID = id

	if err := repository.UpdateUserTx(&user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func DeleteUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("%w: invalid id", ErrInvalidInput)
	}

	existingUser := repository.FindUserById(id)
	if existingUser.ID == 0 {
		return ErrUserNotFound
	}
	existingUser.Active = false
	if err := repository.UpdateUserTx(&existingUser); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
