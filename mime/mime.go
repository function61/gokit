// MIME (Content Type) database. The stdlib MIME package relies on OS to have the list, so
// the results are inconsistent (f.ex. on stock Alpine Linux it doesn't even recognize JSON)
package mime

import (
	"strings"
)

const (
	OctetStream = "application/octet-stream"
	NoFallback  = ""
)

type Type string

// taken from https://en.wikipedia.org/wiki/Media_type
const (
	TypeApplication Type = "application"
	TypeAudio       Type = "audio"
	TypeExample     Type = "example"
	TypeFont        Type = "font"
	TypeImage       Type = "image"
	TypeMessage     Type = "message"
	TypeModel       Type = "model"
	TypeMultipart   Type = "multipart"
	TypeText        Type = "text"
	TypeVideo       Type = "video"
)

// JSON tags are defined due to importing in code generation phase,
// but please do not rely on those (don't JSON-marshal this spec)
type Spec struct {
	Extensions   []string `json:"extensions"`
	Compressible *bool    `json:"compressible"` // nil if unknown
	CharEncoding string   `json:"charset"`      // not always recorded
	Source       string   `json:"source"`
}

// jpg => image/jpeg
var extLookup = func() map[string]string {
	lookup := map[string]string{}

	for contentType, spec := range mimeTypes {
		for _, ext := range spec.Extensions {
			lookup[ext] = contentType
		}
	}

	return lookup
}()

// Supports:
// - ".json"
// - "json"
// - "JSON"
// - "JsOn"
func TypeByExtension(ext string, fallback string) string {
	// .JSON => "json"
	extNormalized := strings.ToLower(strings.TrimPrefix(ext, "."))

	if contentType, found := extLookup[extNormalized]; found {
		return contentType
	} else {
		return fallback
	}
}

// image/jpeg => jpg
func ExtensionByType(contentType string, fallback string) string {
	// some overrides
	switch contentType {
	case "image/jpeg":
		return "jpg" // first extension in DB is jpeg which is less universally used
	}

	spec, found := mimeTypes[contentType]
	if !found || len(spec.Extensions) == 0 {
		return fallback
	}

	return spec.Extensions[0]
}

// Is("image/jpeg", TypeImage) => true
// Is("text/plain", TypeImage) => false
func Is(contentType string, typ Type) bool {
	return strings.HasPrefix(contentType, string(typ)+"/")
}
