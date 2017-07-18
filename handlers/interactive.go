package handlers

import (
	"encoding/json"
	"net/http"

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
			}
			text = text + "\nAre you looking for a:"

			// insert user's self tag

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
			}
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

var helpMessage = []interface{}{
	map[string]interface{}{
		"title":           "Add some tags to yourself",
		"text":            "`/listtags` _returns a list of popular tags in your team_\n`/addmytags [business, software development, ...]` _describe yourself!_\n`/addsearchtags [marketing, ...]` _what are you looking for specifically?_\n`/clearalltags` _clears your own tags and search tags_",
		"mrkdwn_in":       []string{"text"},
		"fallback":        "Unsupported slack client",
		"callback_id":     "signup_self",
		"color":           "#3AA3E3",
		"attachment_type": "default",
	},
	map[string]interface{}{
		"title":           "Add meetup preferences",
		"text":            "`/mydiet [none|vegetarian|vegan]` _add dietary restrictions_\n`/myavailability [9AM-2PM] [Monday,Tuesday,Wednesday]` _restrict meetup times_\n`/mylocation [Vancouver]` _set your location for meetup suggestions_",
		"mrkdwn_in":       []string{"text"},
		"fallback":        "Unsupported slack client",
		"callback_id":     "signup_self",
		"color":           "#008000",
		"attachment_type": "default",
	},
	map[string]interface{}{
		"title":           "To see these options again",
		"text":            "`/paird` _displays help message_",
		"mrkdwn_in":       []string{"text"},
		"fallback":        "Unsupported slack client",
		"callback_id":     "signup_self",
		"color":           "#DAF7A6",
		"attachment_type": "default",
	},
}
