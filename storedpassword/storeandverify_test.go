package storedpassword

import (
	"github.com/function61/gokit/assert"
	"testing"
)

func alwaysFailingResolver(strategyId string) DerivationStrategy {
	return nil
}

func TestStoreAndVerify(t *testing.T) {
	stored, err := Store("hunter2", CurrentBestDerivationStrategy)
	assert.Assert(t, err == nil)
	assert.Assert(t, len(stored) == 106)

	strategyId, _, _, err := deserialize(stored)
	assert.Assert(t, err == nil)
	assert.EqualString(t, strategyId, "pbkdf2-sha256-100k")

	// pretend above strategy is not found
	match, err := Verify(stored, "hunter2", alwaysFailingResolver)
	assert.Assert(t, match == false)
	assert.EqualString(t, err.Error(), "unknown strategy")

	// strategy should now be found
	match, err = Verify(stored, "hunter2", BuiltinStrategies)
	assert.Assert(t, match == true)
	assert.Assert(t, err == nil)

	match, err = Verify(stored, "hunter INCORRECT", BuiltinStrategies)
	assert.Assert(t, match == false)
	assert.Assert(t, err == nil)
}

func TestUpgradeRequired(t *testing.T) {
	assert.EqualString(t, CurrentBestDerivationStrategy.Id(), "pbkdf2-sha256-100k")

	// upgrade is required if we give in anything other than the current best strategy
	stored, _ := Store("hunter2", CurrentBestDerivationStrategy)

	required, err := UpgradeRequired(stored)
	assert.Assert(t, err == nil)
	assert.Assert(t, required == false)

	outdatedStrategy := &pbkdf2Sha256{"pbkdf2-sha256-1", 1}

	stored, _ = Store("hunter2", outdatedStrategy)

	required, err = UpgradeRequired(stored)
	assert.Assert(t, err == nil)
	assert.Assert(t, required == true)

	required, err = UpgradeRequired(StoredPassword("malformed password"))
	assert.EqualString(t, err.Error(), "invalid structure of StoredPassword")
}
