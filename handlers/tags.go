package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ListTags() echo.HandlerFunc {
	return func(c echo.Context) error {
		// fetch team's tags:
		// c.FormValue("team_domain")
		// store.GetDomainTags(domain)

		tags := map[string]int{
			"business":    10,
			"drawing":     20,
			"programming": 20,
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
		// raw := c.FormValue("text")
		// raw = strings.Trim(raw, "[]")
		// tags = raw.Split(",")
		// store.AddSelfTags(c.FormValue("team_domain"), c.FormValue("user_name"), tags)

		return c.JSON(http.StatusOK, H{
			"text": "Tags added!",
		})
	}
}

func AddSearchTags() echo.HandlerFunc {
	return func(c echo.Context) error {
		// raw := c.FormValue("text")
		// raw = strings.Trim(raw, "[]")
		// tags = raw.Split(",")
		// store.AddSearchTags(c.FormValue("team_domain"), c.FormValue("user_name"), tags)

		return c.JSON(http.StatusOK, H{
			"text": "Tags added!",
		})
	}
}

func ClearAllTags() echo.HandlerFunc {
	return func(c echo.Context) error {
		// store.ClearAllTags(c.FormValue("team_domain"), c.FormValue("user_name"))

		return c.JSON(http.StatusOK, H{
			"text": "Tags cleared!",
		})
	}
}
