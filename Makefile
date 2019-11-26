# TODO: add versioning
# TODO: add ldflags for version of sdk

all: tools install

mod:
	@go mod tidy

tools:
	@go get -u github.com/go-bindata/go-bindata/...

generate: 
	@go-bindata --pkg cmd -o cmd/bindata.go -prefix "templates/" templates/...

build: generate mod
	@go build -mod=readonly -o build/scaffold main.go

install: generate mod
	@go install -mod=readonly ./...

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	go mod verify

.PHONY: all tools generate build install