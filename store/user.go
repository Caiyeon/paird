package store

import (
	"errors"
	"fmt"
	"strings"

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

func AddUserIfNotExists(username, teamname string) error {
	return db.Update(func(tx *bolt.Tx) error {
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}
		team := teams.Bucket([]byte(teamname))
		if team == nil {
			return errors.New("team " + teamname + " does not exist")
		}
		_, err := team.CreateBucketIfNotExists([]byte(username))
		return err
	})
}

func AddSelfTags(username, teamname string, tags []string) error {
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

		// key : value => 'tags' : []string{}
		// fetch existing tags
		var prevTags []string
		if exists := user.Get([]byte("self-tags")); exists != nil {
			prevTags = strings.Split(fmt.Sprintf("%s", exists), ",")
		}
		return user.Put([]byte("self-tags"), []byte(strings.Join(append(prevTags, tags...), ",")))
	})
}

func AddSearchTags(username, teamname string, tags []string) error {
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

		// key : value => 'tags' : []string{}
		// fetch existing tags
		var prevTags []string
		if exists := user.Get([]byte("search-tags")); exists != nil {
			prevTags = strings.Split(fmt.Sprintf("%s", exists), ",")
		}
		return user.Put([]byte("search-tags"), []byte(strings.Join(append(prevTags, tags...), ",")))
	})
}

func ClearAllTags(username, teamname string) error {
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
		if err := user.Delete([]byte("self-tags")); err != nil {
			return err
		}
		if err := user.Delete([]byte("search-tags")); err != nil {
			return err
		}
		return nil
	})
}
