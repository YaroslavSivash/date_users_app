package user

import (
	"date_users_app/models"
	"github.com/labstack/echo"
)

type UserRepository interface {
	CreateUserDB(c echo.Context, user *models.User) error
	GetAllUsersDB(c echo.Context, skip, limit int) ([]*models.User, error)
	UpdateUserDB(c echo.Context, user *models.User) error
}
