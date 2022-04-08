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

	// DB Connection Local
	//dsn := "host=localhost user=postgres password=paramadaksa dbname=pustaka port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// DB Connection Heroku
	dsn := "host=ec2-3-217-251-77.compute-1.amazonaws.com user=hbweazruihhypg password=2ba20672d00c6aaf213ddc22e5fe6c1a3f06d72920cf250872dbeca0ccbd9911 dbname=d8ut4jqc4ec0aa port=5432 sslmode=require TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("DB Connection Failed")
	}
	fmt.Println("DB Connection Success")

	// Migration
	db.AutoMigrate(&book.Book{})

	// Book
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// Routing
	router := gin.Default()

	// Versioning
	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/health", bookHandler.Health)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run()
}
