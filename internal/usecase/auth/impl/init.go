package impl

import (
	"golang-boilerplate/internal/middleware"
	repository "golang-boilerplate/internal/repository/user"
	auth "golang-boilerplate/internal/usecase/auth"
)

type authUsecaseImpl struct {
	userRepo repository.UserRepository
	tokens   *middleware.TokenManager
}

func NewAuthUsecase(userRepo repository.UserRepository, tokens *middleware.TokenManager) auth.AuthUsecase {
	return &authUsecaseImpl{userRepo: userRepo, tokens: tokens}
}
