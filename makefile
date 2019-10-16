GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test
test: 
	$(GOTEST) -v ./...
