package aehcl

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
)

func fetchToken() (string, error) {
	// fetch idToken from metadata of gcp
	if idt, err := fetchIDToken(); err == nil {
		return idt, nil
	}

	// fetch accesstoken from local `GOOGLE_APPLICATION_CREDENTIALS`
	lat, err := fetchLocalAccessToken()
	if err != nil {
		return "", err
	}
	return lat, nil
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
