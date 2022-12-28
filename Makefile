all: lint test

build:
	go build -v ./...

test:
	go test -cover  -v ./...

cover:
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

lint:
	golangci-lint run

bench:
	go test -benchmem -count 100 -bench .

mod:
	go mod tidy && go mod vendor
