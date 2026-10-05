package impl

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"golang-boilerplate/internal/entity"
	"golang-boilerplate/internal/middleware"
)

func (u *authUsecaseImpl) Register(ctx context.Context, name, email, password string) (*entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &entity.User{Name: name, Email: email, Password: string(hash)}
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *authUsecaseImpl) Login(ctx context.Context, email, password string) (*entity.User, string, string, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", "", errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", "", errors.New("invalid email or password")
	}
	access, err := u.tokens.AccessToken(user.ID)
	if err != nil {
		return nil, "", "", err
	}
	refresh, err := u.tokens.RefreshToken(user.ID)
	if err != nil {
		return nil, "", "", err
	}
	return &user, access, refresh, nil
}

func (u *authUsecaseImpl) Refresh(_ context.Context, refreshToken string) (string, string, error) {
	userID, err := u.tokens.Validate(refreshToken, middleware.TypeRefresh)
	if err != nil {
		return "", "", err
	}
	access, err := u.tokens.AccessToken(userID)
	if err != nil {
		return "", "", err
	}
	refresh, err := u.tokens.RefreshToken(userID)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}
