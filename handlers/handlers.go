package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// insert handler functions here
// to test temporary handler functions, make another .go file to avoid git conflicts

// for returning JSON bodies
type H map[string]interface{}

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

func DisplayHelpMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"text":        "You will be notified when a pairing is made!\nTo enhance your experience:",
			"attachments": helpMessage,
		})
	}
}
