package goHTTP

import (
	"fmt"
	"net"
)

type ResponseWriter interface {
	// Used to get access to the header map.
	// Changes to the header must be done before Write.
	Header() Header
	// Will write the bytes to the connection.
	// Should call WriteHeader(200) if it has not been called before.
	Write(b []byte) (int, error)
	// Will write the header and code to the connection.
	WriteHeader(code Code)
}

type Response struct {
	request     *Request
	version     string
	statuscode  Code
	status      string
	header      Header
	wroteCode   bool
	wroteHeader bool
	conn        net.Conn
}

func newResponseWriter(conn net.Conn, r *Request) *Response {
	return &Response{
		request: r,
		header:  Header{},
		conn:    conn,
		version: r.Version,
	}
}

func (response *Response) Header() Header {
	return response.header
}

func (response *Response) Write(b []byte) (int, error) {
	if !response.wroteCode {
		response.WriteHeader(Ok)
	}
	return response.conn.Write(b)
}

func (r *Response) WriteHeader(code Code) {
	if r.wroteHeader { // should not be able to write the header twice
		return
	}
	r.statuscode = code
	r.status = "ok"
	r.wroteCode = true

	r.conn.Write([]byte(fmt.Sprintf("%s %d %s\r\n", r.version, r.statuscode, r.status)))
	r.conn.Write([]byte(r.header.String()))
}
