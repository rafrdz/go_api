package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rafrdz/go_api/controller"
	"github.com/rafrdz/go_api/controllers"
)

var bookController controller.BookController = controller.NewBookController()

func main() {
	r := gin.Default()

	// Public routes, no authentication
	apiRoutes := r.Group("/api") {
		apiapiRoutes.POST("/books/new", func(c *gin.Context) {
			bookID, err := contrcontrollers.HandleCreateBook(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
			c.JSON(http.StatusOK, gin.H{"message": "Book Created!", "id": bookID})
		})
	}
	// TODO: Finish these other routes
	//r.GET("/books", controllers.FindBooks)
	//r.GET("/books/:id", controllers.FindBook)
	//r.PATCH("/books/:id", controllers.UpdateBook)
	//r.DELETE("/books/:id")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
