package config

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	envFile, _ := godotenv.Read(".env")
	connectionString := envFile["DATABASE_URL"]

	gormDB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v.", err)
	}

	return gormDB
}
