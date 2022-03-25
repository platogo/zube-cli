BINARY_NAME=zube

build:
	go build -o bin/${BINARY_NAME} .

compile:
	@echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 .
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64 .

clean:
	go clean
	rm ./bin/*

run:
	go run main.go

test:
	go test

all: test build