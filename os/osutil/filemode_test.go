package osutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestFileMode(t *testing.T) {
	// combos

	assert.Assert(t, FileMode(OwnerR, GroupRW, OtherRWX) == 0467)
	assert.Assert(t, FileMode(OwnerRWX, GroupRWX, OtherRWX) == 0777)
	assert.Assert(t, FileMode(OwnerNone, GroupNone, OtherRWX) == 0007)

	// each individually

	assert.Assert(t, FileMode(OwnerNone, GroupNone, OtherNone) == 0000)
	assert.Assert(t, FileMode(OwnerR, GroupNone, OtherNone) == 0400)
	assert.Assert(t, FileMode(OwnerRX, GroupNone, OtherNone) == 0500)
	assert.Assert(t, FileMode(OwnerRW, GroupNone, OtherNone) == 0600)
	assert.Assert(t, FileMode(OwnerRWX, GroupNone, OtherNone) == 0700)

	assert.Assert(t, FileMode(OwnerNone, GroupNone, OtherNone) == 0000)
	assert.Assert(t, FileMode(OwnerNone, GroupR, OtherNone) == 0040)
	assert.Assert(t, FileMode(OwnerNone, GroupRX, OtherNone) == 0050)
	assert.Assert(t, FileMode(OwnerNone, GroupRW, OtherNone) == 0060)
	assert.Assert(t, FileMode(OwnerNone, GroupRWX, OtherNone) == 0070)

	assert.Assert(t, FileMode(OwnerNone, GroupNone, OtherNone) == 0000)
	assert.Assert(t, FileMode(OwnerNone, GroupNone, OtherR) == 0004)
	assert.Assert(t, FileMode(OwnerNone, GroupNone, OtherRX) == 0005)
	assert.Assert(t, FileMode(OwnerNone, GroupNone, OtherRW) == 0006)
	assert.Assert(t, FileMode(OwnerNone, GroupNone, OtherRWX) == 0007)
}
