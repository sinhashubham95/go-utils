
BIN=go

build:
	${BIN} build -v ./...

test:
	go test -race -v ./...

bench:
	go test -benchmem -count 3 -bench ./...

coverage:
	${BIN} test -v -coverprofile=cover.out -covermode=atomic ./...
	${BIN} tool cover -html=cover.out -o cover.html
