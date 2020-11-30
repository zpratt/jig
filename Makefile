all: test build build-windows

build:
	go build -a

build-windows:
	GOOS=windows GOARCH=amd64 go build -a

clean:
	rm -f {jig,jig.exe}

test: clean
	go test ./...
