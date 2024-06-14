package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func PostFullname(c *gin.Context) {

	var requestBody struct {
		FirstName string
		LastName  string
	}
	c.Bind(&requestBody)

	c.JSON(200, gin.H{
		"fullname": requestBody.FirstName + " " + requestBody.LastName,
		"author":   os.Getenv("AUTHOR"),
	})
}
