package storedpassword

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

// entries from here can never be removed to preserve backwards compatibility,
// but you should always use UpgradeRequired() to enable automatically upgrading to
// state-of-the-art.
var builtinStrategies = map[string]DerivationStrategy{
	NewPbkdf2Sha256At100k().Id(): NewPbkdf2Sha256At100k(),
}

// provide this to Store() to automatically keep newly generated passwords
// up-to-date according to current recommendations
var CurrentBestDerivationStrategy = NewPbkdf2Sha256At100k()

// this works as a resolver to Verify()
func BuiltinStrategies(id string) DerivationStrategy {
	strategy, found := builtinStrategies[id]
	if !found {
		return nil
	}

	return strategy
}

// with this function you can easily check if the stored password needs to be upgraded
// during user login to a stronger strategy. this does not mean changing the plaintext password.
func UpgradeRequired(stored StoredPassword) (bool, error) {
	strategyId, _, _, err := deserialize(stored)
	if err != nil {
		return false, err
	}

	return strategyId != CurrentBestDerivationStrategy.Id(), nil
}

type pbkdf2Sha256 struct {
	id         string
	iterations int
}

// 1.4sec @ 100k on Raspberry Pi 2
// https://github.com/borgbackup/borg/issues/77#issuecomment-130459726
func NewPbkdf2Sha256At100k() DerivationStrategy {
	return &pbkdf2Sha256{"pbkdf2-sha256-100k", 100 * 1000}
}

func (p *pbkdf2Sha256) Id() string { return p.id }

func (p *pbkdf2Sha256) Derive(plaintext []byte, salt []byte) []byte {
	derivedKey := pbkdf2.Key(
		plaintext,
		salt,
		p.iterations,
		sha256.Size,
		sha256.New)

	if len(derivedKey) != sha256.Size {
		panic("pbkdf2 derived key unexpected length")
	}

	return derivedKey
}
