package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/caiyeon/lunch-with-us/store"
	"github.com/labstack/echo"
)

type Payload struct {
	Actions       []ActionStruct
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

type ActionStruct struct {
	Name  string
	Type  string
	Value string
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

		text := ""
		switch payload.Callback_id {
		case "signup_self":
			if payload.Actions[0].Value == "mentor" || payload.Actions[0].Value == "mentee" {
				text = "Ok, you are signing up as a " + payload.Actions[0].Value
			} else {
				text = "Ok, you don't mind being either"
				payload.Actions[0].Value = "both"
			}
			text = text + "\nAre you looking for a:"

			store.SetUserKeyValue(payload.User.Name, payload.Team.Domain, "self-type", payload.Actions[0].Value)

			return c.JSON(http.StatusOK, H{
				"text": "Step 2:",
				"attachments": []interface{}{
					map[string]interface{}{
						"text":            text,
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

		case "signup_seek":
			if payload.Actions[0].Value == "mentor" || payload.Actions[0].Value == "mentee" {
				text = "Ok, you are seeking a " + payload.Actions[0].Value
			} else {
				text = "Ok, you don't mind meeting up with a mentor or mentee"
				payload.Actions[0].Value = "both"
			}

			store.SetUserKeyValue(payload.User.Name, payload.Team.Domain, "search-type", payload.Actions[0].Value)

			go FindPair(payload.User.Name, payload.Team.Domain, payload.Response_url)

			return c.JSON(http.StatusOK, H{
				"text":        text + "\nDone! You will be notified when a pairing is made!\nTo enhance your experience:",
				"attachments": helpMessage,
			})

		default:
			// no callback found, return nothing in fallthrough
		}

		return c.JSON(http.StatusOK, H{})
	}
}
