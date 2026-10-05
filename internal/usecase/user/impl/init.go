package impl

import (
	repository "golang-boilerplate/internal/repository/user"
	usecase "golang-boilerplate/internal/usecase/user"
)

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}
