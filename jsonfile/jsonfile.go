package jsonfile

// I got tired of writing the same readConfigFile(), writeConfigFile() magic over and over again..

import (
	"encoding/json"
	"io"
	"os"
)

func Read(path string, data interface{}, disallowUnknownFields bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return Unmarshal(file, data, disallowUnknownFields)
}

func Write(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return Marshal(file, data)
}

func Unmarshal(source io.Reader, data interface{}, disallowUnknownFields bool) error {
	jsonDecoder := json.NewDecoder(source)
	if disallowUnknownFields {
		jsonDecoder.DisallowUnknownFields()
	}
	if err := jsonDecoder.Decode(data); err != nil {
		return err
	}

	return nil
}

func Marshal(destination io.Writer, data interface{}) error {
	jsonEncoder := json.NewEncoder(destination)
	jsonEncoder.SetIndent("", "    ")
	return jsonEncoder.Encode(data)
}
