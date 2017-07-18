package yelp

import (
	"strings"
	"time"
)

var (
	AccessToken string
)

type Location struct {
	city string
	long float64
	lat  float64
}

func GetSuggestions(loc Location, t time.Time, opts []string) ([]string, error) {
	AccessToken = "abc"
	if strings.HasPrefix(AccessToken, "vault:v1") {
		// use vault to decipher access token
	}

	// fetch list of resturants from yelp API with access token

	return []string{"starbucks", "mcdonalds", "tim hortons"}, nil
}
