package main

import (
	"os"

	"github.com/kmuju/goHTTP"
)

func main() {
	server := goHTTP.NewServer("127.0.0.1:8000")
	server.HandleFunc("/hello", func(w goHTTP.ResponseWriter, r *goHTTP.Request) {
		w.Write([]byte("Hello World"))
	})
	server.HandleFunc("/", func(w goHTTP.ResponseWriter, r *goHTTP.Request) {
		w.Header()["Content-Type"] = []string{"text/html; charset=utf-8"}
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
	server.HandleFunc("/img", func(w goHTTP.ResponseWriter, r *goHTTP.Request) {
		fileName := "README.md"
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		w.Write(data)
	})
	server.Run()
}
