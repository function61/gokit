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
	// use temp-file scheme to write file atomically (as much as POSIX allows)
	tempName := path + ".temp"

	file, err := os.Create(tempName)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := Marshal(file, data); err != nil {
		return err
	}

	if err := file.Close(); err != nil { // double close is intentional
		return err
	}

	return os.Rename(tempName, path)
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
