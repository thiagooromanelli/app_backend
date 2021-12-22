package controllers

import (
	"log"
	"net/http"

	"app.com/backend/database"
	"app.com/backend/models"
	"github.com/gin-gonic/gin"
)

func HandleGetUsers(c *gin.Context) {
	var loadedTasks, err = database.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": loadedTasks})
}

func HandlePostUsers(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var oid, err = database.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"oid": oid})
}

func HandleGetUserById(c *gin.Context) {
	var user models.User
	if err := c.BindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var loadedUser, err = database.GetUserByID(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": loadedUser})
}
