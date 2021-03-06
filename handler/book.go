package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/book"
	"strconv"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(context *gin.Context) {

	books, err := h.bookService.FindAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertToBookResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	context.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	// Convert to DTO Response
	bookResponse := convertToBookResponse(b)

	context.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
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

func (h *bookHandler) UpdateBookHandler(context *gin.Context) {
	var bookRequest book.UpdateBookRequest

	err := context.ShouldBindJSON(&bookRequest)
	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf(`Error on field %s, condition %s`, e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := context.Param("id")
	id, err := strconv.Atoi(idString)
	book, err := h.bookService.UpdateBook(id, bookRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (h *bookHandler) DeleteBook(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.Atoi(idString)

	b, err := h.bookService.DeleteBook(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	// Convert to DTO Response
	bookResponse := convertToBookResponse(b)

	context.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		Title: b.Title,
		Price: b.Price,
		Description: b.Description,
		Rating: b.Rating,
		Discount: b.Discount,
	}
}
