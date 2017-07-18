package store

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func Initialize(filename string) (err error) {
	// use local file as db store
	DB, err = bolt.Open(filename, 0600, nil)

	// start a transaction
	tx, err := DB.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// ensure root buckets exist
	if _, err := tx.CreateBucketIfNotExists([]byte("teams")); err != nil {
		return fmt.Errorf("create bucket %s error: %s", "teams", err)
	}
	if _, err := tx.CreateBucketIfNotExists([]byte("stats")); err != nil {
		return fmt.Errorf("create bucket %s error: %s", "stats", err)
	}

	// commit transaction and check for error
	err = tx.Commit()
	return
}

func CloseDB() {
	DB.Close()
}
