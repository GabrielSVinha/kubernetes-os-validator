GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
PROJECT_ROOT=$(GOPATH)/src/github.com/GabrielSVinha/kubernetes-os-validator
FIXTURES_DIR=$(PROJECT_ROOT)/fixtures
OUTPUT=$(PROJECT_ROOT)/bin/validate

all: test build

test:
	GOCACHE=off FIXTURES=$(FIXTURES_DIR) $(GOTEST) -v ./...

build:
	$(GOBUILD) -o $(OUTPUT)
