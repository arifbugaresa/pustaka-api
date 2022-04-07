package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/book"
)

func RootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "Agung",
	})
}

func BookHandler(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func PostBookHandler(context *gin.Context) {
	var bookInput book.BookInput

	err := context.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf(`Error on field %s, condition %s`, e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"errors" : errorMessages,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}