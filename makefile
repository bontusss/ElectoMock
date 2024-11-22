.PHONY: run build test docker-build docker-up deps clean

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Binary Name
BINARY_NAME=electomock

#Build parameters
BUILD_DIR=build
MAIN_PKG=cmd/server/main.go

#DEV commands
run:
	air

build:
	$(GOBUILD) -o $(BUild_DIR)/$(BINARY_NAME) $(MAIN_PKG)

test:
	$(GOTEST) ./... -v

#Database Migration
migrate-up:
	migrate -path