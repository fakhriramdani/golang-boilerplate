package usecase

import (
	"fmt"

	"go.uber.org/dig"

	authimpl "golang-boilerplate/internal/usecase/auth/impl"
	"golang-boilerplate/internal/usecase/user/impl"
)

func Register(container *dig.Container) error {
	if err := container.Provide(impl.NewUserUsecase); err != nil {
		return fmt.Errorf("provide user usecase: %w", err)
	}
	if err := container.Provide(authimpl.NewAuthUsecase); err != nil {
		return fmt.Errorf("provide auth usecase: %w", err)
	}
	return nil
}
