package vault

import (
	"errors"
	"strings"
	"net/http"
	"net/url"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"

)

const VaultAddr = "https://vault-ui.io"

var VaultToken="ca13caaa-6ff6-a6e1-8770-e0b107719a8b"
type encryptData struct {
	Result string `json:"result"`
}

// given a string, encrypt it with vault and return the cipher
func Encrypt(text string) (string, error) {
	// make call to vault-ui.io
	client := &http.Client{}
	payload := strings.NewReader("key=lwu&plaintext="+ url.QueryEscape(text))
	req, err := http.NewRequest("POST", "https://vault-ui.io/api/transit/encrypt", payload)
    if err != nil {
        log.Print(err)
    }
    req.Header.Add("X-Vault-Token", VaultToken)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp, err := client.Do(req)
    defer resp.Body.Close()
    fmt.Println("Status is ",resp.Status)
    respString, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response: ", string(respString))

    // extact response string
    stripped := encryptData{}
    json.Unmarshal([]byte(respString),&stripped)
    fmt.Println("Encryption token is: ",stripped.Result)
	return stripped.Result, nil
}

// given a string, decrypt it with vault and return the plaintext
func Decrypt(cipher string) (string, error) {
	// check if the given cipher is valid
	if !strings.HasPrefix(cipher, "vault:") {
		return "", errors.New("Not a vault cipher!")
	}
	// make call to vault-ui.io
	client := &http.Client{}
	payload := strings.NewReader("key=lwu&cipher="+ url.QueryEscape(cipher))
	req, err := http.NewRequest("POST", "https://vault-ui.io/api/transit/decrypt", payload)
    if err != nil {
        log.Print(err)
        return "", err
    }
    req.Header.Add("X-Vault-Token", VaultToken)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp, err := client.Do(req)
    if err != nil{
    	return "", err
    }
    defer resp.Body.Close()
    fmt.Println("vault Status is ",resp.Status)
    respString, err := ioutil.ReadAll(resp.Body)
    if err != nil{
    	return "", err
    }
    fmt.Println("Decrypted result: ", string(respString))

    //extract response string
    stripped := encryptData{}
    json.Unmarshal([]byte(respString),&stripped)
    fmt.Println("Decrypted text is: ", stripped.Result)
	return stripped.Result, nil
}
