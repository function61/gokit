package jsonfile

// I got tired of writing the same readConfigFile(), writeConfigFile() magic over and over again..

import (
	"encoding/json"
	"os"
)

func Read(path string, content interface{}, disallowUnknownFields bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)
	if disallowUnknownFields {
		jsonDecoder.DisallowUnknownFields()
	}
	if err := jsonDecoder.Decode(content); err != nil {
		return err
	}

	return nil
}

func Write(path string, content interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonEncoder := json.NewEncoder(file)
	jsonEncoder.SetIndent("", "    ")
	return jsonEncoder.Encode(content)
}
