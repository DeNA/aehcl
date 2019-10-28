package aehcl

import (
	"net/http"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	client := &http.Client{
		Transport: Transport(http.DefaultTransport),
	}
	req, _ := http.NewRequest("GET", "/", nil)
	client.Do(req)

	if h := req.Header.Get("Authorization"); h != "" {
		t.Fatal("Authorization shoud be concealed.")
	}
}
