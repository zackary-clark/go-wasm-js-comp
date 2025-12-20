.DEFAULT_GOAL := build

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
generate: build
	./go-wasm-calc generate
serve: generate
	./go-wasm-calc serve
