package impl

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"golang-boilerplate/internal/entity"
	"golang-boilerplate/pkg/response"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type authResponse struct {
	User entity.User `json:"user"`
	tokenResponse
}

func (ctrl *authControllerImpl) Register(c echo.Context) error {
	var req registerRequest
	if err := c.Bind(&req); err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	user, err := ctrl.authUsecase.Register(c.Request().Context(), req.Name, req.Email, req.Password)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.JSON(c, http.StatusCreated, "success", user)
}

func (ctrl *authControllerImpl) Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	user, access, refresh, err := ctrl.authUsecase.Login(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		return response.JSON(c, http.StatusUnauthorized, err.Error(), nil)
	}
	return response.JSON(c, http.StatusOK, "success", authResponse{User: *user, tokenResponse: tokenResponse{AccessToken: access, RefreshToken: refresh}})
}

func (ctrl *authControllerImpl) Refresh(c echo.Context) error {
	var req refreshRequest
	if err := c.Bind(&req); err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	access, refresh, err := ctrl.authUsecase.Refresh(c.Request().Context(), req.RefreshToken)
	if err != nil {
		return response.JSON(c, http.StatusUnauthorized, err.Error(), nil)
	}
	return response.JSON(c, http.StatusOK, "success", tokenResponse{AccessToken: access, RefreshToken: refresh})
}
