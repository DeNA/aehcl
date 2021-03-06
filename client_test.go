package aehcl

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		arg     http.RoundTripper
		handler http.Handler
	}{
		{
			name: "success to get authorization header",
			arg:  Transport(http.DefaultTransport, WithTokenSource(FetchIDToken)),
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if h := r.Header.Get("Authorization"); h == "" {
					t.Fatalf("Authorization Header is required")
				}
			}),
		},
		{
			name: "success to get authorization header in empty options",
			arg:  Transport(http.DefaultTransport, []Option{}...),
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
