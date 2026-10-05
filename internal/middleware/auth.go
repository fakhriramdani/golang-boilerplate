package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"golang-boilerplate/pkg/response"
)

func JWTAuth(tokens *TokenManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer ") {
				return response.JSON(c, http.StatusUnauthorized, "missing or malformed token", nil)
			}
			userID, err := tokens.Validate(strings.TrimPrefix(header, "Bearer "), TypeAccess)
			if err != nil {
				return response.JSON(c, http.StatusUnauthorized, "invalid or expired token", nil)
			}
			c.Set("user_id", userID)
			return next(c)
		}
	}
}
