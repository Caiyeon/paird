package store

import (
	"errors"

	"github.com/boltdb/bolt"
)

func DoesUserExistInTeam(username, teamname string) (bool, error) {
	found := false

	// open a read-only transaction
	if err := db.View(func(tx *bolt.Tx) error {
		// get the root teams bucket
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}

		// if the sub bucket returns nil, then team does not exist
		team := teams.Bucket([]byte(teamname))
		if team == nil {
			// if team doesn't exist, user doesn't exist
			found = false
			return nil
		}

		found = team.Bucket([]byte(username)) != nil
		return nil
	}); err != nil {
		return false, err
	}

	// no errors occured using database, return found flag
	return found, nil
}

func AddUserIfNotExists(username, teamname string, properties map[string]string) error {
	return nil
}
