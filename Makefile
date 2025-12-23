.DEFAULT_GOAL := generate

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
generate: build
	./go-wasm-js-comp generate
watch: build
	while true; do \
		./go-wasm-js-comp generate; \
		fswatch -1 ./ts ./templates ./stylesheets || break; \
    done
serve: build
	./go-wasm-js-comp serve
