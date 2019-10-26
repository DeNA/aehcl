package aehcl

import (
	"net/http"

	"github.com/emahiro/aehcl/internal/gcp"
)

var tokenSource = gcp.Token()

// Transport is aehcl transport.
type Transport struct {
	base http.RoundTripper
}

// RoundTrip retrieves token from token source and set it to request header.
func (t *Transport) RoundTrip(ireq *http.Request) (*http.Response, error) {
	token, err := tokenSource()
	if err != nil {
		return nil, err
	}

	req := new(http.Request)
	*req = *ireq

	if req.Header == nil {
		req.Header = make(http.Header)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return t.base.RoundTrip(req)
}
