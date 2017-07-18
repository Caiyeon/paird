package store

import (
	"errors"

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
