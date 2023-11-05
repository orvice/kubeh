.PHONY: build lint build-release
build:
	go build -o ${GOBIN}/kubeh ./cmd/kubeh/main.go
lint:
	golangci-lint run

build-release:
	go build -o bin/kubeh ./cmd/kubeh/main.go