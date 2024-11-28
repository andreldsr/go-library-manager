package repository

import (
	"go-library-manager/internal/database"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/models"
	"gorm.io/gorm"
	"time"
)

func CreateLending(lending *models.Lending) *models.Lending {
	database.DB.Create(lending)
	return lending
}

func ReturnLending(lending models.Lending) models.Lending {
	lending.ReturnedAt = time.Now()
	database.DB.Save(&lending)
	return lending
}

func FindLendingById(id int) (lending models.Lending) {
	lending.ID = id
	database.DB.Joins("Book").Find(&lending)
	return
}

func FindLendingDetailById(id int) (result dtos.LendingDetailDto) {
	database.DB.
		Model(models.Lending{ID: id}).
		Joins("Book").
		Joins("User").
		Select(`"lending".id, "Book".title book_title, "Book".id book_id, "Book".copy book_copy,
						"Book".location book_location, "Book".observation book_observation, "lending".returned_at,
						"lending".return_date, "User".name user_name, "lending".created_at`).
		Scan(&result)
	return
}

func FindAllLendingsActive(pageNumber, pageSize int) dtos.Page[dtos.LendingListDto] {
	return findLendingsPage(pageNumber, pageSize, returnAll())
}

func FindAllLendingsDueToday(pageNumber, pageSize int) dtos.Page[dtos.LendingListDto] {
	return findLendingsPage(pageNumber, pageSize, dueToday())
}

func FindAllLendingsOverdue(pageNumber, pageSize int) dtos.Page[dtos.LendingListDto] {
	return findLendingsPage(pageNumber, pageSize, overdue())
}

func FindAllLendingsBetweenDate(startDate, endDate string, pageNumber, pageSize int) dtos.Page[dtos.LendingListDto] {
	return findLendingsPage(pageNumber, pageSize, betweenDate(startDate, endDate))
}

func findLendingsPage(pageNumber, pageSize int, scope func(db *gorm.DB) *gorm.DB) dtos.Page[dtos.LendingListDto] {
	contentChan := make(chan []dtos.LendingListDto)
	countChan := make(chan int)
	go findLendingsContent(pageNumber, pageSize, scope, contentChan)
	go countLendings(returnAll(), countChan)
	return dtos.BuildPage(<-contentChan, <-countChan, pageNumber, pageSize)
}

func findLendingsContent(pageNumber, pageSize int, scope func(db *gorm.DB) *gorm.DB, contentChan chan []dtos.LendingListDto) (result []dtos.LendingListDto) {
	database.DB.
		Limit(pageSize).
		Offset(pageSize * pageNumber).
		Model(&models.Lending{}).
		Joins("Book").
		Joins("User").
		Where("returned_at is null").
		Scopes(scope).
		Select(`"lending".id id, 
						"Book".id book_id, 
						"Book".title book_title, 
						"User".name user_name, 
						"lending".return_date, 
						"lending".returned_at,
						"lending".created_at`).
		Order(`"User".name`).
		Scan(&result)
	contentChan <- result
	return
}

func countLendings(scope func(db *gorm.DB) *gorm.DB, countChan chan int) (count int64) {
	database.DB.
		Model(&models.Lending{}).
		Where("returned_at is null").
		Scopes(scope).
		Count(&count)
	countChan <- int(count)
	return
}

func returnAll() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func dueToday() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("lending.return_date = ?", time.Now().Format("2006-01-02"))
	}
}

func overdue() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("lending.return_date < ?", time.Now().Format("2006-01-02"))
	}
}

func betweenDate(startDate, endDate string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("lending.return_date between ? and ?", startDate, endDate)
	}
}
