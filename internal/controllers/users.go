package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/services"
	"go-library-manager/internal/util"
	"net/http"
)

func FindUserList(c *gin.Context) {
	query := c.Query("name")
	fmt.Sprint("Find user with query" + query)
	pageNumber := util.IntOrDefault(c.Query("page"), 0)
	pageSize := util.IntOrDefault(c.Query("size"), 10)
	c.JSON(200, services.FindUsersByName(query, pageNumber, pageSize))
}

func Login(c *gin.Context) {
	var dto dtos.UserLoginDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := services.Login(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{
		"message": "token generated succesfully",
	})
}

func Register(c *gin.Context) {
	var dto dtos.CreateUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := services.CreateUser(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "new user created",
	})
}
