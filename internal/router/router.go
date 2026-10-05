package router

import (
	"github.com/labstack/echo/v4"

	authController "golang-boilerplate/internal/controller/auth"
	controller "golang-boilerplate/internal/controller/user"
	"golang-boilerplate/internal/middleware"
)

func NewRouter(userController controller.UserController, authController authController.AuthController, tokens *middleware.TokenManager) *echo.Echo {
	e := echo.New()

	api := e.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
	auth.POST("/refresh", authController.Refresh)

	users := api.Group("/users", middleware.JWTAuth(tokens))

	users.GET("", userController.GetAll)
	users.GET("/:id", userController.GetByID)
	users.POST("", userController.Create)
	users.PUT("/:id", userController.Update)
	users.DELETE("/:id", userController.Delete)

	return e
}
