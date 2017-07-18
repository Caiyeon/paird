package store

import (
	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func CreateAndOpen(filename string) (err error) {
	DB, err = bolt.Open(filename, 0600, nil)
	return
}

func CloseDB() {
	DB.Close()
}
