package user

import (
	"context"
	"date_users_app/models"
)

type UserRepository interface {
	CreateUserDB(ctx context.Context, user *models.User) error
	GetAllUsersDB(ctx context.Context, skip, limit int) ([]*models.User, error)
	UpdateUserDB(ctx context.Context, user *models.User, id string) error
}
