package impl

import (
	"context"

	"golang-boilerplate/internal/entity"
)

func (r *userRepositoryImpl) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uint) (entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	return user, err
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}
