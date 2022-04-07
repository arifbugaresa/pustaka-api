package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
)

func main() {

	// DB Connection
	dsn := "host=localhost user=postgres password=paramadaksa dbname=pustaka port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("DB Connection Failed")
	}
	fmt.Println("DB Connection Success")

	// Migration
	db.AutoMigrate(&book.Book{})

	// Book Repository
	bookRepository := book.NewRepository(db)
	books, err := bookRepository.FindAll()
	for _, book := range books {
		fmt.Println("Title :", book.Title)
	}

	// Routing
	router := gin.Default()

	// Versioning
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/book/:id", handler.BookHandler)
	v1.POST("/books", handler.PostBookHandler)

	router.Run()
}