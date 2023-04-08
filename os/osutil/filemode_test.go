package osutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestFileMode(t *testing.T) {
	// combos

	assert.Equal(t, FileMode(OwnerR, GroupRW, OtherRWX), 0467)
	assert.Equal(t, FileMode(OwnerRWX, GroupRWX, OtherRWX), 0777)
	assert.Equal(t, FileMode(OwnerNone, GroupNone, OtherRWX), 0007)

	// each individually

	assert.Equal(t, FileMode(OwnerNone, GroupNone, OtherNone), 0000)
	assert.Equal(t, FileMode(OwnerR, GroupNone, OtherNone), 0400)
	assert.Equal(t, FileMode(OwnerRX, GroupNone, OtherNone), 0500)
	assert.Equal(t, FileMode(OwnerRW, GroupNone, OtherNone), 0600)
	assert.Equal(t, FileMode(OwnerRWX, GroupNone, OtherNone), 0700)

	assert.Equal(t, FileMode(OwnerNone, GroupNone, OtherNone), 0000)
	assert.Equal(t, FileMode(OwnerNone, GroupR, OtherNone), 0040)
	assert.Equal(t, FileMode(OwnerNone, GroupRX, OtherNone), 0050)
	assert.Equal(t, FileMode(OwnerNone, GroupRW, OtherNone), 0060)
	assert.Equal(t, FileMode(OwnerNone, GroupRWX, OtherNone), 0070)

	assert.Equal(t, FileMode(OwnerNone, GroupNone, OtherNone), 0000)
	assert.Equal(t, FileMode(OwnerNone, GroupNone, OtherR), 0004)
	assert.Equal(t, FileMode(OwnerNone, GroupNone, OtherRX), 0005)
	assert.Equal(t, FileMode(OwnerNone, GroupNone, OtherRW), 0006)
	assert.Equal(t, FileMode(OwnerNone, GroupNone, OtherRWX), 0007)
}
