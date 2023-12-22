package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/services"
	"go-library-manager/internal/util"
)

func FindAllBooks(c *gin.Context) {
	query := c.Query("query")
	page := util.IntOrDefault(c.Query("page"), 0)
	size := util.IntOrDefault(c.Query("size"), 10)
	c.JSON(200, services.FindAllBooks(query, page, size))
}

func FindBookById(c *gin.Context) {
	id := util.IntOrDefault(c.Param("id"), 0)
	if id == 0 {
		c.JSON(400, gin.H{
			"error": "Invalid id",
		})
	}

	book := services.FindBookById(id)

	if book.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Book not found",
		})
		return
	}
	c.JSON(200, book)
}

func GetBookStats(c *gin.Context) {
	c.JSON(200, services.GetBookStats())
}

func CreateBook(c *gin.Context) {
	var dtos []dtos.CreateBookDto
	err := c.ShouldBindJSON(&dtos)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	for _, dto := range dtos {
		book := services.CreateBook(dto)
		if book.ID == 0 {
			panic(errors.New("error creating new book"))
		}
	}
	c.JSON(201, gin.H{
		"Message": "Books created succesfully",
	})
}
