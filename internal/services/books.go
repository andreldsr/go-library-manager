package services

import (
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/models"
	"go-library-manager/internal/repository"
	"strconv"
	"strings"
	"time"
)

func FindAllBooks(query string, page, size int) dtos.Page[dtos.BookListDto] {
	return repository.FindAllBooks(query, page, size)
}

func FindBookById(id int) models.Book {
	return repository.FindBookById(id)
}

func GetBookStats() dtos.BookStats {
	return repository.GetStats()
}

func CreateBook(dto dtos.CreateBookDto) models.Book {
	registerNumber, err := strconv.Atoi(dto.RegisterNumber)
	if err != nil {
		panic(err.Error())
		return models.Book{}
	}
	if repository.ExistsByCode(registerNumber) {
		return models.Book{}
	}
	authors := getAuthors(dto)
	publisher := repository.FindOrCreatePublisher(dto.Publisher)
	book := buildBook(dto, registerNumber, publisher, authors)
	repository.CreateBook(&book)
	return book
}

func buildBook(dto dtos.CreateBookDto, registerNumber int, publisher models.Publisher, authors []models.Author) models.Book {
	book := models.Book{
		Title:            dto.Title,
		Volume:           dto.Volume,
		Copy:             dto.Copy,
		Location:         dto.Location,
		PublicationYear:  dto.PublicationYear,
		AcquisitionForm:  dto.AcquisitionForm,
		Index:            dto.Index,
		CDD:              dto.Cdd,
		Observation:      dto.Observation,
		RegisterNumber:   registerNumber,
		RegistrationDate: time.Now(),
		Publisher:        publisher,
		Authors:          authors,
	}
	return book
}

func getAuthors(dto dtos.CreateBookDto) []models.Author {
	authorsNames := strings.Split(dto.Authors, " e ")
	var authors []models.Author
	for _, authorName := range authorsNames {
		author := repository.FindOrCreateAuthor(authorName)
		authors = append(authors, author)
	}
	return authors
}
