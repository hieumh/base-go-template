package main

import (
	api "base-go-template/src/api/controllers"
	"base-go-template/src/application/services"
	"base-go-template/src/infrastructure"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	envFile, _ := godotenv.Read(".env")
	connectionString := envFile["DATABASE_URL"]

	gormDB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v.", err)
	}

	userRepo := infrastructure.NewUserRepository(gormDB)

	userService := services.NewUserService(userRepo)

	_echo := echo.New()
	api.NewUserController(_echo, userService)

	if err := _echo.Start("localhost:3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
