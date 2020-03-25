package storedpassword

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

// this works as a StrategyResolver to Verify()
func BuiltinStrategies(id string) (DerivationStrategy, DerivationStrategy) {
	var strategy DerivationStrategy

	// entries from here can never be removed to preserve backwards compatibility,
	// but you should always use UpgradeRequired() to enable automatically upgrading to
	// state-of-the-art.
	switch id {
	case "pbkdf2-sha256-100k":
		// 1.4sec @ 100k on Raspberry Pi 2
		// https://github.com/borgbackup/borg/issues/77#issuecomment-130459726
		strategy = &pbkdf2Sha256{id, 100 * 1000}
	}

	var upgradeStrategy DerivationStrategy

	if strategy != nil && strategy.Id() != CurrentBestDerivationStrategy.Id() {
		upgradeStrategy = CurrentBestDerivationStrategy
	}

	return strategy, upgradeStrategy
}

// provide this to Store() to automatically keep newly generated passwords
// up-to-date according to current recommendations
var CurrentBestDerivationStrategy = &pbkdf2Sha256{"pbkdf2-sha256-100k", 100 * 1000}

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
