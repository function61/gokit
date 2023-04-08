package mime

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestTypeByExtension(t *testing.T) {
	assert.Equal(t, TypeByExtension(".json", NoFallback), "application/json")
	assert.Equal(t, TypeByExtension("json", NoFallback), "application/json")
	assert.Equal(t, TypeByExtension("JSON", NoFallback), "application/json")
	assert.Equal(t, TypeByExtension("JsOn", NoFallback), "application/json")

	assert.Equal(t, TypeByExtension("mkv", NoFallback), "video/x-matroska")
	assert.Equal(t, TypeByExtension("mkv", OctetStream), "video/x-matroska")

	assert.Equal(t, TypeByExtension("doesnotexist", NoFallback), "")
	assert.Equal(t, TypeByExtension("doesnotexist", OctetStream), "application/octet-stream")
}

func TestExtensionByType(t *testing.T) {
	assert.Equal(t, ExtensionByType("image/jpeg", "bin"), "jpg")
	assert.Equal(t, ExtensionByType("video/x-matroska", "bin"), "mkv")
	assert.Equal(t, ExtensionByType("application/json", "bin"), "json")

	assert.Equal(t, ExtensionByType("dunno/notfound", "bin"), "bin")
	assert.Equal(t, ExtensionByType("", "bin"), "bin")
	assert.Equal(t, ExtensionByType("", NoFallback), "")
}

func TestIs(t *testing.T) {
	assert.Equal(t, Is("image/jpeg", TypeImage), true)
	assert.Equal(t, Is("image/", TypeImage), true)
	assert.Equal(t, Is("image", TypeImage), false)
	assert.Equal(t, Is("text/plain", TypeImage), false)
}
