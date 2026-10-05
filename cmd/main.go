package main

import (
	"golang-boilerplate/pkg/config"
	"log"

	"go.uber.org/dig"

	"golang-boilerplate/di"

	"github.com/labstack/echo/v4"
)

func main() {
	container := dig.New()
	if err := di.Register(container); err != nil {
		log.Fatalf("failed to register DI: %v", err)
	}

	if err := container.Invoke(func(e *echo.Echo, cfg *config.Config) {
		e.Logger.Fatal(e.Start(":" + cfg.AppPort))
	}); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
