run:
	go run cmd/gendiff/main.go $(ARGS)
json:
	go run cmd/gendiff/main.go testdata/fixtures/file1.json testdata/fixtures/file2.json
yml:
	go run cmd/gendiff/main.go testdata/fixtures/file1.yml testdata/fixtures/file2.yaml

build:
	go build -o bin/gendiff cmd/gendiff/main.go

test:
	go test ./...