package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rafrdz/go_api/controller"
)

var bookController controller.BookController = controller.NewBookController()

func main() {
	r := gin.Default()

	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/book/new", func(c *gin.Context) {
			bookID, err := bookController.HandleCreateBook(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Book Created!", "id": bookID})
			}
		})
		apiRoutes.GET("/book/all", func(c *gin.Context) {
			allBooks, err := bookController.HandleGetAllBooks(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			} else {
				c.JSON(http.StatusOK, allBooks)
			}
		})
		apiRoutes.GET("/book/id/:id", func(c *gin.Context) {
			book, err := bookController.HandleGetBookByID(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"book": book})
			}
		})
		apiRoutes.POST("/book/delete/:id", func(c * gin.Context) {
			deleteStatus, err := bookController.HandleDeleteBookByID(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": deleteStatus})
			}
		})
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
