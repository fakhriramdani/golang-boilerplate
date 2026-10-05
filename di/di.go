package di

import (
	"fmt"
	"golang-boilerplate/pkg/config"

	"go.uber.org/dig"

	"golang-boilerplate/internal/controller"
	"golang-boilerplate/internal/middleware"
	"golang-boilerplate/internal/repository"
	"golang-boilerplate/internal/router"
	"golang-boilerplate/internal/usecase"
	"golang-boilerplate/resource"
)

func Register(container *dig.Container) error {
	if err := container.Provide(config.LoadConfig); err != nil {
		return fmt.Errorf("provide config: %w", err)
	}
	if err := container.Provide(resource.Connect); err != nil {
		return fmt.Errorf("provide database: %w", err)
	}
	if err := container.Provide(func(cfg *config.Config) *middleware.TokenManager {
		return middleware.NewTokenManager(cfg.JWTSecret)
	}); err != nil {
		return fmt.Errorf("provide token manager: %w", err)
	}
	if err := repository.Register(container); err != nil {
		return err
	}
	if err := usecase.Register(container); err != nil {
		return err
	}
	if err := controller.Register(container); err != nil {
		return err
	}
	if err := container.Provide(router.NewRouter); err != nil {
		return fmt.Errorf("provide router: %w", err)
	}
	return nil
}
