package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"result": "Pong!",
		})
	}
}
