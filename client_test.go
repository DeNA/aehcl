package aehcl

import (
	"net/http"
	"strings"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	client := &http.Client{
		Transport: &Transport{
			base: http.DefaultTransport,
		},
	}

	req, _ := http.NewRequest("GET", "/", nil)
	client.Do(req)

	if h := req.Header.Get("Authorization"); h == "" || strings.Index(h, "Bearer ") == -1 {
		t.Fatalf("Authirization is not found. header: %#v", h)
	}
}
