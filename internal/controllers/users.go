package controllers

import (
	"github.com/gin-gonic/gin"
	"go-library-manager/internal/dtos"
	"go-library-manager/internal/services"
	"go-library-manager/internal/util"
	"net/http"
)

func FindUserList(c *gin.Context) {
	query := c.Query("name")
	pageNumber := util.IntOrDefault(c.Query("page"), 0)
	pageSize := util.IntOrDefault(c.Query("size"), 10)
	c.JSON(200, services.FindUsersByName(query, pageNumber, pageSize))
}

func FindUserById(c *gin.Context) {
	id := util.IntOrDefault(c.Param("id"), 0)
	c.JSON(200, services.FindUserById(id))
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
	if dto.Login == "" {
		dto.Login = dto.Document
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

func UpdateUser(context *gin.Context) {
	id := util.IntOrDefault(context.Param("id"), 0)
	var dto dtos.CreateUserDto
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := services.UpdateUser(id, dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "user updated",
	})
}
