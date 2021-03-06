package mime

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
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

func TestExtensionByType(t *testing.T) {
	assert.EqualString(t, ExtensionByType("image/jpeg", "bin"), "jpg")
	assert.EqualString(t, ExtensionByType("video/x-matroska", "bin"), "mkv")
	assert.EqualString(t, ExtensionByType("application/json", "bin"), "json")

	assert.EqualString(t, ExtensionByType("dunno/notfound", "bin"), "bin")
	assert.EqualString(t, ExtensionByType("", "bin"), "bin")
	assert.EqualString(t, ExtensionByType("", NoFallback), "")
}

func TestIs(t *testing.T) {
	assert.Assert(t, Is("image/jpeg", TypeImage))
	assert.Assert(t, Is("image/", TypeImage))
	assert.Assert(t, !Is("image", TypeImage))
	assert.Assert(t, !Is("text/plain", TypeImage))
}
