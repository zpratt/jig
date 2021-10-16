all: test build build-windows test-e2e

build:
	go build -a

build-windows:
	GOOS=windows GOARCH=amd64 go build -a

clean:
	go clean -cache
	go clean

test: clean
	go test -v ./...

test-e2e: clean build
	brew uninstall httping || echo skip uninstalling httping because it is not installed
	./jig
	httping -v
