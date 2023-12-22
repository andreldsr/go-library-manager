package repository

import (
	"go-library-manager/internal/database"
	"go-library-manager/internal/models"
)

func FindOrCreatePublisher(name string) models.Publisher {
	publisher := models.Publisher{Name: name}
	database.DB.FirstOrCreate(&publisher, publisher)
	return publisher
}
