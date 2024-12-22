package handler

import (
	"net/http"

	"github.com/Melom01/go-boilerplate/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	usecase "github.com/Melom01/go-boilerplate/pkg/usecase/interfaces"
)

type BookHandler struct {
	bookUseCase usecase.BookUseCaseInterface
}

func CreateBookHandler(usecase usecase.BookUseCaseInterface) *BookHandler {
	return &BookHandler{
		bookUseCase: usecase,
	}
}

func (h *BookHandler) FindAll(c *gin.Context) {
	books, err := h.bookUseCase.FindAll()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		var response []domain.Book
		copierErr := copier.Copy(&response, &books)

		if copierErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h *BookHandler) AddNewBook(c *gin.Context) {
	var book domain.Book

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	book, err := h.bookUseCase.AddNewBook(book)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		response := domain.Book{}
		copierErr := copier.Copy(&response, &book)
		if copierErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, response)
	}
}
