GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_DIR=./bin
BINARY_NAME=mathch
    
all: test build
build: 
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v ./...
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm $(BINARY_DIR)/*
run: build
	$(BINARY_DIR)/$(BINARY_NAME)
