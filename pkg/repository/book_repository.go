package repository

import (
	"gorm.io/gorm"

	"github.com/Melom01/go-boilerplate/pkg/domain"
	"github.com/Melom01/go-boilerplate/pkg/repository/interfaces"
)

type bookDatabase struct {
	db *gorm.DB
}

func CreateBookRepository(db *gorm.DB) interfaces.BookRepositoryInterface {
	return &bookDatabase{db: db}
}

func (c *bookDatabase) FindAll() ([]domain.Book, error) {
	var books []domain.Book
	err := c.db.Find(&books).Error

	return books, err
}

func (c *bookDatabase) FindByID(id uint) (domain.Book, error) {
	var book domain.Book
	err := c.db.Find(&book, id).Error

	return book, err
}

func (c *bookDatabase) AddNewBook(book domain.Book) (domain.Book, error) {
	err := c.db.Save(&book).Error

	return book, err
}

func (c *bookDatabase) DeleteBookByID(id uint) error {
	var book domain.Book
	err := c.db.Delete(&book, id).Error

	return err
}
