# App Engine Http Client

[![GoDoc](https://godoc.org/github.com/emahiro/aehcl?status.svg)](https://godoc.org/github.com/emahiro/aehcl)
[![Go Report Card](https://goreportcard.com/badge/github.com/emahiro/aehcl)](https://goreportcard.com/report/github.com/emahiro/aehcl)

## Description

App Engine HTTP Client provides HTTP RoundTripper for authentication service-to-service in Google App Engine.  
This package is inspired by [Authentication Service-to-Service](https://cloud.google.com/run/docs/authenticating/service-to-service)

## Usage

```go

client := &http.Client {
    Transport: aehcl.Transport(http.DefaultTransport)
}

```

## License

MIT
