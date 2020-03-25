package hcl2json

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/function61/gokit/assert"
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
	assert.Assert(t, err == nil)

	jsonBytes, err := ioutil.ReadAll(jsonBytesReader)
	assert.Assert(t, err == nil)

	assert.EqualString(t, string(jsonBytes), `{
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
