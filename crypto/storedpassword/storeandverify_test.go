package storedpassword

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestStoreAndVerify(t *testing.T) {
	stored, err := Store("hunter2", CurrentBestDerivationStrategy)
	assert.Ok(t, err)
	assert.Equal(t, len(stored), 107)

	strategyId, _, _, err := deserialize(stored)
	assert.Ok(t, err)
	assert.Equal(t, strategyId, "pbkdf2-sha256-100k")

	// pretend above strategy is not found
	upgrade, err := Verify(stored, "hunter2", alwaysFailingResolver)
	assert.Equal(t, err.Error(), "unknown strategy")
	assert.Equal(t, upgrade, "")

	// strategy should now be found
	upgrade, err = Verify(stored, "hunter2", BuiltinStrategies)
	assert.Ok(t, err)
	assert.Equal(t, upgrade, "")

	upgrade, err = Verify(stored, "hunter INCORRECT", BuiltinStrategies)
	assert.Equal(t, err, ErrIncorrectPassword)
	assert.Equal(t, upgrade, "")

	// Verify() should now suggest upgrade with this resolver
	upgrade, err = Verify(stored, "hunter2", downgradingResolver)
	assert.Ok(t, err)
	assert.Equal(t, upgrade != "", true)

	// upgraded password should now use the ridiculously insecure strategy
	strategyId, _, _, err = deserialize(upgrade)
	assert.Ok(t, err)
	assert.Equal(t, strategyId, "pbkdf2-sha256-1")

	// verify upgraded password
	upgrade, err = Verify(upgrade, "hunter2", downgradingResolver)
	assert.Ok(t, err)
	assert.Equal(t, upgrade, "")
}

func alwaysFailingResolver(strategyId string) (DerivationStrategy, DerivationStrategy) {
	return nil, nil
}

var insecureStrategy = &pbkdf2Sha256{"pbkdf2-sha256-1", 1}

func downgradingResolver(strategyId string) (DerivationStrategy, DerivationStrategy) {
	if strategyId == insecureStrategy.Id() {
		return insecureStrategy, nil
	}

	strategy, _ := BuiltinStrategies(strategyId)
	if strategy == nil {
		return nil, nil
	}

	// recommend upgrade (actuall downgrade) to with just one iteration
	// (this is ridiculous example, I know)
	return strategy, insecureStrategy
}
