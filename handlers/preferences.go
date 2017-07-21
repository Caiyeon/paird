package handlers

import (
	"net/http"
	"strings"

	"github.com/caiyeon/paird/store"
	"github.com/labstack/echo"
)

func SetDiet() echo.HandlerFunc {
	return func(c echo.Context) error {
		diet := strings.ToLower(c.FormValue("text"))
		if diet != "" && diet != "vegetarian" && diet != "none" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "Only vegetarian options are available at the moment",
			})
		}

		err := store.SetUserKeyValue(c.FormValue("user_name"), c.FormValue("team_domain"), "diet", diet)
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, H{
			"text": "Successfully set diet to " + diet + "!",
		})
	}
}

func SetAvailability() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := store.SetUserKeyValue(c.FormValue("user_name"), c.FormValue("team_domain"), "avail", c.FormValue("text"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, H{
			"text": "Availability set!",
		})
	}
}

func SetLocation() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := store.SetUserKeyValue(c.FormValue("user_name"), c.FormValue("team_domain"), "location", c.FormValue("text"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, H{
			"text": "Location set!",
		})
	}
}
