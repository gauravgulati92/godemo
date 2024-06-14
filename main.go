package main

import (
	"apidemo/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file")
	}
}

func main() {

	router := gin.Default()

	router.GET("/ping", controllers.GetPing)
	router.POST("/createFullName", controllers.PostFullname)
	router.Run()

}
