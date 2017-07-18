package handlers

import (
	"net/http"
	"strings"

	"github.com/caiyeon/lunch-with-us/store"

	"github.com/labstack/echo"
)

func SetDiet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// make sure team_domain, user_name, and text exist
		// enforce text to be none or vegeterian
		// set user's diet key with store.SetUserDiet(username, teamname, diet)
		diet := strings.ToLower(c.FormValue("text"))
		if diet != "" && diet != "vegetarian" && diet != "none" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "Only vegetarian options are available at the moment",
			})
		}

		err := store.SetUserKeyValue(c.FormValue("user_name"), c.FormValue("team_domain"), "diet", diet)
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": "Only vegetarian options are available at the moment",
			})
		}

		return c.JSON(http.StatusOK, H{
			"text": "Successfully set diet to " + diet + "!",
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
