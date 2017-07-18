package main
import (
// "github.com/caiyeon/lunch-with-us/vault"
"github.com/caiyeon/lunch-with-us/yelp"
"time"


)
// main page to test functions
func main(){

// vault.Decrypt("vault:v1:oXS9ezMRv6UFulP9n2q4DSZ2WPc9JyuSTTzEbByXaGg=")
yelp.GetSuggestions(yelp.Location{},time.Time{},[]string{})
}