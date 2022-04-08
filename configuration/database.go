package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {

	// Load Environment
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
	//dsn := fmt.Sprintf(`host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta`, dbHost, dbUser, dbPass, dbPort, dbName)

	// Sandbox Config
	dsn := fmt.Sprintf(`host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Jakarta`, dbHost, dbUser, dbPass, dbPort, dbName)

	// Open DB Connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("DB Connection Failed")
	}
	fmt.Println("DB Connection Success")

	return db
}

func CloseDatabaseConnection(db *gorm.DB)  {
	database, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	database.Close()
}