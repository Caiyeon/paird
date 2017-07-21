package handlers

import (
	"net/http"

	"github.com/caiyeon/paird/store"
	"github.com/labstack/echo"
)

func Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		if team := c.FormValue("team_domain"); team == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "team domain cannot be empty",
			})
		} else {
			store.AddTeamIfNotExists(team)
		}

		return c.JSON(http.StatusOK, H{
			"text": "Welcome!",
			"attachments": []interface{}{
				map[string]interface{}{
					"text":            "Do you want to be a:",
					"fallback":        "Unsupported slack client",
					"callback_id":     "signup_self",
					"color":           "#3AA3E3",
					"attachment_type": "default",
					"actions": []interface{}{
						map[string]string{
							"name":  "self",
							"text":  "Mentor",
							"type":  "button",
							"value": "mentor",
						},
						map[string]string{
							"name":  "self",
							"text":  "Mentee",
							"type":  "button",
							"value": "mentee",
						},
						map[string]string{
							"name":  "self",
							"text":  "Doesn't matter",
							"type":  "button",
							"value": "",
						},
					},
				},
			},
		})
	}
}
