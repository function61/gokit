package mime

import (
	"testing"

	"github.com/function61/gokit/assert"
)

func TestTypeByExtension(t *testing.T) {
	assert.EqualString(t, TypeByExtension(".json", NoFallback), "application/json")
	assert.EqualString(t, TypeByExtension("json", NoFallback), "application/json")
	assert.EqualString(t, TypeByExtension("JSON", NoFallback), "application/json")
	assert.EqualString(t, TypeByExtension("JsOn", NoFallback), "application/json")

	assert.EqualString(t, TypeByExtension("mkv", NoFallback), "video/x-matroska")
	assert.EqualString(t, TypeByExtension("mkv", OctetStream), "video/x-matroska")

	assert.EqualString(t, TypeByExtension("doesnotexist", NoFallback), "")
	assert.EqualString(t, TypeByExtension("doesnotexist", OctetStream), "application/octet-stream")
}
