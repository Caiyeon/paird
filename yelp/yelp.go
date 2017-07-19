package yelp

import (
	"strings"
	"time"
	"github.com/caiyeon/lunch-with-us/vault"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"strconv"
	"encoding/json"
)


var (
	AccessToken string
)

type Location struct {
	City string
	Long float64
	Lat  float64
}

type Businesses struct {
	Restos []Restaurant `json:"businesses"`
}

type Restaurant struct {
	Name string `json:"name"`  
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
	// check if location is included
	if (loc.City == "" && loc.Long == 0){
		fmt.Println("Location required!")
	}

	client := &http.Client{}

	urlPrefix := "https://api.yelp.com/v3/businesses/search?"
	
	// location string is given
	if loc.City != ""{
		urlPrefix = urlPrefix + "location=" + loc.City
	}else if loc.Lat != 0 && loc.Long != 0  {
		urlPrefix = urlPrefix + "latitude=" + strconv.FormatFloat(loc.Lat,'f',3,64)
		urlPrefix = urlPrefix + "&longitude=" + strconv.FormatFloat(loc.Long,'f',3,64)
	}

	// limit response to 3
	urlPrefix += "&limit=3"
	// add time to query
	comp :=time.Time{}
	if t != comp{
	urlPrefix += "open_at" + string(t.Unix())
	}

	//format link with given options
	if len(opts) != 0 {
		urlPrefix += "&term="
		for _, option := range opts {
			urlPrefix += option
		}
	}
	fmt.Println(urlPrefix)

	req, err := http.NewRequest("GET", urlPrefix, nil)
    if err != nil {
        log.Print(err)
    }
    req.Header.Add("authorization", "Bearer " + AccessToken)
    resp, err := client.Do(req)
    defer resp.Body.Close()
    fmt.Println("Yelp request status is ",resp.Status)
    respString, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Respond string is: ", string(respString))

    // extract response string into structs
    stripped := Businesses{}
    json.Unmarshal([]byte(respString),&stripped)
    fmt.Println("Parsed text is: ", stripped)
    var result []string 

    // convert restaurant objects into string array
    for _, rest := range stripped.Restos {
    	result = append(result, rest.Name)
    }
	return result, nil
}
