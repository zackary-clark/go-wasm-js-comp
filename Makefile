.DEFAULT_GOAL := generate

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
generate: build
	./go-wasm-js-comp generate
serve: generate
	./go-wasm-js-comp serve
