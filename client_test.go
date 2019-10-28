package aehcl

import (
	"net/http"
	"strings"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	client := &http.Client{
		Transport: Transport(http.DefaultTransport),
	}
	req, _ := http.NewRequest("GET", "/", nil)
	client.Do(req)

	if h := req.Header.Get("Authorization"); h == "" || strings.Index(h, "Bearer ") == -1 {
		t.Fatalf("Authirization is not found. header: %#v", h)
	}
}

func Benchmark_cloneHeader(b *testing.B) {
	header := http.Header{}
	for i := 0; i < 10000; i++ {
		header.Add(string(i), string(i))
	}

	for i := 0; i < b.N; i++ {
		cloneHeader(header)
	}
}

func Benchmark_cloneHeaderV1(b *testing.B) {
	header := http.Header{}
	for i := 0; i < 10000; i++ {
		header.Add(string(i), string(i))
	}

	for i := 0; i < b.N; i++ {
		cloneHeaderV1(header)
	}
}

func Benchmark_cloneHeaderV2(b *testing.B) {
	header := http.Header{}
	for i := 0; i < 10000; i++ {
		header.Add(string(i), string(i))
	}

	for i := 0; i < b.N; i++ {
		cloneHeaderV2(header)
	}
}
