package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load("config.env")

	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error loading config.env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	SSLMode := os.Getenv("SSL_MODE")
	TimeZone := os.Getenv("TIMEZONE")

	DBURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		DbHost, DbUser, DbPassword, DbName, DbPort, SSLMode, TimeZone,
	)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	DB.AutoMigrate(&User{})
}
