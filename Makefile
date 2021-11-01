all: test build build-darwin build-windows build-linux test-e2e test-release

BREW_BIN ?= brew
GO_BIN ?= go # in case you want to run some other version / containerize this run

NAME=jig
TEST_PACKAGE=bat
#BASE_DIR=$(shell pwd) # useful to have in case the cwd is needed

.PHONY: build
build: 
	$(GO_BIN) build -a -v .

.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -a

.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -a

.PHONY: build-linux
build-windows:
	GOOS=linux GOARCH=amd64 go build -a

.PHONY: build-all
build-all: build build-darwin build-windows build-linux

.PHONY: test-release
test-release: goreleaser build --snapshot --rm-dist

.PHONY: clean
clean:
	$(GO_BIN) clean -cache
	$(GO_BIN) clean

.PHONY: test
test: clean
	$(GO_BIN) test -v -cover ./...

.PHONY: test-e2e
test-e2e: clean build
	$(BREW_BIN) uninstall $(TEST_PACKAGE) || echo skip uninstalling hping because it is not installed
	./jig
	$(TEST_PACKAGE) -V
