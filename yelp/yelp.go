package yelp

import (
	"strings"
	"time"
	"github.com/caiyeon/lunch-with-us/vault"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
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
	AccessToken = "vault:v1:ZI/vH4d8aRcPmYHd4+FBrkKgnpQHCMCGU4JapfzkKr9Wb4XI3ol+mYi6DVaoFWrTP+dNct6EHo6EU4rTtLcgGkuIGxdfQ58IHgZd96VCwWsl28W4q85nUR5PjFBm0pKLCEAk+i0WS+1B9JF7kt6nkihVs1Ja5F/oyy2FD9+rTbT7QvUSQsvBtprSKyvWj0xEj+AbiFUiOSRbmUvI"
	if strings.HasPrefix(AccessToken, "vault:v1") {
		// use vault to decipher access token
		Decrypted, err := vault.Decrypt(AccessToken)
		if err != nil{
			return []string{}, err
		}
		AccessToken = Decrypted
	}
	fmt.Println(AccessToken)
	// fetch list of resturants from yelp API with access token
	// if (location.city == nil & location.long == nil){
	// 	fmt.Println("Location required!")
	// }

	client := &http.Client{}

	// // location string is given
	// if (location.city != nil){
	// 	locaString = location.city
	// }
	
	// payload := strings.NewReader("key=lwu&cipher="+ cipher)
	req, err := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search?location=vancouver", nil)
    if err != nil {
        log.Print(err)
    }
    req.Header.Add("authorization", "Bearer " + AccessToken)
    // req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp, err := client.Do(req)
    defer resp.Body.Close()
    fmt.Println("Status is ",resp.Status)
    respString, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Decrypted result: ", string(respString))

    //extract response string
    // stripped := encryptData{}
    // json.Unmarshal([]byte(respString),&stripped)
    // fmt.Println("Decrypted text is: ", stripped.Result)
	return []string{"starbucks", "mcdonalds", "tim hortons"}, nil
}
