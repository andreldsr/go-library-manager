package controllers

import (
	"github.com/gin-gonic/gin"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/services"
	"go-library-manager/internal/util"
)

func FindAllLendingsActive(c *gin.Context) {
	pageNumber := util.IntOrDefault(c.Query("page"), 0)
	pageSize := util.IntOrDefault(c.Query("size"), 10)
	c.JSON(200, services.FindAllLendingsActive(pageNumber, pageSize))
}
func FindAllLendingsDueToday(c *gin.Context) {
	pageNumber := util.IntOrDefault(c.Query("page"), 0)
	pageSize := util.IntOrDefault(c.Query("size"), 10)
	c.JSON(200, services.FindAllLendingsDueToday(pageNumber, pageSize))
}
func FindAllLendingsOverdue(c *gin.Context) {
	pageNumber := util.IntOrDefault(c.Query("page"), 0)
	pageSize := util.IntOrDefault(c.Query("size"), 10)
	c.JSON(200, services.FindAllLendingsOverdue(pageNumber, pageSize))
}

func FindLendingById(c *gin.Context) {
	id := util.IntOrDefault(c.Param("id"), 0)
	if id == 0 {
		c.JSON(400, gin.H{
			"error": "Invalid id",
		})
		return
	}
	lending := services.FindLendingById(id)
	if lending.Id == 0 {
		c.JSON(404, gin.H{
			"error": "Lending not found",
		})
		return
	}
	c.JSON(200, lending)
}

func CreateLending(c *gin.Context) {
	dto := dtos.CreateLendingDto{}
	if err := c.ShouldBindJSON(&dto); err != nil {
		print(err.Error())
		return
	}
	c.JSON(200, services.CreateLending(dto))
}

func ReturnLending(c *gin.Context) {
	id := util.IntOrDefault(c.Param("id"), 0)
	if id == 0 {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}
	services.ReturnLending(id)
}
