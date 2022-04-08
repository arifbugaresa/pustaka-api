package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"pustaka-api/book"
	"pustaka-api/handler"
)

func main() {

	// DB Connection Local
	//dsn := "host=localhost user=postgres password=paramadaksa dbname=pustaka port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// DB Connection Heroku
	//dsn := fmt.Sprintf(`host=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta`,
	//	"c2-3-217-251-77.compute-1.amazonaws.com",
	//	"2ba20672d00c6aaf213ddc22e5fe6c1a3f06d72920cf250872dbeca0ccbd9911",
	//	"d8ut4jqc4ec0aa",
	//	"5432")

	// Load Environtment
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error env")
	} else {
		println("Succes read env file")
	}

	// Setting Env Database
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Local Config
	//dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbPort, dbName)

	//Sandbox Config
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)

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
