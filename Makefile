build:
	@go build -o bin/simple examples/simple/main.go

run:
	@go build -o bin/simple examples/simple/main.go
	./bin/simple
