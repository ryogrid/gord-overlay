GOCMD=go
GOPATH?=$(shell if test -x `which go`; then go env GOPATH; else echo "$(HOME)/go"; fi)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gordctl
BUILD_TARGET=main.go

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) $(BUILD_TARGET)
test:
	$(GOTEST) -v ./...