package repository

import (
	"database/sql"
	"fmt"
	"go-library-manager/internal/database"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/models"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func FindAllBooks(query string, pageNumber, pageSize int) dtos.Page[dtos.BookListDto] {
	booksChan := make(chan []dtos.BookListDto)
	countChan := make(chan int)
	go findAllBooksContent(query, pageNumber, pageSize, booksChan)
	go findAllBooksCount(query, countChan)
	return dtos.BuildPage(<-booksChan, <-countChan, pageNumber, pageSize)
}

func UpdateBookTx(book models.Book) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Update book
		authors := book.Authors
		if err := tx.Model(&book).Omit("Authors").Updates(book).Error; err != nil {
			return err
		}

		if authors != nil {
			if err := tx.Model(&book).Association("Authors").Clear(); err != nil {
				return err
			}
			if err := tx.Model(&book).Association("Authors").Replace(authors); err != nil {
				return err
			}
		}

		return nil
	})
}

func ExistsByCode(code int) bool {
	var count int64
	database.DB.Model(&models.Book{}).Where("register_number = ?", code).Count(&count)
	return count > 0
}

func RemoveLending(book models.Book, lending models.Lending) models.Book {
	err := database.DB.
		Model(&lending).
		Association("Book").
		Delete(&book)
	if err != nil {
		panic(err.Error())
		return models.Book{}
	}
	return book
}

func FindBookById(id int) (book models.Book) {
	database.DB.Preload("Authors").Joins("Publisher").Find(&book, id)
	return
}

func FindBookByIdAsync(id int, c chan models.Book) {
	book := models.Book{ID: id}
	database.DB.Find(&book)
	c <- book
}

func GetStats() (stats dtos.BookStats) {
	database.DB.
		Raw(`SELECT 
            count(*) as total, 
            SUM(CASE WHEN l.id IS NOT NULL then 1 else 0 END) AS lent,
            SUM(CASE WHEN l.id IS NOT NULL AND l.return_date = ? THEN 1 ELSE 0 END) AS today,
            SUM(CASE WHEN l.id IS NOT NULL AND l.return_date < ? THEN 1 ELSE 0 END) AS delayed
       FROM book b LEFT JOIN lending l ON b.lending_id = l.id`,
			time.Now().Format("2006-01-02"),
			time.Now().Format("2006-01-02")).
		Scan(&stats)
	return
}

func findAllBooksCount(query string, countChan chan int) (count int) {
	var term = "%" + query + "%"
	var registerNumber, err = strconv.Atoi(query)
	if err != nil {
		fmt.Println(err)
		registerNumber = 0
	}
	database.DB.
		Raw(`SELECT COUNT(DISTINCT b.id)
				FROM book b 
				LEFT JOIN publisher p on b.publisher_id = p.id
				LEFT JOIN books_authors ba on b.id = ba.book_id
				LEFT JOIN author a on a.id = ba.author_id
				WHERE b.title ILIKE @term
				OR p.name ILIKE @term
				OR a.name ILIKE  @term
				OR b.register_number = @registerNumber`,
			sql.Named("term", term),
			sql.Named("registerNumber", registerNumber)).
		Scan(&count)
	countChan <- count
	return
}

func findAllBooksContent(query string, pageNumber, pageSize int, booksChan chan []dtos.BookListDto) []dtos.BookListDto {
	var results []dtos.BookListDto
	var registerNumber, err = strconv.Atoi(query)
	if err != nil {
		fmt.Println(err)
		registerNumber = 0
	}
	database.DB.
		Raw(`SELECT 
    				DISTINCT b.id AS id, 
    			    b.title AS title, 
    			    b.register_number,
    			    STRING_AGG(a.name, ', ') as authors_names, 
    			    p.name AS publisher_name, 
    			    b.cdd AS cdd,
    			    b.index AS index,
    			    l.id AS lending_id, 
    			    l.return_date AS lending_return_date 
				FROM book b 
				    LEFT JOIN publisher p ON b.publisher_id = p.id 
				    LEFT JOIN lending l on b.lending_id = l.id
				    LEFT JOIN books_authors ba ON b.id = ba.book_id
					LEFT JOIN author a ON ba.author_id = a.id 
				WHERE b.title ILIKE @term
				      OR a.name ILIKE @term
					  OR p.name ILIKE @term
					  OR b.register_number = @registerNumber
				GROUP BY b.id, p.name, l.id
				ORDER BY b.title
				LIMIT @pageSize OFFSET @offset`,
			sql.Named("term", "%"+query+"%"),
			sql.Named("registerNumber", registerNumber),
			sql.Named("pageSize", pageSize),
			sql.Named("offset", pageNumber*pageSize)).
		Scan(&results)
	booksChan <- results
	return results
}

func CreateBook(book *models.Book) {
	database.DB.Create(book)
}
