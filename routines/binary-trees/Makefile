.DEFAULT_GOAL:=build

.PHONY: fmt vet clean build

fmt:
	go fmt .

vet: fmt
	go vet .

clean: vet
	go clean

build: clean
	go mod tidy
	go build -o binary-trees binary-trees.go
