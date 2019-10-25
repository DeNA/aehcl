// +build local

package gcp

import (
	"fmt"
	"context"

	"golang.org/x/oauth2/google"
)

// FetchToken is ...
func FetchToken() (string, error) {
	creds, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to find default credentials. err: %v", err)
	}

	token, err := creds.TokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("failed to fetch token. err: %v", err)
	}
	return token.AccessToken, nils
}