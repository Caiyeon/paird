package store

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
)

func SetUserKeyValue(username, teamname, key, value string) error {
	if username == "" || teamname == "" || key == "" || value == "" {
		return errors.New("SetUserKeyValue: no parameters can be empty")
	}

	return db.Update(func(tx *bolt.Tx) error {
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}
		team := teams.Bucket([]byte(teamname))
		if team == nil {
			return errors.New("team " + teamname + " does not exist")
		}
		user, err := team.CreateBucketIfNotExists([]byte(username))
		if err != nil {
			return err
		}
		return user.Put([]byte(key), []byte(value))
	})
}

func GetUserKeyValue(username, teamname, key string) (string, error) {
	if username == "" || teamname == "" || key == "" {
		return "", errors.New("GetUserKeyValue: no parameters can be empty")
	}

	s := ""
	err := db.View(func(tx *bolt.Tx) error {
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}
		team := teams.Bucket([]byte(teamname))
		if team == nil {
			return errors.New("team " + teamname + " does not exist")
		}
		user, err := team.CreateBucketIfNotExists([]byte(username))
		if err != nil {
			return err
		}
		s = fmt.Sprintf("%s", user.Get([]byte("diet")))
		return nil
	})
	return s, err
}

func DeleteUserKeyValue(username, teamname, key string) error {
	if username == "" || teamname == "" || key == "" {
		return errors.New("DeleteUserKeyValue: no parameters can be empty")
	}

	return db.Update(func(tx *bolt.Tx) error {
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}
		team := teams.Bucket([]byte(teamname))
		if team == nil {
			return errors.New("team " + teamname + " does not exist")
		}
		user, err := team.CreateBucketIfNotExists([]byte(username))
		if err != nil {
			return err
		}
		return user.Delete([]byte(key))
	})
}
