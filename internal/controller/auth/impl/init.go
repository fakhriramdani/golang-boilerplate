package impl

import (
	auth "golang-boilerplate/internal/controller/auth"
	usecase "golang-boilerplate/internal/usecase/auth"
)

type authControllerImpl struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) auth.AuthController {
	return &authControllerImpl{authUsecase: authUsecase}
}
