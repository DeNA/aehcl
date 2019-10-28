package aehcl

import (
	"net/http"
)

// Transport is aehcl transport.
type transport struct {
	base  http.RoundTripper
	token tokenSource
}

// Transport ...
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

// RoundTrip retrieves token from token source and set it to request header.
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

func cloneHeaderV1(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, v := range h {
		v2 := make([]string, len(v))
		copy(v2, v)
		h2[k] = v2
	}
	return h2
}

func cloneHeaderV2(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, v := range h {
		h2[k] = v
	}
	return h2
}
