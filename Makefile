run:
	go run cmd/gendiff/main.go $(ARGS)

build:
	go build -o bin/gendiff cmd/gendiff/main.go

test:
	go test ./...