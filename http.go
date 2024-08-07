package goHTTP

import (
	"fmt"
	"strings"
)

type method string

type message struct {
	raw []byte
}

type Header map[string][]string

type Request struct {
	Method  string
	URL     string
	Version string
	Header  Header
	Body    []byte
}

type RequestWriter struct {
	bytes []byte
}

func (r Request) String() string {
	builder := strings.Builder{}
	builder.WriteString(r.Method)
	builder.WriteString(" ")
	builder.WriteString(r.URL)
	builder.WriteString(" ")
	builder.WriteString(r.Version)
	builder.WriteString("\r\n")
	for k, v := range r.Header {
		builder.WriteString(fmt.Sprintf("%s: %s\r\n", k, strings.Join(v, ", ")))
	}
	builder.WriteString("\r\n")
	builder.Write(r.Body)
	return builder.String()
}

func (h Header) String() string {
	builder := strings.Builder{}
	for k, v := range h {
		builder.WriteString(fmt.Sprintf("%s: %s\r\n", k, strings.Join(v, ", ")))
	}
	builder.WriteString("\r\n")

	return builder.String()
}
