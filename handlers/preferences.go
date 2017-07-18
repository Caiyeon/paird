package handlers

import (
	// "net/http"
	// "strconv"
	// "strings"

	// "github.com/caiyeon/lunch-with-us/store"
	"net/http"

	"github.com/labstack/echo"
)

func SetDiet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// make sure team_domain, user_name, and text exist
		// enforce text to be none or vegeterian
		// set user's diet key with store.SetUserDiet(username, teamname, diet)

		return c.JSON(http.StatusOK, H{
			"text": "Some text to respond back to user",
		})
	}
}

func SetAvailability() echo.HandlerFunc {
	return func(c echo.Context) error {
		// make sure team_domain, user_name, and text exist
		// enforce text to be none or vegeterian
		// set user's diet key with store.SetUserDiet(username, teamname, diet)

		return c.JSON(http.StatusOK, H{
			"text": "Some text to respond back to user",
		})
	}
}

func SetLocation() echo.HandlerFunc {
	return func(c echo.Context) error {
		// make sure team_domain, user_name, and text exist
		// enforce text to be none or vegeterian
		// set user's diet key with store.SetUserDiet(username, teamname, diet)

		return c.JSON(http.StatusOK, H{
			"text": "Some text to respond back to user",
		})
	}
}
