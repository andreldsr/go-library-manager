package repository

import (
	"go-library-manager/internal/database"
	"go-library-manager/internal/models"
)

func FindAllAuthors(query string, pageNumber, pageSize int) []models.Author {
	var result []models.Author
	database.DB.
		Preload("Books").
		Limit(pageSize).
		Offset(pageSize*pageNumber).
		Where("name ilike ?", "%"+query+"%").
		Find(&result)
	return result
}

func FindOrCreateAuthor(name string) models.Author {
	author := models.Author{Name: name}
	database.DB.FirstOrCreate(&author, author)
	return author
}
