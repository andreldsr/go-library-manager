package services

import (
	"errors"
	"fmt"
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
	authors := getAuthors(dto.Authors)
	publisher := repository.FindOrCreatePublisher(dto.Publisher)
	book := buildBook(dto, registerNumber, publisher, authors)
	repository.CreateBook(&book)
	return book
}

// Custom error types for better error handling
var (
	ErrBookNotFound      = errors.New("book not found")
	ErrInvalidInput      = errors.New("invalid input")
	ErrDuplicateRegister = errors.New("register number already exists")
)

func UpdateBook(bookID int, dto dtos.UpdateBookDto) (models.Book, error) {
	// Find the book by ID
	book := repository.FindBookById(bookID)
	if book.ID == 0 {
		return models.Book{}, ErrBookNotFound
	}

	// Validate register number if provided
	if dto.RegisterNumber != nil {
		registerNumber, err := strconv.Atoi(*dto.RegisterNumber)
		if err != nil {
			return models.Book{}, fmt.Errorf("%w: invalid register number", ErrInvalidInput)
		}
		// Check if register number exists (skip if it's the same book)
		if registerNumber != book.RegisterNumber && repository.ExistsByCode(registerNumber) {
			return models.Book{}, ErrDuplicateRegister
		}
	}

	// Update fields if they are provided in the DTO
	if err := buildBookUpdate(dto, &book); err != nil {
		return models.Book{}, fmt.Errorf("failed to update book fields: %w", err)
	}

	// Save the updated book within a transaction
	if err := repository.UpdateBookTx(book); err != nil {
		return models.Book{}, fmt.Errorf("failed to save book: %w", err)
	}

	return book, nil
}

func buildBookUpdate(dto dtos.UpdateBookDto, book *models.Book) error {
	if dto.RegisterNumber != nil {
		registerNumber, err := strconv.Atoi(*dto.RegisterNumber)
		if err != nil {
			return errors.New("invalid register number")
		}
		book.RegisterNumber = registerNumber
	}

	if dto.Authors != nil {
		book.Authors = getAuthors(*dto.Authors)
	}
	if dto.Title != nil {
		book.Title = *dto.Title
	}
	if dto.Volume != nil {
		book.Volume = *dto.Volume
	}
	if dto.Copy != nil {
		book.Copy = *dto.Copy
	}
	if dto.Location != nil {
		book.Location = *dto.Location
	}
	if dto.Publisher != nil {
		publisher := repository.FindOrCreatePublisher(*dto.Publisher)
		book.Publisher = publisher
	}
	if dto.PublicationYear != nil {
		book.PublicationYear = *dto.PublicationYear
	}
	if dto.AcquisitionForm != nil {
		book.AcquisitionForm = *dto.AcquisitionForm
	}
	if dto.Index != nil {
		book.Index = *dto.Index
	}
	if dto.Cdd != nil {
		book.CDD = *dto.Cdd
	}
	if dto.Observation != nil {
		book.Observation = *dto.Observation
	}
	return nil
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

func getAuthors(authorNames string) []models.Author {
	authorsNames := strings.Split(authorNames, " e ")
	var authors []models.Author
	for _, authorName := range authorsNames {
		author := repository.FindOrCreateAuthor(authorName)
		authors = append(authors, author)
	}
	return authors
}
