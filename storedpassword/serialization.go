package storedpassword

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

// format: <strategyId>:<salt>:<derived>
type StoredPassword string

func serialize(strategyId string, salt []byte, derived []byte) StoredPassword {
	serialized := fmt.Sprintf(
		"%s:%s:%s",
		strategyId,
		base64.RawURLEncoding.EncodeToString(salt),
		base64.RawURLEncoding.EncodeToString(derived))

	return StoredPassword(serialized)
}

func deserialize(stored StoredPassword) (string, []byte, []byte, error) {
	storedParts := strings.Split(string(stored), ":")

	if len(storedParts) != 3 {
		return "", nil, nil, errors.New("invalid structure of StoredPassword")
	}

	storedStrategyId := storedParts[0]
	storedSalt, err := base64.RawURLEncoding.DecodeString(storedParts[1])
	if err != nil {
		return "", nil, nil, err
	}
	storedBytes, err := base64.RawURLEncoding.DecodeString(storedParts[2])
	if err != nil {
		return "", nil, nil, err
	}

	return storedStrategyId, storedSalt, storedBytes, nil
}
