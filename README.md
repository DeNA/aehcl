# App Engine Http Client

[![GoDoc](https://godoc.org/github.com/DeNA/aehcl?status.svg)](https://godoc.org/github.com/DeNA/aehcl)
[![Go Report Card](https://goreportcard.com/badge/github.com/DeNA/aehcl)](https://goreportcard.com/report/github.com/DeNA/aehcl)

## Description

App Engine HTTP Client provides HTTP RoundTripper for authentication service-to-service in Google App Engine.  
This package is inspired by [Authentication Service-to-Service](https://cloud.google.com/run/docs/authenticating/service-to-service)

## Installation

```sh
go get github.com/DeNA/aehcl
```

## Usage

```go

client := &http.Client {
    Transport: aehcl.Transport(nil)
}

```

## License

MIT
