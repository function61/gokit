![Build](https://github.com/function61/gokit/workflows/Build/badge.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/function61/gokit)

Frequently needed, stupid simple, packages in Go.

Directory structure follows [Go's stdlib](https://pkg.go.dev/std?tab=packages) where there
are equivalents. E.g. `httputils` is found in [net/http/httputils](net/http/httputils/).


Deprecations
------------

- `csrf/` - just use SameSite cookies
