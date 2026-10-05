package response

import "github.com/labstack/echo/v4"

type Body struct {
	Code    int  `json:"code"`
	Message string `json:"message"`
	Data    any  `json:"data,omitempty"`
}

func JSON(c echo.Context, status int, message string, data any) error {
	return c.JSON(status, Body{Code: status, Message: message, Data: data})
}
