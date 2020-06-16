package controller

import (
	"github.com/rafrdz/go_api/service"
	viewmodel "github.com/rafrdz/go_api/viewModel"
	"log"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	HandleCreateBook(c *gin.Context) (string, error)
	HandleGetAllBooks(c *gin.Context) ([]viewmodel.Book, error)
	HandleGetBookByID(c *gin.Context) (*viewmodel.Book, error)
	HandleDeleteBookByID(c *gin.Context) (string, error)
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

func (controller *bookController) HandleGetAllBooks(c *gin.Context) ([]viewmodel.Book, error) {
	allBooks, err := controller.bookService.GetAllBooks()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return allBooks, nil
}

func (controller *bookController) HandleGetBookByID(c *gin.Context) (*viewmodel.Book, error) {
	id := c.Param("id")
	book, err := controller.bookService.GetBookByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (controller *bookController) HandleDeleteBookByID(c *gin.Context) (string, error) {
	id := c.Param("id")
	deleteResult, err := controller.bookService.DeleteBookByID(id)
	if err != nil {
		return "", err
	}
	return deleteResult, nil
}
