package jsonfile

import (
	"strings"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestRead(t *testing.T) {
	assert.Equal(t, ReadDisallowUnknownFields("notfound.json", &struct{}{}).Error(), "open notfound.json: no such file or directory")
	assert.Equal(t, ReadDisallowUnknownFields("testdata/example.json", &struct{}{}).Error(), `testdata/example.json: decode failed: json: unknown field "foo"`)
}

func TestUnmarshalTrailer(t *testing.T) {
	assert.Equal(t,
		UnmarshalDisallowUnknownFields(strings.NewReader(`{"Name": "Joonas"}     invalidtrailer`), &struct{ Name string }{}).Error(),
		"validate JSON trailer: expecting EOF got invalid character 'i' looking for beginning of value")

	// whitespace after data is ok
	assert.Ok(t,
		UnmarshalDisallowUnknownFields(strings.NewReader("{\"Name\": \"Joonas\"}  \n\t  "), &struct{ Name string }{}))
}
