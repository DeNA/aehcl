package gcp

import (
	"fmt"
	"testing"
)

func Test_fetchLocalAccessToken(t *testing.T) {
	token, err := fetchLocalAccessToken()
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	fmt.Printf("token: %s\n", token)
	if token == "" {
		t.Fatalf("faild to get gcp credentilas in GOOGLE_APPLICATION_CREDENTIALS. err: %v\n", err)
	}
}
