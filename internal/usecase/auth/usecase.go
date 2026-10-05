package auth

import (
	"context"

	"golang-boilerplate/internal/entity"
)

type AuthUsecase interface {
	Register(ctx context.Context, name, email, password string) (*entity.User, error)
	Login(ctx context.Context, email, password string) (*entity.User, string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, string, error)
}
