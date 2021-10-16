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
	brew uninstall hping || echo skip uninstalling hping because it is not installed
	./jig
	hping -v
