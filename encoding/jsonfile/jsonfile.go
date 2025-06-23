// I got tired of writing the same readConfigFile(), writeConfigFile() boilerplate over and over again..
package jsonfile

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/function61/gokit/os/osutil"
)

func ReadDisallowUnknownFields(path string, data interface{}) error {
	return read(path, data, UnmarshalDisallowUnknownFields)
}

func ReadAllowUnknownFields(path string, data interface{}) error {
	return read(path, data, UnmarshalAllowUnknownFields)
}

func read(path string, data interface{}, unmarshal func(io.Reader, interface{}) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := unmarshal(file, data); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	return nil
}

func Write(path string, data interface{}, options ...osutil.WriteFileOption) error {
	return osutil.WriteFileAtomic(path, func(writer io.Writer) error {
		return Marshal(writer, data)
	}, options...)
}

func UnmarshalDisallowUnknownFields(source io.Reader, data interface{}) error {
	return unmarshalInternal(source, data, true)
}

func UnmarshalAllowUnknownFields(source io.Reader, data interface{}) error {
	return unmarshalInternal(source, data, false)
}

func unmarshalInternal(source io.Reader, data interface{}, disallowUnknownFields bool) error {
	jsonDecoder := json.NewDecoder(source)
	if disallowUnknownFields {
		jsonDecoder.DisallowUnknownFields()
	}
	if err := jsonDecoder.Decode(data); err != nil {
		// sadly, line numbers are only possible with hacks (requiring buffering):
		//   https://github.com/hashicorp/packer/blob/master/common/json/unmarshal.go
		//
		// no need mention JSON as context in error message, because `err` already mentions that.
		return fmt.Errorf("decode failed: %s", err.Error())
	}

	// validate there's no (non-whitespace) trailers after the decoded data to prevent smuggling undefined data
	// after the payload. https://blog.trailofbits.com/2025/06/17/unexpected-security-footguns-in-gos-parsers/
	if err := jsonDecoder.Decode(data); err != io.EOF {
		return fmt.Errorf("validate JSON trailer: expecting EOF got %v", err)
	}

	return nil
}

func Marshal(destination io.Writer, data interface{}) error {
	jsonEncoder := json.NewEncoder(destination)
	jsonEncoder.SetIndent("", "    ")
	return jsonEncoder.Encode(data)
}
