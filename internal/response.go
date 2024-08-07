package internal

import (
	"fmt"
	"net"

	"github.com/kmuju/goHTTP/internal/status"
)

type ResponseWriter interface {
	Header() Header
	Write(b []byte) (int, error)
	WriteHeader(code status.Code)
}

type Response struct {
	request     *Request
	version     string
	statuscode  status.Code
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
		response.WriteHeader(status.Ok)
	}
	return response.conn.Write(b)
}

func (r *Response) WriteHeader(code status.Code) {
	r.statuscode = code
	r.status = "ok"
	r.wroteCode = true

	r.conn.Write([]byte(fmt.Sprintf("%s %d %s\r\n", r.version, r.statuscode, r.status)))
	r.conn.Write([]byte(r.header.String()))
}
