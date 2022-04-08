package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pustaka-api/book"
	"pustaka-api/configuration"
	"pustaka-api/handler"
)

var (
	db             *gorm.DB = configuration.SetupDatabaseConnection()
	bookRepository          = book.NewRepository(db)
	bookService             = book.NewService(bookRepository)
	bookHandler             = handler.NewBookHandler(bookService)
)

func main() {
	defer configuration.CloseDatabaseConnection(db)

	// Migration
	db.AutoMigrate(&book.Book{})

	// Routing
	router := gin.Default()

	// Versioning
	version := "/v1"

	// Grouping
	rootRoutes := router.Group(version+"/api")
	{
		rootRoutes.GET("/health", bookHandler.Health)
	}

	bookRoutes := router.Group(version+"/api")
	{
		bookRoutes.GET("/books/:id", bookHandler.GetBook)
		bookRoutes.GET("/books", bookHandler.GetBooks)
		bookRoutes.POST("/books", bookHandler.PostBookHandler)
		bookRoutes.PUT("/books/:id", bookHandler.UpdateBookHandler)
		bookRoutes.DELETE("/books/:id", bookHandler.DeleteBook)
	}

	router.Run()
}
