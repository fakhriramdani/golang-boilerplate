package impl

import (
	controller "golang-boilerplate/internal/controller/user"
	usecase "golang-boilerplate/internal/usecase/user"
)

type userControllerImpl struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) controller.UserController {
	return &userControllerImpl{userUsecase: userUsecase}
}
