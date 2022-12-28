build:
	go build -v ./...

test:
	go test -cover  -v ./...

lint:
	golangci-lint run

bench:
	go test -benchmem -count 100 -bench .

mod:
	go mod tidy && go mod vendor
