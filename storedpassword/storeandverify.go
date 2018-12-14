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

// both returns MUST NOT refer to same derivation strategy, so upgrade is nil if stored
// password is already succicient
type StrategyResolver func(id string) (found DerivationStrategy, upgrade DerivationStrategy)

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

// 1st return: true if given password matches stored password, false if not
// 2nd return: if != "" is the upgraded version of the stored password, if upgraded DerivationStrategy found
// 3rd return: error if there's an internal error in the checking process (e.g. unknown derivation strategy)
//
// this function is safe from timing attacks
func Verify(stored StoredPassword, givenPlaintext string, resolver StrategyResolver) (bool, StoredPassword, error) {
	upgradedPassword := StoredPassword("")

	storedStrategyId, storedSalt, storedBytes, err := deserialize(stored)
	if err != nil {
		return false, upgradedPassword, err
	}

	strategy, betterStrategy := resolver(storedStrategyId)
	if strategy == nil {
		return false, upgradedPassword, errors.New("unknown strategy")
	}

	givenBytes := strategy.Derive([]byte(givenPlaintext), storedSalt)

	if subtle.ConstantTimeCompare(storedBytes, givenBytes) != 1 {
		return false, upgradedPassword, nil
	}

	if betterStrategy != nil {
		upgradedPassword, err = Store(givenPlaintext, betterStrategy)
		if err != nil {
			return true, upgradedPassword, err
		}
	}

	return true, upgradedPassword, nil
}
