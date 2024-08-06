package main

import "github.com/kmuju/goHTTP/internal"

func main() {
	server := internal.HttpServer{Address: "127.0.0.1:8000"}
	server.Run()
}
