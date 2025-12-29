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
	./go-wasm-js-comp generate && cp ./wasm/wasm_exec.js out/ && cp ./wasm/bundle.wasm out/
watch: wasm build
	while true; do \
		$(MAKE) generate; \
		fswatch -1 ./ts ./templates ./stylesheets ./wasm ./commands || break; \
    done
serve: build
	./go-wasm-js-comp serve
