package storedpassword

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

// entries from here can never be removed to preserve backwards compatibility,
// but you should always use UpgradeRequired() to enable automatically upgrading to
// state-of-the-art.
var builtinStrategies = map[string]DerivationStrategy{
	// 1.4sec @ 100k on Raspberry Pi 2
	// https://github.com/borgbackup/borg/issues/77#issuecomment-130459726
	"pbkdf2-sha256-100k": &pbkdf2Sha256{"pbkdf2-sha256-100k", 100 * 1000},
}

// this works as a StrategyResolver to Verify()
func BuiltinStrategies(id string) (DerivationStrategy, DerivationStrategy) {
	strategy, found := builtinStrategies[id]
	if !found {
		return nil, nil
	}

	// only non-nil if upgrade needed
	upgradeStrategy := func() DerivationStrategy {
		if strategy.Id() != CurrentBestDerivationStrategy.Id() { // need upgrade
			return CurrentBestDerivationStrategy
		} else {
			return nil
		}
	}()

	return strategy, upgradeStrategy
}

// provide this to Store() to automatically keep newly generated passwords
// up-to-date according to current recommendations
var CurrentBestDerivationStrategy = builtinStrategies["pbkdf2-sha256-100k"]

type pbkdf2Sha256 struct {
	id         string
	iterations int
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
