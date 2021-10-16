all: test build build-windows

build:
	go build -a

build-windows:
	GOOS=windows GOARCH=amd64 go build -a

clean:
	go clean -cache
	rm -f {jig,jig.exe}

test: clean
	go test -v ./...

test-e2e: clean build
	brew uninstall hping || echo skip uninstalling hping because it is not installed
	./jig
	hping -v
