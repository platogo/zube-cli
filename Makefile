build:
	go build -o zb .

run:
	go run main.go

test:
	go test

all: test build