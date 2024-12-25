package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/Melom01/go-boilerplate/pkg/domain"
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
		return
	} else {
		var response []domain.Book

		if copierErr := copier.Copy(&response, &books); copierErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h *BookHandler) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, parseErr := strconv.Atoi(idParam)

	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse book ID"})
		return
	} else {
		var response domain.Book

		book, findBookErr := h.bookUseCase.FindByID(uint(id))
		if findBookErr != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if copierErr := copier.Copy(&response, &book); copierErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
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
		return
	}

	response := domain.Book{}
	if copierErr := copier.Copy(&response, &book); copierErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *BookHandler) DeleteBookById(c *gin.Context) {
	idParam := c.Param("id")
	id, parseErr := strconv.Atoi(idParam)

	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse book ID"})
		return
	}

	if deleteBookErr := h.bookUseCase.DeleteBookByID(uint(id)); deleteBookErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}
