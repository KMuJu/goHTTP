package goHTTP

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequestLine(t *testing.T) {
	cases := []struct {
		name    string
		input   []byte
		method  string
		url     string
		version string
	}{
		{
			name:    "Simple GET /",
			input:   []byte("GET / HTTP/1.1"),
			method:  "GET",
			url:     "/",
			version: "HTTP/1.1",
		},
		{
			name:    "test 2",
			input:   []byte("POST /index.html HTTP/1.1"),
			method:  "POST",
			url:     "/index.html",
			version: "HTTP/1.1",
		},
		{
			name:    "Additional bytes",
			input:   []byte("POST /index.html HTTP/1.1\r\nTesting Data"),
			method:  "POST",
			url:     "/index.html",
			version: "HTTP/1.1",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, m, u, v := parseRequestLine(tc.input)
			assert.Equal(t, tc.method, m)
			assert.Equal(t, tc.url, u)
			assert.Equal(t, tc.version, v)
		})
	}
}
func TestParseHeaders(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected Header
	}{
		{
			name: "Test 1",
			input: []string{
				"Access-Control-Allow-Origin: *",
				"Connection: Keep-Alive",
				"Keep-Alive: timeout=5,max=999",
			},
			expected: Header{
				"Access-Control-Allow-Origin": {"*"},
				"Connection":                  {"Keep-Alive"},
				"Keep-Alive":                  {"timeout=5", "max=999"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output, _ := parseHeaders(tc.input)
			assert.Equal(t, tc.expected, output)
		})
	}
}

func TestHeaderEnd(t *testing.T) {
	cases := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "random data",
			input:    []byte("testing\r\nabs\r\n\r\n"),
			expected: 12,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output := findHeaderEnd(tc.input)
			assert.Equal(t, tc.expected, output)
		})
	}
}

func TestParseRequest(t *testing.T) {
	cases := []struct {
		name     string
		input    []byte
		expected Request
	}{
		{
			name:  "Simple test",
			input: []byte("GET / HTTP/1.1\r\nHost: localhost:8000\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate, br, zstd \r\n\r\n"),

			expected: Request{
				Method:  "GET",
				URL:     "/",
				Version: "HTTP/1.1",
				Header: Header{
					"Host":            {"localhost:8000"},
					"Accept-Language": {"en-US", "en;q=0.5"},
					"Accept-Encoding": {"gzip", "deflate", "br", "zstd"},
				},
				Body: []byte{},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := parseRequest(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, output)
		})
	}
}
