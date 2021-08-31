package controllers

import "github.com/gin-gonic/gin"

func GetOk(c *gin.Context) {
	c.JSON(200, gin.H{
		"value": "OK",
	})
}
