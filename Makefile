test:
	go test -v ./cmd/... ./internal/...

build:
	go clean
	go build -o csv-converter ./cmd/
