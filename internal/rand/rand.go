package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil || nRead != n {
		return nil, fmt.Errorf("rand.Read: %v", err)
	}
	return b, nil
}

func randString(n int) (string, error) {
	b, err := bytes(n)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// DefaultSessionToken generates a 32 bit token string and returns it along with an error
func DefaultSessionToken() (string, error) {
	return randString(32)
}
