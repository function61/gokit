// I got tired of writing the same readConfigFile(), writeConfigFile() magic over and over again..
package jsonfile

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/function61/gokit/os/osutil"
)

func Read(path string, data interface{}, disallowUnknownFields bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := Unmarshal(file, data, disallowUnknownFields); err != nil {
		return fmt.Errorf("%s: %s", path, err.Error())
	}

	return nil
}

func Write(path string, data interface{}) error {
	return osutil.WriteFileAtomic(path, func(writer io.Writer) error {
		return Marshal(writer, data)
	})
}

func Unmarshal(source io.Reader, data interface{}, disallowUnknownFields bool) error {
	jsonDecoder := json.NewDecoder(source)
	if disallowUnknownFields {
		jsonDecoder.DisallowUnknownFields()
	}
	if err := jsonDecoder.Decode(data); err != nil {
		// sadly, line numbers are only possible with hacks (requiring buffering):
		//   https://github.com/hashicorp/packer/blob/master/common/json/unmarshal.go
		return fmt.Errorf("JSON parsing failed: %s", err.Error())
	}

	return nil
}

func Marshal(destination io.Writer, data interface{}) error {
	jsonEncoder := json.NewEncoder(destination)
	jsonEncoder.SetIndent("", "    ")
	return jsonEncoder.Encode(data)
}
