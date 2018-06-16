# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOLIST=$(GOCMD) list
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet


all: unit build


.PHONY: unit
unit: ## @testing Run the unit tests 
	$(GOFMT) ./...
	$(GOVET) ./smgp/...
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic $(shell go list ./smgp/...)

.PHONY: build
build:
	$(GOBUILD) -o ./bin/transmitter ./cmd/transmitter
	$(GOBUILD) -o ./bin/receiver ./cmd/receiver
	$(GOBUILD) -o ./bin/mockserver ./cmd/mockserver


.PHONY: build_linux
build_linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o ./bin/transmitter ./cmd/transmitter
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o ./bin/receiver ./cmd/receiver
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o ./bin/mockserver ./cmd/mockserver

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf ./bin/ coverage.txt
