package controller

import (
	"github.com/rafrdz/go_api/service"
	viewmodel "github.com/rafrdz/go_api/viewModel"
	"log"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	HandleCreateBook(c *gin.Context) (string, error)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController() BookController {
	return &bookController{
		bookService: service.NewBookService(),
	}
}

func (controller *bookController) HandleCreateBook(c *gin.Context) (string, error) {
	var book viewmodel.NewBook
	if err := c.ShouldBindJSON(&book); err != nil {
		log.Print(err)
		return "", err
	}

	id, err := controller.bookService.CreateNewBook(&book)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return id, nil
}

// func FindBooks(c *gin.Context) {
// 	var books []models.Book
// 	models.DB.Find(&books)

// 	c.JSON(http.StatusOK, gin.H{"data": books})
// }

// func FindBook(c *gin.Context) {
// 	var book models.Book

// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }

// func UpdateBook(c *gin.Context) {
// 	// Get model if exist
// 	var book models.Book
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input UpdateBookInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.DB.Model(&book).Updates(input)

// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }

// func DeleteBook(c *gin.Context) {
// 	// Get model if exist
// 	var book models.Book
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	models.DB.Delete(&book)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }
