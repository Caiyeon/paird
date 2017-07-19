package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/caiyeon/lunch-with-us/store"
)

const sleepDuration = 15 * time.Second

// run this as a background job to find a pairing after a certain amount of time has passed
func FindPair(username, teamname, webhook string) {
	log.Println("goroutine started...")

	// delay response
	time.Sleep(sleepDuration)

	log.Println("finding pair...")

	// find matches
	matches, err := store.GetMatchingUsers(username, teamname)
	if err != nil {
		log.Println(err.Error())
		return
	}

	text := ""
	if len(matches) == 0 {
		text = "No matches found so far..."
	} else {
		text = "Your top matches are: " + strings.Join(matches, ", ")
	}

	log.Println(text)

	// alert user about match
	payload, err := json.Marshal(
		map[string]interface{}{
			"response_type":    "ephemeral",
			"replace_original": false,
			"text":             text,
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
