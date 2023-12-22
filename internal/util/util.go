package util

import (
	"github.com/golang-jwt/jwt"
	"go-library-manager/internal/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

var JWT_SECRET_KEY = "TOKEN_SECRET"

func IntOrDefault(value string, def int) int {
	convertedValue, err := strconv.Atoi(value)
	if err != nil || convertedValue == 0 {
		return def
	}
	return convertedValue
}

func CreateJwt(user *models.User) (tokenString string, err error) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":          user.ID,
		"name":        user.Name,
		"sub":         user.Name,
		"authorities": user.Roles,
		"profile":     user.Profile.Description,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func EncryptPassword(text string) (encrypted string, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func CheckPassword(password string, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	return err == nil
}
