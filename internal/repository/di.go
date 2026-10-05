package repository

import (
	"fmt"

	"go.uber.org/dig"

	"golang-boilerplate/internal/repository/user/impl"
)

func Register(container *dig.Container) error {
	if err := container.Provide(impl.NewUserRepository); err != nil {
		return fmt.Errorf("provide user repository: %w", err)
	}
	return nil
}
