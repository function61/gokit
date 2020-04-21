![Build](https://github.com/function61/gokit/workflows/Build/badge.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/function61/gokit)

Frequently needed, stupid simple, packages in Go.


Deprecations
------------

- `stopper/` - just use the stdlib `context` for cancellation
  * if you need to wait for single "stopped" just wait for fn to return
  * you could use `taskrunner/` for multiple return waits
