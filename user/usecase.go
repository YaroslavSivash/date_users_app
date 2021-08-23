package user

import (
	"date_users_app/models"
	"github.com/labstack/echo"
)

type UseCase interface {
	CreateUser(c echo.Context, user *models.User) error
	GetAllUsers(c echo.Context, perPage, page int) ([]*models.User, error)
	UpdateUser(c echo.Context, user *models.User) error
}
