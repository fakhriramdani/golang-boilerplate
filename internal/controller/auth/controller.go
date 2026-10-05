package auth

import "github.com/labstack/echo/v4"

type AuthController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	Refresh(c echo.Context) error
}
