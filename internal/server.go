package internal

import (
	"fmt"
	"net"
)

type HttpServer struct {
	Address string
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

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		handleError(err)
		return
	}

	fmt.Printf("Received %d bytes:\n%s\n", n, buf[:n])

	m, _, _ := getMethod(buf[:n])
	fmt.Printf("Received method %s\n", m)

	conn.Write([]byte("Hello World"))
}

func handleError(_ error) {

}
