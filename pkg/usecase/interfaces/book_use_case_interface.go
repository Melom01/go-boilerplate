package interfaces

import "github.com/Melom01/go-boilerplate/pkg/domain"

type BookUseCaseInterface interface {
	FindAll() ([]domain.Book, error)
	FindByID(id uint) (domain.Book, error)
	AddNewBook(book domain.Book) (domain.Book, error)
	DeleteBookByID(id uint) error
}
