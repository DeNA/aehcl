package gcp

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
)

var idTokenURL = "/instance/service-accounts/default/identity?audience=" + os.Getenv("GOOGLE_CLOUD_PROJECT")

// TokenSource returns token.
type TokenSource func() (string, error)

// Token returns function TokenSource implementation.
func Token() TokenSource {
	return func() (string, error) {
		return fetchToken()
	}
}

func fetchToken() (string, error) {
	// get idToken from metadata of gcp
	idToken, err := fetchIDToken()
	if err != nil || idToken == "" {
		// get accesstoken from local `GOOGLE_APPLICATION_CREDENTIALS`
		lat, err := fetchLocalAccessToken()
		if err != nil {
			return "", err
		}
		return lat, nil
	}
	return idToken, nil
}

func fetchIDToken() (string, error) {
	return metadata.Get("/instance/service-accounts/default/identity?audience=" + os.Getenv("GOOGLE_CLOUD_PROJECT"))
}

func fetchLocalAccessToken() (string, error) {
	creds, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to find default credentials. err: %v", err)

	}

	token, err := creds.TokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("failed to fetch token. err: %v", err)

	}

	return token.AccessToken, nil
}
