package controllers

import (
	"net/http"

	"app.com/backend/database"
	"github.com/gin-gonic/gin"
)

func HandleGetUsers(c *gin.Context) {
	var loadedTasks, err = database.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": loadedTasks})
}
