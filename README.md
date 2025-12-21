# Pure Go Web Calculator

## Idea

Use Pure Go (standard libraries only), to write a static site that compares
various JS implementations to Go+wasm implementations, in terms of resources,
and with code rendered side by side.

#### Dependencies

- Go
- typescript
- esbuild
- npm

### Dev Guide

`make serve` does it all. Compiles the Go project, installs (the few) npm
dependencies, builds the JS/TS/CSS bundles, generates the static site from templates,
bundle, and assets, and serves the site at `localhost:8080`

### Features

- [x] bundled typescript
- [x] bundled css
- [ ] favicon
- [ ] GH Pages
- [ ] GH Action
- [ ] GH Page custom domain
- [ ] watch command 1 - watch for css and ts file changes
- [ ] watch command 2 - watch for go file changes
- [ ] WASM module
- [ ] Calc.Add
