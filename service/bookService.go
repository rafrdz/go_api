package service

import (
	"github.com/rafrdz/go_api/model"
	"github.com/rafrdz/go_api/repository"
	viewmodel "github.com/rafrdz/go_api/viewModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type BookService interface {
	CreateNewBook(newBook *viewmodel.NewBook) (string, error)
	GetAllBooks() ([]viewmodel.Book, error)
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
	return service.bookRepo.SaveNewBook(&book)
}

func (service *bookService) GetAllBooks() ([]viewmodel.Book, error) {
	allDBBooks, err := service.bookRepo.GetAllBooks()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	var allBooks []viewmodel.Book
	for _, book := range allDBBooks {
		var newBook viewmodel.Book
		newBook.ID = book.ID.Hex()
		newBook.Title = book.Title
		newBook.Author = book.Author
		allBooks = append(allBooks, newBook)
	}
	return allBooks, nil
}