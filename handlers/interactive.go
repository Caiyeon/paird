package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type Payload struct {
	Actions       []map[string]string
	Callback_id   string
	Team          TeamStruct
	Channel       ChannelStruct
	User          UserStruct
	Action_ts     string
	Message_ts    string
	Attachment_id string
	Token         string
	Is_app_unfurl bool
	Response_url  string
}

type TeamStruct struct {
	Id     string
	Domain string
}

type ChannelStruct struct {
	Id   string
	Name string
}

type UserStruct struct {
	Id   string
	Name string
}

func Interactive() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.FormValue("payload")
		if p == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "empty payload",
			})
		}

		// payload is a json object, so let's marshal it
		var payload Payload
		err := json.Unmarshal([]byte(p), &payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": "could not decipher payload",
			})
		}

		switch payload.Callback_id {
		case "signup_self":
			// insert user's self tag
		default:
			// do nothing
		}

		return c.JSON(http.StatusOK, H{
			"text": "Step 2:",
			"attachments": []interface{}{
				map[string]interface{}{
					"text":            "Are you looking for a:",
					"fallback":        "Unsupported slack client",
					"callback_id":     "signup_seek",
					"color":           "#3AA3E3",
					"attachment_type": "default",
					"actions": []interface{}{
						map[string]string{
							"name":  "seek",
							"text":  "Mentor",
							"type":  "button",
							"value": "mentor",
						},
						map[string]string{
							"name":  "seek",
							"text":  "Mentee",
							"type":  "button",
							"value": "mentee",
						},
						map[string]string{
							"name":  "seek",
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