package user

import (
	"context"
	"date_users_app/models"
)

type UserRepository interface {
	CreateUser (ctx context.Context, user *models.User, Email string) error
	GetUsers (ctx context.Context, user *models.User) error
	UpdateUser (ctx context.Context, user *models.User) (Id string, error)
}
