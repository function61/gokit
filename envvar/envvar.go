// Small wrappers for dealing with ENV vars
package envvar

import (
	"encoding/base64"
	"errors"
	"os"
)

func Get(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("ENV not defined: " + key)
	}

	return value, nil
}

func GetFromBase64Encoded(key string) ([]byte, error) {
	dataBase64, err := Get(key)
	if err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		return nil, err
	}

	return data, nil
}
