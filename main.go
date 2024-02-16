package main

import (
	"base-go-template/config"
	api "base-go-template/src/api/controllers"
	"base-go-template/src/application/services"
	"base-go-template/src/infrastructure"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	gormDB := config.NewConnection()
	userRepo := infrastructure.NewUserRepository(gormDB)

	userService := services.NewUserService(userRepo)

	_echo := echo.New()
	api.NewUserController(_echo, userService)

	// start server
	if err := _echo.Start("localhost:3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
