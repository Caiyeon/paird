package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/caiyeon/paird/store"
	"github.com/caiyeon/paird/yelp"
)

const sleepDuration = 5 * time.Second

// run this as a background job to find a pairing after a certain amount of time has passed
func FindPair(username, teamname, webhook string) {
	log.Println("goroutine started...")

	// delay response
	time.Sleep(sleepDuration)

	// find matches
	matches, err := store.GetMatchingUsers(username, teamname)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(matches) == 0 {
		payload, err := json.Marshal(
			map[string]interface{}{
				"response_type":    "ephemeral",
				"replace_original": false,
				"text":             "No matches found so far...",
			},
		)
		if err != nil {
			log.Println(err.Error())
			return
		}
		_, err = http.Post(webhook, "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// if user has location, use yelp api for suggestions
	details := []string{}
	for _, m := range matches {
		text := ""
		loc, err := store.GetUserKeyValue(username, teamname, "location")
		if err != nil {
			details = append(details, "")
			continue
		}
		loc2, err := store.GetUserKeyValue(m, teamname, "location")
		if err != nil {
			details = append(details, "")
			continue
		}
		if loc == loc2 {
			text = text + "You are both located in *" + loc + "*."
		}
		log.Println("Location: " + loc)

		// optionally include dietary options
		res, err := yelp.GetSuggestions(yelp.Location{City: loc}, time.Time{}, []string{})
		if err != nil {
			details = append(details, "")
			continue
		}

		text = text + " Here's a list of nearby places that fit *both* of your dietary restrictions!\n"
		for _, r := range res {
			text = text + "\t" + r
		}
		details = append(details, text)
	}

	var attachments []interface{}
	for i, m := range matches {
		attachments = append(attachments, map[string]interface{}{
			"title":           "Match #" + strconv.Itoa(i) + ": *@" + m + "*",
			"text":            details[i],
			"mrkdwn_in":       []string{"text"},
			"fallback":        "Unsupported slack client",
			"color":           "#3AA3E3",
			"attachment_type": "default",
		})
	}

	// alert user about match
	payload, err := json.Marshal(
		map[string]interface{}{
			"response_type":    "ephemeral",
			"replace_original": false,
			"text":             "Pairings have been found for you!",
			"attachments":      attachments,
		},
	)
	if err != nil {
		log.Println(err.Error())
		return
	}

	_, err = http.Post(webhook, "application/json", bytes.NewReader(payload))
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Successfully paired!")
	return
}
