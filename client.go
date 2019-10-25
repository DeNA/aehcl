package aehcl

import (
	"net/http"
)

// NewAEHttpClient is ...
func NewAEHttpClient(t *http.Transport, token string) *http.Client {
	return &http.Client{
		Transport: NewAEHttpTransport(token),
	}
}

// AEHttpTransport ...
type AEHttpTransport struct {
	base  http.RoundTripper
	Token string
}

// NewAEHttpTransport ...
func NewAEHttpTransport(token string) *AEHttpTransport {
	return &AEHttpTransport{
		base:  http.DefaultClient.Transport,
		Token: token,
	}
}

// RoundTrip ...
func (t *AEHttpTransport) RoundTrip(ireq *http.Request) (*http.Response, error) {
	req := new(http.Request)
	*req = *ireq // shallow clone request

	if req.Header == nil {
		req.Header = make(http.Header)
	}
	req.Header.Set("Authorization", "Bearer "+t.Token)
	return t.base.RoundTrip(req)
}
