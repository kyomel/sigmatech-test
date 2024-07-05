package controller

import (
	"encoding/json"
	"net/http"
	"sigmatech-test/pkg/model"
	"sigmatech-test/pkg/usecase"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userUC usecase.UserUseCase
}

type UserController interface {
	RegisterUser(c echo.Context) error
	Login(c echo.Context) error
}

func NewUserController(userUC usecase.UserUseCase) UserController {
	return &userController{
		userUC,
	}
}

func (u *userController) RegisterUser(c echo.Context) error {
	var request model.RegisterUser
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = u.userUC.RegisterUser(c.Request().Context(), &request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (u *userController) Login(c echo.Context) error {
	var request model.LoginParam
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	result, err := u.userUC.Login(c.Request().Context(), &request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"token":   result,
	})

}
