build:
	go build -o bin/zube .

compile:
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/zube-darwin-amd64 .
	GOOS=linux GOARCH=386 go build -o bin/zube-linux-386 .
	GOOS=windows GOARCH=amd64 go build -o bin/zube-windows-amd64 .

run:
	go run main.go

test:
	go test

all: test build