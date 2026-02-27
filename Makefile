lint:
	golangci-lint run ./...
	@echo "ok"

test:
	go test ./src/...

build:
	go build -o bin/tiktaktoe ./src