package service

import viewmodel "github.com/rafrdz/go_api/viewModel"

type BookService interface {
	CreateNewBook(newBook *viewmodel.NewBook) (string, error)
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService() BookService {
	return &bookService{
		bookRepo: repository.NewBookRepository(),
	}
}

func (service *bookService) CreateNewBook(newBook *viewmodel.NewBook) (string, error) {

}