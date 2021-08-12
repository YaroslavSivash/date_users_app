package user

import (
	"context"
	"date_users_app/models"
)

type UseCase interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context, skip, limit int) ([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User, id string) error
}
