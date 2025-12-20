# Pure Go Web Calculator

## Idea

Use Pure Go (standard libraries only), to write a static site that compares
various JS implementations to Go+wasm implementations, in terms of resources,
and with code rendered side by side.

### Dev Guide

`make serve` does it all. Compiles the Go project, installs (the few) npm
dependencies, builds the JS/TS bundle, generates the static site from templates,
bundle, and assets, and serves the site at `localhost:8080`
