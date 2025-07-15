package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gnanasuryateja-kurapati/simple-rest-api-golang/model"
	"github.com/gnanasuryateja-kurapati/simple-rest-api-golang/src"
)

func GetUsers(c *gin.Context) {
	fmt.Println(src.GetUsers())
	c.JSON(http.StatusOK, src.GetUsers())
}

func GetUserById(c *gin.Context) {
	userId := c.Param("id")
	userData, err := src.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *userData)
}

func CreateUser(c *gin.Context) {
	var userData model.UserDataInput
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	message, err := src.CreateNewUser(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusAccepted, gin.H{"message": *message})
}

func UpdateUser(c *gin.Context) {
	var userData model.UserDataInput
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	message, err := src.UpdateUser(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusAccepted, gin.H{"message": *message})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	message, err := src.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusAccepted, gin.H{"message": *message})
}
