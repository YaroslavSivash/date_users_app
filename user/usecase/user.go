package usecase

import (
	"date_users_app/models"
	"date_users_app/user"
	"github.com/labstack/echo"
)

type UserUseCase struct {
	repo user.UserRepository
}

func NewUserUseCase(repo user.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u UserUseCase) CreateUser(c echo.Context, user *models.User) error {
	return u.repo.CreateUserDB(c, user)
}

func (u UserUseCase) GetAllUsers(c echo.Context, skip, limit int) ([]*models.User, error) {
	return u.repo.GetAllUsersDB(c, skip, limit)
}

func (u UserUseCase) UpdateUser(c echo.Context, user *models.User) error {
	return u.repo.UpdateUserDB(c, user)
}
