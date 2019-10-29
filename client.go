// Package aehcl provides service-to-service authentication in Google App Engine.
package aehcl

import (
	"net/http"
)

type transport struct {
	base http.RoundTripper
}

// Transport is an implementation of RoundTripper required service-to-service authentication.
// If base http RoundTripper is nil, it sets DefaultTransport.
func Transport(base http.RoundTripper) http.RoundTripper {
	t := &transport{
		base: base,
	}
	if base == nil {
		t.base = http.DefaultTransport
	}
	return t
}

func (t *transport) RoundTrip(ireq *http.Request) (*http.Response, error) {
	token, err := t.token()
	if err != nil {
		return nil, err
	}

	req := new(http.Request)
	*req = *ireq

	if req.Header == nil {
		req.Header = make(http.Header)
	}
	req.Header = cloneHeader(req.Header)
	req.Header.Set("Authorization", "Bearer "+token)

	return t.base.RoundTrip(req)
}

func (t *transport) token() (string, error) {
	return fetchToken()
}

func cloneHeader(h http.Header) http.Header {
	nv := 0
	for _, v := range h {
		nv += len(v)
	}

	sv := make([]string, nv) // shared backing array for headers' values
	h2 := make(http.Header, len(h))
	for k, v := range h {
		n := copy(sv, v)
		h2[k] = sv[:n:n]
		sv = sv[:n]
	}
	return h2
}
