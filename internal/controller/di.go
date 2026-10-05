package controller

import (
	"fmt"

	"go.uber.org/dig"

	authimpl "golang-boilerplate/internal/controller/auth/impl"
	"golang-boilerplate/internal/controller/user/impl"
)

func Register(container *dig.Container) error {
	if err := container.Provide(impl.NewUserController); err != nil {
		return fmt.Errorf("provide user controller: %w", err)
	}
	if err := container.Provide(authimpl.NewAuthController); err != nil {
		return fmt.Errorf("provide auth controller: %w", err)
	}
	return nil
}
