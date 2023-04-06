package hcl2json

import (
	"bytes"
	"io"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestConvert(t *testing.T) {
	demo := `
person {
	name = "Foo"
	age = 32
	title = "CEO"
}

person {
	name = "Bar"
	age = 28
	title = "Software engineer"
}
`
	jsonBytesReader, err := Convert(bytes.NewBufferString(demo))
	assert.Ok(t, err)

	jsonBytes, err := io.ReadAll(jsonBytesReader)
	assert.Ok(t, err)

	assert.Equal(t, string(jsonBytes), `{
  "person": [
    {
      "age": 32,
      "name": "Foo",
      "title": "CEO"
    },
    {
      "age": 28,
      "name": "Bar",
      "title": "Software engineer"
    }
  ]
}`)
}
