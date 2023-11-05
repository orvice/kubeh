.PHONY: build lint build-release
build:
	go build -o ${GOBIN}/kubeh ./cmd/kubeh/main.go
lint:
	golangci-lint run -c=./.golangci.yml --out-format=line-number ./... --timeout 3m

build-release:
	go build -o bin/kubeh ./cmd/kubeh/main.go