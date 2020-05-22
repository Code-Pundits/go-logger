#Go parameters
GOCMD=go
GOTIDY=$(GOCMD) mod tidy
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOBUILD=$(GOCMD) build
BINARY_NAME=auth-service
BINARY_UNIX=$(BINARY_NAME)_unix
CMD_PATH=cmd/api/main.go
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
CONFIG_PATH=cmd/api/config.json

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

tidy: 
	$(GOTIDY)

test:
	$(GOTEST) -v ./...

vet:
	$(GOVET) ./...

build:
	echo APP_VERSION=${VERSION} > .env
	$(GOBUILD) -o $(BINARY_NAME) -v $(CMD_PATH)

run: build 
	./$(BINARY_NAME)