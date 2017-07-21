package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/caiyeon/paird/store"
	"github.com/labstack/echo"
)

func ListTags() echo.HandlerFunc {
	return func(c echo.Context) error {
		// fetch team's tags:
		team := c.FormValue("team_domain")
		if team == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "team domain name is required",
			})
		}
		tags, err := store.ListTeamTags(team)
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err.Error(),
			})
		}

		text := "Popular tags in your team:\n"
		for key, value := range tags {
			text = text + "*" + key + "* - " + strconv.Itoa(value) + " tags found\n"
		}

		return c.JSON(http.StatusOK, H{
			"text": text,
		})
	}
}

func AddMyTags() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tags []string
		if raw := c.FormValue("text"); raw == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "text parameters are required",
			})
		} else {
			raw = strings.Trim(raw, "[]")
			tags = strings.Split(raw, ",")
		}

		team := c.FormValue("team_domain")
		if team == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "domain name is required",
			})
		}

		user := c.FormValue("user_name")
		if user == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "user_name is required",
			})
		}

		err := store.AddSelfTags(user, team, tags)
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, H{
			"text": "Tags added!",
		})
	}
}

func AddSearchTags() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tags []string
		if raw := c.FormValue("text"); raw == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "text parameters are required",
			})
		} else {
			raw = strings.Trim(raw, "[]")
			tags = strings.Split(raw, ",")
		}

		team := c.FormValue("team_domain")
		if team == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "domain name is required",
			})
		}

		user := c.FormValue("user_name")
		if user == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "user_name is required",
			})
		}

		err := store.AddSearchTags(user, team, tags)
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, H{
			"text": "Tags added!",
		})
	}
}

func ClearAllTags() echo.HandlerFunc {
	return func(c echo.Context) error {
		team := c.FormValue("team_domain")
		if team == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "domain name is required",
			})
		}

		user := c.FormValue("user_name")
		if user == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "user_name is required",
			})
		}

		err := store.ClearAllTags(user, team)
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, H{
			"text": "Tags cleared!",
		})
	}
}
