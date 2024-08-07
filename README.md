# HTTP Server in Golang

## Usage

You can find examples in [examples](./examples/)

```go
func main() {
    server := goHTTP.NewServer("127.0.0.1:8000")
    server.HandleFunc("/hello", func(w goHTTP.ResponseWriter, r *goHTTP.Request) {
        w.Write([]byte("Hello World"))
    })
}
```

This will write "Hello World" to the connection

## Idea

Server
    - Running on a port
    - Receiving and sending tcp requests
    - Parsing HTTP message
    - Handlers for the different paths

## TODO

Server
- [x] Sending and receiving requests
- [x] Parsing http message
- [x] Send back some files via http
- [x] Adding handlers for different paths
- [ ] Add middleware
