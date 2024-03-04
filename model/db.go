package model

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
  "github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

func init() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/test?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "test:test@tcp(go-test-db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(dsn + " database can't connect")
	}
	DB.AutoMigrate(&User{})
}