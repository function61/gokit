package logex

import (
	"io/ioutil"
	"log"
)

// a logger that discards its output
var Discard = log.New(ioutil.Discard, "", 0)

// helper for libraries in getting an alwas non-nil logger. if a logger given to you was
// nil, you get a Discard logger back, which you can call safely - it doesn't write anywhere
func NonNil(ref *log.Logger) *log.Logger {
	if ref == nil {
		return Discard
	}

	return ref
}
