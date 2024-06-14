package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {

	var responseBody struct {
		StatusCode uint
		Message    string
		TimeStamp  string
	}

	responseBody.Message = "sever is running"
	responseBody.StatusCode = 200
	responseBody.TimeStamp = time.Now().GoString()

	c.JSON(int(responseBody.StatusCode), gin.H{
		"data": responseBody,
	})

}
