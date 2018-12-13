package storedpassword

import (
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"io"
)

type DerivationStrategy interface {
	Id() string
	Derive(plaintext []byte, salt []byte) []byte
}

// transforms a password into a form that is safe to store in a database, provided
// that you pass CurrentBestDerivationStrategy as the strategy.
func Store(plaintext string, strategy DerivationStrategy) (StoredPassword, error) {
	// generate salt, 256 bits ought to be enough since salt is not a secret per se
	// https://stackoverflow.com/questions/9619727
	salt := [32]byte{}
	if _, err := io.ReadFull(rand.Reader, salt[:]); err != nil {
		return "", err
	}

	derived := strategy.Derive([]byte(plaintext), salt[:])

	return serialize(strategy.Id(), salt[:], derived), nil
}

// returns true if given password matches stored password, false if not
// return error if there's an internal error in the checking process (e.g. unknown derivation strategy)
//
// this is safe from timing attacks
func Verify(stored StoredPassword, givenPlaintext string, strategyResolver func(string) DerivationStrategy) (bool, error) {
	storedStrategyId, storedSalt, storedBytes, err := deserialize(stored)
	if err != nil {
		return false, err
	}

	strategy := strategyResolver(storedStrategyId)
	if strategy == nil {
		return false, errors.New("unknown strategy")
	}

	givenBytes := strategy.Derive([]byte(givenPlaintext), storedSalt)

	if subtle.ConstantTimeCompare(storedBytes, givenBytes) == 1 {
		return true, nil
	}

	return false, nil
}
