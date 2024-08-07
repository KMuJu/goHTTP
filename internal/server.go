package internal

import (
	"fmt"
	"net"
)

type HttpServer struct {
	Address  string
	routes   []*Route
	routemap map[string]*Route
}

func NewServer(address string) HttpServer {
	return HttpServer{
		Address:  address,
		routes:   []*Route{},
		routemap: make(map[string]*Route),
	}
}

type Route struct {
	name    string
	handler Handler
}

func (server *HttpServer) Run() error {
	ln, err := net.Listen("tcp", server.Address)
	if err != nil {
		return err
	}

	fmt.Printf("Listening for packets on %s\n", server.Address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go server.handleConnection(conn)
	}

}

func (server *HttpServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		handleError(err)
		return
	}

	request, err := parseRequest(buf[:n])
	if err != nil {
		handleError(err)
		return
	}
	fmt.Printf("Request: %s %s\n", request.Method, request.URL)

	route, ok := server.routemap[request.URL]
	if !ok {
		handleError(notFound)
		return
	}
	route.handler.ServeHTTP(newResponseWriter(conn, &request), &request)
}

func (server *HttpServer) HandleFunc(pattern string, handler HandlerFunc) {
	route := &Route{name: pattern, handler: handler}
	server.routes = append(server.routes, route)
	server.routemap[pattern] = route
}

func handleError(_ error) {
	fmt.Printf("Error in request\n")
}
