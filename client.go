// package aehcl provides HTTP RoundTripper for authentication service-to-service
// in Google App Engine.

package aehcl

import (
	"net/http"
)

type transport struct {
	base  http.RoundTripper
	token tokenSource
}

// Transport is an implementation of RoundTripper with TokenSource required authentication service-to-service.
// If base http RoundTripper is nil, it sets DefaultTransport.
func Transport(base http.RoundTripper) http.RoundTripper {
	t := &transport{
		base:  base,
		token: token(),
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
