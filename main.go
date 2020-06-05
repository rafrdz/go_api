package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafrdz/go_api/controllers"
)

func main() {
	r := gin.Default()

	//r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.HandleCreateBook)
	//r.GET("/books/:id", controllers.FindBook)
	//r.PATCH("/books/:id", controllers.UpdateBook)
	//r.DELETE("/books/:id")

	r.Run()
}
