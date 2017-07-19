package store

import (
	"errors"
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
)

// returns whether teamname (aka domain name) exists in db
func DoesTeamExist(teamname string) (bool, error) {
	// found flag keeps track of whether the team name has been found in db
	found := false

	// open a read-only transaction
	if err := db.View(func(tx *bolt.Tx) error {
		// get the root teams bucket
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}

		// if the sub bucket returns nil, then team does not exist
		found = teams.Bucket([]byte(teamname)) != nil
		return nil
	}); err != nil {
		return false, err
	}

	// no errors occured using database, return found flag
	return found, nil
}

// if team exists, do nothing. If team doesn't exist, create a bucket for it
func AddTeamIfNotExists(teamname string) error {
	return db.Update(func(tx *bolt.Tx) error {
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}
		_, err := teams.CreateBucketIfNotExists([]byte(teamname))
		return err
	})
}

func ListTeamTags(teamname string) (map[string]int, error) {
	tags := make(map[string]int)

	err := db.View(func(tx *bolt.Tx) error {
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}
		team := teams.Bucket([]byte(teamname))

		// for each user
		c := team.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			// fetch the user's tags
			if v == nil { // if this is a bucket, it must be a user
				user := team.Bucket(k)
				var usertags []string
				if exists := user.Get([]byte("self-tags")); exists != nil {
					usertags = strings.Split(fmt.Sprintf("%s", exists), ",")
				}
				if exists := user.Get([]byte("search-tags")); exists != nil {
					usertags = append(usertags, strings.Split(fmt.Sprintf("%s", exists), ",")...)
				}

				// add to tags list
				for _, tag := range usertags {
					if count, ok := tags[tag]; !ok {
						tags[tag] = 1
					} else {
						tags[tag] = count + 1
					}
				}
			}
		}
		return nil
	})

	return tags, err
}

func SetTeamWebhook(teamname, webhook string) error {
	if teamname == "" || webhook == "" {
		return errors.New("Teamname and webhook cannot be empty")
	}
	return db.Update(func(tx *bolt.Tx) error {
		webhooks, err := tx.CreateBucketIfNotExists([]byte("webhooks"))
		if err != nil {
			return err
		}
		return webhooks.Put([]byte(teamname), []byte(webhook))
	})
}

func GetTeamWebhook(teamname string) (string, error) {
	if teamname == "" {
		return "", errors.New("Teamname cannot be empty")
	}

	s := ""
	err := db.Update(func(tx *bolt.Tx) error {
		webhooks, err := tx.CreateBucketIfNotExists([]byte("webhooks"))
		if err != nil {
			return err
		}
		s = fmt.Sprintf("%s", webhooks.Get([]byte(teamname)))
		return nil
	})

	return s, err
}
