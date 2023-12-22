package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err.Error())
	}
	if err != nil {
		panic(err.Error())
	}
}
