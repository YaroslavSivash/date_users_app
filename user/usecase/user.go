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
	return u.repo.CreateUser(ctx, user)
}

func (u UserUseCase) GetUsers(ctx context.Context, skip, limit int) ([]*models.User, error) {
	u.repo.GetUsers()
}

func (u UserUseCase) UpdateUser(ctx context.Context, user *models.User, id string) error {
	panic("implement me")
}
