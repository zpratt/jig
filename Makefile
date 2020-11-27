build: clean
	go build -a

build-windows: clean
	GOOS=windows GOARCH=amd64 go build -a

clean:
	rm -f {jig,jig.exe}
