package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafrdz/go_api/controllers"
	"github.com/rafrdz/go_api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
