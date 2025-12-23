.DEFAULT_GOAL := generate

.PHONY: fmt vet build wasm generate watch serve

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
wasm:
	$(MAKE) -C wasm build
generate: wasm build
	./go-wasm-js-comp generate && cp ./wasm/wasm_exec.js out/ && cp ./wasm/main.wasm out/
watch: build
	while true; do \
		./go-wasm-js-comp generate; \
		fswatch -1 ./ts ./templates ./stylesheets || break; \
    done
serve: build
	./go-wasm-js-comp serve
