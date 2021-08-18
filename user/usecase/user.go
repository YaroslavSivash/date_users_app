package usecase

import (
	"context"
	"date_users_app/models"
	"date_users_app/user"
)

type UserUseCase struct {
	repo user.UserRepository
}

func NewUserUseCase(repo user.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u UserUseCase) CreateUser(ctx context.Context, user *models.User) error {
	return u.repo.CreateUserDB(ctx, user)
}

func (u UserUseCase) GetAllUsers(ctx context.Context, skip, limit int) ([]*models.User, error) {
	return u.repo.GetAllUsersDB(ctx, skip, limit)
}

func (u UserUseCase) UpdateUser(ctx context.Context, user *models.User, id string) error {
	return u.repo.UpdateUserDB(ctx, user, id)
}
