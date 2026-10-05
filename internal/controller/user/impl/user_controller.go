package impl

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"golang-boilerplate/internal/entity"
	"golang-boilerplate/pkg/response"
)

func (ctrl *userControllerImpl) GetAll(c echo.Context) error {
	users, err := ctrl.userUsecase.GetAll(c.Request().Context())
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.JSON(c, http.StatusOK, "success", users)
}

func (ctrl *userControllerImpl) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, "invalid id", nil)
	}

	user, err := ctrl.userUsecase.GetByID(c.Request().Context(), uint(id))
	if err != nil {
		return response.JSON(c, http.StatusNotFound, err.Error(), nil)
	}
	return response.JSON(c, http.StatusOK, "success", user)
}

func (ctrl *userControllerImpl) Create(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := ctrl.userUsecase.Create(c.Request().Context(), &user); err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.JSON(c, http.StatusCreated, "success", user)
}

func (ctrl *userControllerImpl) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, "invalid id", nil)
	}

	var user entity.User
	if err := c.Bind(&user); err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	user.ID = uint(id)

	if err := ctrl.userUsecase.Update(c.Request().Context(), &user); err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.JSON(c, http.StatusOK, "success", user)
}

func (ctrl *userControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, "invalid id", nil)
	}

	if err := ctrl.userUsecase.Delete(c.Request().Context(), uint(id)); err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.JSON(c, http.StatusOK, "success", nil)
}
