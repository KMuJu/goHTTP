build:
	@go build -o goHTTP cmd/main.go

run:
	@go build -o goHTTP cmd/main.go
	./goHTTP
