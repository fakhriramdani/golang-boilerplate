package user

import (
	"context"

	"golang-boilerplate/internal/entity"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id uint) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
}
