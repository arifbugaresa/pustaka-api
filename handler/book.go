package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "Agung",
	})
}

func (h *bookHandler) BookHandler(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *bookHandler) PostBookHandler(context *gin.Context) {
	var bookRequest book.BookRequest

	err := context.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf(`Error on field %s, condition %s`, e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	// Service
	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
