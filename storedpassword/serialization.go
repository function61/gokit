package storedpassword

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

// format: $<strategyId>$<salt>$<derived>
// trying to be somewhat compatible with
// 	https://passlib.readthedocs.io/en/stable/modular_crypt_format.html#application-defined-hashes
// with the exception that we're storing the cost in the <strategyId> so that we don't have
// to implement different parsing formats per strategy
type StoredPassword string

func serialize(strategyId string, salt []byte, derived []byte) StoredPassword {
	serialized := fmt.Sprintf(
		"$%s$%s$%s",
		strategyId,
		base64.RawURLEncoding.EncodeToString(salt),
		base64.RawURLEncoding.EncodeToString(derived))

	return StoredPassword(serialized)
}

func deserialize(stored StoredPassword) (string, []byte, []byte, error) {
	storedParts := strings.Split(string(stored), "$")

	if len(storedParts) != 4 || storedParts[0] != "" {
		return "", nil, nil, errors.New("invalid structure of StoredPassword")
	}

	storedStrategyId := storedParts[1]
	storedSalt, err := base64.RawURLEncoding.DecodeString(storedParts[2])
	if err != nil {
		return "", nil, nil, err
	}
	storedBytes, err := base64.RawURLEncoding.DecodeString(storedParts[3])
	if err != nil {
		return "", nil, nil, err
	}

	return storedStrategyId, storedSalt, storedBytes, nil
}
