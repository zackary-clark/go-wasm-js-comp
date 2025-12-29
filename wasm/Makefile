.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt:
	GOOS=js GOARCH=wasm go fmt ./...
vet: fmt
	GOOS=js GOARCH=wasm go vet ./...
build: vet
	GOOS=js GOARCH=wasm go build -o bundle.wasm

