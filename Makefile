BINARY_NAME=zube

PREFIX := /usr/local

all: test build

build:
	go build -o bin/${BINARY_NAME} .

compile:
	@echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o bin/${BINARY_NAME}-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64.exe .

clean:
	go clean
	rm ./bin/*

run:
	go run main.go

test:
	go test -v ./zube ./utils

format:
	@echo "Formatting the entire project"
	go fmt

install: ./bin/$(BINARY_NAME)
	install ./bin/$(BINARY_NAME) $(PREFIX)/bin/
