package api

import (
	"base-go-template/src/api/request"
	"base-go-template/src/application/interfaces"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(_echo *echo.Echo, service interfaces.UserService) *UserController {
	controller := &UserController{
		service: service,
	}

	_echo.GET("/users", controller.GetAllUsers)
	_echo.GET("/users/:id", controller.GetUserById)
	_echo.POST("/users", controller.CreateUser)

	return controller
}

func (uc *UserController) CreateUser(_echo echo.Context) error {
	var createUserRequest request.CreateUserRequest
	if err := _echo.Bind(&createUserRequest); err != nil {
		return _echo.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	userCommand, err := createUserRequest.ToCreateUserCommand()
	if err != nil {
		return _echo.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID format",
		})
	}

	result, err := uc.service.CreateUser(userCommand)
	if err != nil {
		return _echo.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return _echo.JSON(http.StatusCreated, result.Result)
}

func (uc *UserController) GetAllUsers(_echo echo.Context) error {
	users, err := uc.service.GetAllUser()
	if err == nil {
		return _echo.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch users",
		})
	}

	return _echo.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserById(_echo echo.Context) error {
	id, err := uuid.Parse(_echo.Param("id"))
	if err != nil {
		return _echo.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user id",
		})
	}

	user, err := uc.service.FindUserById(id)
	if err != nil {
		return _echo.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch user",
		})
	}

	return _echo.JSON(http.StatusOK, user)
}
