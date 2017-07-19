package store

import (
	"errors"
	"fmt"
	"strings"

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
	err := db.Update(func(tx *bolt.Tx) error {
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
		s = fmt.Sprintf("%s", user.Get([]byte("location")))
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

func GetMatchingUsers(username, teamname string) ([]string, error) {
	selfType, err := GetUserKeyValue(username, teamname, "self-type")
	if err != nil {
		return []string{}, err
	}
	searchType, err := GetUserKeyValue(username, teamname, "search-type")
	if err != nil {
		return []string{}, err
	}

	raw, err := GetUserKeyValue(username, teamname, "self-tags")
	if err != nil {
		return []string{}, err
	}
	selfTags := strings.Split(raw, ",")

	raw, err = GetUserKeyValue(username, teamname, "search-tags")
	if err != nil {
		return []string{}, err
	}
	searchTags := strings.Split(raw, ",")

	candidates := make(map[string]int)
	err = db.View(func(tx *bolt.Tx) error {
		teams := tx.Bucket([]byte("teams"))
		if teams == nil {
			return errors.New("'teams' bucket does not exist")
		}
		team := teams.Bucket([]byte(teamname))
		if team == nil {
			return errors.New("team " + teamname + " does not exist")
		}

		// for each user
		c := team.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			// fetch the user's tags
			if v == nil { // if this is a bucket, it must be a user
				// log.Println("testing candidate" + fmt.Sprint("%s", k))
				if fmt.Sprintf("%s", k) == username {
					continue // dont match user with user's self
				}

				// get candidate's list of tags
				user := team.Bucket(k)
				var candidateSelfType string
				var candidateSearchType string
				var candidateSelf []string
				var candidateSearch []string

				if exists := user.Get([]byte("self-type")); exists != nil {
					candidateSelfType = fmt.Sprintf("%s", exists)
				}
				if exists := user.Get([]byte("search-type")); exists != nil {
					candidateSearchType = fmt.Sprintf("%s", exists)
				}
				if exists := user.Get([]byte("self-tags")); exists != nil {
					candidateSelf = strings.Split(fmt.Sprintf("%s", exists), ",")
				}
				if exists := user.Get([]byte("search-tags")); exists != nil {
					candidateSearch = strings.Split(fmt.Sprintf("%s", exists), ",")
				}

				// mentorship MUST match
				if selfType == "mentor" && candidateSearchType == "mentee" {
					continue
				}
				if searchType == "mentee" && candidateSelfType == "mentor" {
					continue
				}

				matches := 0
				// compare with self tags, and add match total
				for _, tag := range selfTags {
					for _, t := range candidateSearch {
						if tag == t {
							matches = matches + 1
							break
						}
					}
				}
				for _, tag := range searchTags {
					for _, t := range candidateSelf {
						if tag == t {
							matches = matches + 1
							break
						}
					}
				}

				// log.Println("adding candidate" + fmt.Sprint("%s", k))
				candidates[fmt.Sprintf("%s", k)] = matches
			}
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}

	// return the top-3 matched candidates
	var result []string
	for _, _ = range []int{1, 2, 3} {
		maxScore := -1
		maxCandidate := ""
		for name, score := range candidates {
			if score > maxScore {
				maxScore = score
				maxCandidate = name
			}
		}
		if maxScore > -1 {
			result = append(result, maxCandidate)
			candidates[maxCandidate] = -1
		}
	}

	return result, nil
}
