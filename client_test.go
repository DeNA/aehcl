package aehcl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		arg     http.RoundTripper
		ts      func() (string, error)
		handler http.Handler
	}{
		{
			name: "success to get authorization header",
			arg:  Transport(http.DefaultTransport, FetchIDToken),
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if h := r.Header.Get("Authorization"); h == "" {
					t.Fatalf("Authorization Header is required")
				}
			}),
		},
		{
			name: "faield to get authorization header",
			arg:  &http.Transport{},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if h := r.Header.Get("Authorization"); h != "" {
					t.Fatalf("Authorization Header is exist. header: %v", h)
				}
			}),
		},
		{
			name: "faield to get idToken",
			arg: Transport(http.DefaultTransport, func() (string, error) {
				return "", fmt.Errorf("hoge")
			}),
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r != nil {
					t.Fatalf("request should be failed")
				}
			}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{
				Transport: tt.arg,
			}
			server := httptest.NewServer(tt.handler)
			defer server.Close()

			client.Get(server.URL)
		})
	}
}
