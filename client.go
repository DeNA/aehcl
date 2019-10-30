// Package aehcl provides service-to-service authentication in Google App Engine.
package aehcl

import (
	"net/http"
)

// TokenSource is interface of function returns token required service-to-service authentication in App Engine.
type TokenSource func() (string, error)

// TokenSourceOption is interface of function that sets token source required service-to-service authentication.
type TokenSourceOption func(*tokenSourceOption)

type tokenSourceOption struct {
	token TokenSource
}

// WithTokenSource sets token source required service-to-service authentication to tokenSourceOption.
func WithTokenSource(ts TokenSource) TokenSourceOption {
	return func(o *tokenSourceOption) {
		o.token = ts
	}
}

func (o *tokenSourceOption) apply(opts []TokenSourceOption) {
	for _, opt := range opts {
		opt(o)
	}
}

type transport struct {
	base  http.RoundTripper
	token TokenSource
}

// Transport is an implementation of http.RoundTripper for service-to-service authentication.
// When required service-to-service authentication, create http.Client using this transport.
//
// Default RoundTripper is http.DefaultTransport, and default TokenSourceOption is FetchIDToken.
func Transport(base http.RoundTripper, opts ...TokenSourceOption) http.RoundTripper {
	opt := &tokenSourceOption{token: FetchIDToken}
	opt.apply(opts)

	t := &transport{
		base:  base,
		token: opt.token,
	}
	if base == nil {
		t.base = http.DefaultTransport
	}
	return t
}

// RoundTrip issues a request with identity token required service-to-service authentication described in
// https://cloud.google.com/run/docs/authenticating/service-to-service.
// When failed to obtain the identity token from metadata API (e.g. in local environment), RoundTrip returns error.
//
// If uses service-to-serivce authentication, server that receives the request must be implemented to validate the
// identity token added to Authorization header using the public key provided by Google.
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
