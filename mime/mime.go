// MIME (Content Type) database. The stdlib MIME package relies on OS to have the list, so
// the results are inconsistent (f.ex. on stock Alpine Linux it doesn't even recognize JSON)
package mime

import (
	gomime "github.com/cubewise-code/go-mime"
)

const (
	OctetStream = "application/octet-stream"
	NoFallback  = ""
)

// Need to add on top of go-mime because it doesn't support fallback
// Supports:
// - ".json"
// - "json"
// - "JSON"
// - "JsOn"
func TypeByExtension(ext string, fallback string) string {
	if contentType := gomime.TypeByExtension(ext); contentType != "" {
		return contentType
	} else {
		return fallback
	}
}
