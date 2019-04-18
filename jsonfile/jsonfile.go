// I got tired of writing the same readConfigFile(), writeConfigFile() magic over and over again..
package jsonfile

import (
	"encoding/json"
	"fmt"
	"github.com/function61/gokit/atomicfilewrite"
	"io"
	"os"
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
	return atomicfilewrite.Write(path, func(writer io.Writer) error {
		return Marshal(writer, data)
	})
}

func Unmarshal(source io.Reader, data interface{}, disallowUnknownFields bool) error {
	jsonDecoder := json.NewDecoder(source)
	if disallowUnknownFields {
		jsonDecoder.DisallowUnknownFields()
	}
	if err := jsonDecoder.Decode(data); err != nil {
		return fmt.Errorf("JSON parsing failed: %s", err.Error())
	}

	return nil
}

func Marshal(destination io.Writer, data interface{}) error {
	jsonEncoder := json.NewEncoder(destination)
	jsonEncoder.SetIndent("", "    ")
	return jsonEncoder.Encode(data)
}
