package jsonfile

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestRead(t *testing.T) {
	assert.Equal(t, ReadDisallowUnknownFields("notfound.json", &struct{}{}).Error(), "open notfound.json: no such file or directory")
	assert.Equal(t, ReadDisallowUnknownFields("testdata/example.json", &struct{}{}).Error(), `testdata/example.json: decode failed: json: unknown field "foo"`)
}
