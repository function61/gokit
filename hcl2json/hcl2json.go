package hcl2json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl"
	"io"
	"io/ioutil"
)

func Convert(hclContent io.Reader) (io.Reader, error) {
	hclBuffer, err := ioutil.ReadAll(hclContent)
	if err != nil {
		return nil, err
	}

	// read & parse HCL to generic struct
	var v interface{}
	errHcl := hcl.Unmarshal(hclBuffer, &v)
	if errHcl != nil {
		return nil, fmt.Errorf("unable to parse HCL: %s", errHcl)
	}

	// re-encode the generic struct to JSON
	asJson, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(asJson), nil
}
