package vault

import (
	"errors"
	"strings"
)

const VaultAddr = "https://vault-ui.io"

var VaultToken string

// given a string, encrypt it with vault and return the cipher
func Encrypt(plaintext string) (string, error) {
	// make call to vault-ui.io
	return "", nil
}

// given a string, decrypt it with vault and return the plaintext
func Decrypt(cipher string) (string, error) {
	if !strings.HasPrefix(cipher, "vault:") {
		return "", errors.New("Not a vault cipher!")
	}
	// make call to vault-ui.io
	return "", nil
}
