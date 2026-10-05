package user

import (
	"context"

	"golang-boilerplate/internal/entity"
)

type UserUsecase interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id uint) (entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
}
