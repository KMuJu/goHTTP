package main

import "github.com/kmuju/goHTTP/internal"

func main() {
	server := internal.NewServer("127.0.0.1:8000")
	server.HandleFunc("/hello", func(w internal.ResponseWriter, r *internal.Request) {
		w.Write([]byte("Hello World"))
	})
	server.HandleFunc("/", func(w internal.ResponseWriter, r *internal.Request) {
		w.Write([]byte(`<!DOCTYPE html>

<html lang="en">

<head>

  <meta charset="UTF-8" />

  <title>Hello, world!</title>

  <meta name="viewport" content="width=device-width,initial-scale=1" />

  <meta name="description" content="" />

  <link rel="icon" href="favicon.png">

</head>

<body>

  <h1>Hello, world!</h1>

</body>

</html>`))
	})
	server.Run()
}
