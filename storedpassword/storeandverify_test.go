package storedpassword

import (
	"github.com/function61/gokit/assert"
	"testing"
)

func TestStoreAndVerify(t *testing.T) {
	stored, err := Store("hunter2", CurrentBestDerivationStrategy)
	assert.Assert(t, err == nil)
	assert.Assert(t, len(stored) == 107)

	strategyId, _, _, err := deserialize(stored)
	assert.Assert(t, err == nil)
	assert.EqualString(t, strategyId, "pbkdf2-sha256-100k")

	// pretend above strategy is not found
	match, upgrade, err := Verify(stored, "hunter2", alwaysFailingResolver)
	assert.Assert(t, match == false)
	assert.EqualString(t, err.Error(), "unknown strategy")
	assert.Assert(t, upgrade == "")

	// strategy should now be found
	match, upgrade, err = Verify(stored, "hunter2", BuiltinStrategies)
	assert.Assert(t, match == true)
	assert.Assert(t, err == nil)
	assert.Assert(t, upgrade == "")

	match, upgrade, err = Verify(stored, "hunter INCORRECT", BuiltinStrategies)
	assert.Assert(t, match == false)
	assert.Assert(t, err == nil)
	assert.Assert(t, upgrade == "")

	// Verify() should now suggest upgrade with this resolver
	match, upgrade, err = Verify(stored, "hunter2", downgradingResolver)
	assert.Assert(t, match == true)
	assert.Assert(t, err == nil)
	assert.Assert(t, upgrade != "")

	// upgraded password should now use the ridiculously insecure strategy
	strategyId, _, _, err = deserialize(upgrade)
	assert.Assert(t, err == nil)
	assert.EqualString(t, strategyId, "pbkdf2-sha256-1")

	// verify upgraded password
	match, upgrade, err = Verify(upgrade, "hunter2", downgradingResolver)
	assert.Assert(t, match == true)
	assert.Assert(t, err == nil)
	assert.Assert(t, upgrade == "")
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
