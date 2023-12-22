package services

import (
	"errors"
	"fmt"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/models"
	"go-library-manager/internal/repository"
	"time"
)

func FindAllLendingsActive(pageNumber, pageSize int) dtos.Page[dtos.LendingListDto] {
	return repository.FindAllLendingsActive(pageNumber, pageSize)
}

func FindAllLendingsDueToday(pageNumber, pageSize int) dtos.Page[dtos.LendingListDto] {
	return repository.FindAllLendingsDueToday(pageNumber, pageSize)
}

func FindAllLendingsOverdue(pageNumber, pageSize int) dtos.Page[dtos.LendingListDto] {
	return repository.FindAllLendingsOverdue(pageNumber, pageSize)
}

func FindLendingById(id int) dtos.LendingDetailDto {
	return repository.FindLendingDetailById(id)
}

func ReturnLending(id int) models.Lending {
	if id == 0 {
		panic(errors.New("invalid id"))
	}
	lending := repository.FindLendingById(id)
	if lending.ID == 0 {
		panic(errors.New("lending not found"))
	}
	repository.ReturnLending(lending)
	lending.Book.LendingId = 0
	repository.RemoveLending(lending.Book, lending)
	return lending
}

func CreateLending(dto dtos.CreateLendingDto) models.Lending {
	bookChan := make(chan models.Book)
	userChan := make(chan models.User)
	go repository.FindBookByIdAsync(dto.BookId, bookChan)
	go repository.FindUserByIdAsync(dto.UserId, userChan)
	returnDate, err := time.Parse("2006-01-02", dto.ReturnDate)
	if err != nil {
		fmt.Println(err.Error())
		return models.Lending{}
	}
	book := <-bookChan
	user := <-userChan
	lending := models.Lending{
		BookId:     book.ID,
		UserId:     user.ID,
		ReturnDate: returnDate,
	}
	repository.CreateLending(&lending)
	book.LendingId = lending.ID
	repository.UpdateBook(book)
	return lending
}
