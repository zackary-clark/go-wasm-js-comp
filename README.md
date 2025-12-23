# Pure Go Site Comparing WASM to JS Implementations

## Idea

Use Pure Go (standard libraries only), to write a static site that compares
various JS implementations to Go+wasm implementations, in terms of resources,
and with code rendered side by side.

#### Dependencies

- Go
- typescript
- esbuild
- npm

- fswatch (for `make watch` only)

### Dev Guide

#### Local Dev Server

`make watch` builds the go app, generates the static site, and watches the
`ts`, `templates`, and `stylesheets` directories to regenerate the site

`make serve` serves the site (generated with `make generate` or `make watch`) at `localhost:8080`

#### Deployment

Pushing to `main` runs the deploy action, which builds the go app, generates
the site, and publishes to GH Pages.

### Features

- [x] bundled typescript
- [x] bundled css
- [x] favicon
- [x] GH Pages
- [x] GH Action
- [x] GH Page custom domain
- [x] watch command 1 - watch for css and ts file changes
- [ ] WASM module
- [ ] watch command 2 - watch for go file changes
- [ ] Calc.Add
