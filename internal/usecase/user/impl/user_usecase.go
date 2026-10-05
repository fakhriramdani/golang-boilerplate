package impl

import (
	"context"

	"golang-boilerplate/internal/entity"
)

func (u *userUsecaseImpl) GetAll(ctx context.Context) ([]entity.User, error) {
	return u.userRepository.FindAll(ctx)
}

func (u *userUsecaseImpl) GetByID(ctx context.Context, id uint) (entity.User, error) {
	return u.userRepository.FindByID(ctx, id)
}

func (u *userUsecaseImpl) Create(ctx context.Context, user *entity.User) error {
	return u.userRepository.Create(ctx, user)
}

func (u *userUsecaseImpl) Update(ctx context.Context, user *entity.User) error {
	return u.userRepository.Update(ctx, user)
}

func (u *userUsecaseImpl) Delete(ctx context.Context, id uint) error {
	return u.userRepository.Delete(ctx, id)
}
