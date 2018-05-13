run: build
	./csv-converter

test:
	go test -v ./cmd/... ./internal/...

build:
	go clean
	go build -o csv-converter ./cmd/

install: build
	go install ./cmd
