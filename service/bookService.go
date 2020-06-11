package service

import (
	"github.com/rafrdz/go_api/db"
	"github.com/rafrdz/go_api/model"
	"github.com/rafrdz/go_api/repository"
	viewmodel "github.com/rafrdz/go_api/viewModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService interface {
	CreateNewBook(newBook *viewmodel.NewBook) (string, error)
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService() BookService {
	return &bookService{
		bookRepo: repository.NewBookRepository("books", "books"),
	}
}

func (service *bookService) CreateNewBook(newBook *viewmodel.NewBook) (string, error) {
	var book model.Book
	book.ID = primitive.NewObjectID()
	book.Title = newBook.Title
	book.Author = newBook.Author
	return bookRepo.SaveNewBook(&book)
}