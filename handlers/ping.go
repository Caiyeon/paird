package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// for returning JSON bodies
type H map[string]interface{}

func Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"result": "Pong!",
		})
	}
}
