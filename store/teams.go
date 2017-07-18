package store

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
)

// returns whether teamname (aka domain name) exists in db
func DoesTeamExist(teamname string) (bool, error) {
	// start a transaction
	tx, err := db.Begin(true)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	// find teams bucket
	teamsBucket := tx.Bucket([]byte("teams"))
	if teamsBucket == nil {
		return false, errors.New("'teams' bucket does not exist")
	}

	// search in teams bucket for specific team name
	cursor := teamsBucket.Cursor()
	found := false
	for key, value := cursor.First(); key != nil && !found; key, value = cursor.Next() {
		// buckets have value equal to nil
		if value == nil && fmt.Sprintf("%s", key) == teamname {
			found = true
		}
	}

	// commit transaction and check for error
	err = tx.Commit()
	return found, err
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
