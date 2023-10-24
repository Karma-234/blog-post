package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization header: ApiKey {apikey here}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")

	if val == "" {
		return "", errors.New("No authorization header found!")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed authorization header!")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of authorization header!")
	}
	return vals[1], nil
}
