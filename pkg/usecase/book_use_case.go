package usecase

import (
	"github.com/Melom01/go-boilerplate/pkg/domain"
	repository "github.com/Melom01/go-boilerplate/pkg/repository/interfaces"
	usecase "github.com/Melom01/go-boilerplate/pkg/usecase/interfaces"
)

type bookUseCase struct {
	bookRepo repository.BookRepositoryInterface
}

func CreateBookUseCase(repo repository.BookRepositoryInterface) usecase.BookUseCaseInterface {
	return &bookUseCase{
		bookRepo: repo,
	}
}

func (b bookUseCase) FindAll() ([]domain.Book, error) {
	books, err := b.bookRepo.FindAll()
	return books, err
}

func (b bookUseCase) FindByID(id uint) (domain.Book, error) {
	book, err := b.bookRepo.FindByID(id)
	return book, err
}

func (b bookUseCase) AddNewBook(book domain.Book) (domain.Book, error) {
	book, err := b.bookRepo.AddNewBook(book)
	return book, err
}

func (b bookUseCase) DeleteBookByID(id uint) error {
	err := b.bookRepo.DeleteBookByID(id)
	return err
}
