package osutil

// typesafe interface for specifying "chmod bits" in human readable way

import (
	"os"
)

type ownerBits os.FileMode
type groupBits os.FileMode
type otherBits os.FileMode

// leaving out combinations that semantically don't make sense:
// (X), (W,X) - execute without read
// (W) - write without read
const (
	// owner means "user", but I chose to go against the flow because it's more semantic
	OwnerNone ownerBits = 0b000000000
	OwnerR    ownerBits = 0b100000000
	OwnerRX   ownerBits = 0b101000000
	OwnerRW   ownerBits = 0b110000000
	OwnerRWX  ownerBits = 0b111000000

	GroupNone groupBits = 0b000000000
	GroupR    groupBits = 0b000100000
	GroupRX   groupBits = 0b000101000
	GroupRW   groupBits = 0b000110000
	GroupRWX  groupBits = 0b000111000

	OtherNone otherBits = 0b000000000
	OtherR    otherBits = 0b000000100
	OtherRX   otherBits = 0b000000101
	OtherRW   otherBits = 0b000000110
	OtherRWX  otherBits = 0b000000111
)

// typesafe interface for specifying "chmod bits" in human readable way
func FileMode(owner ownerBits, group groupBits, other otherBits) os.FileMode {
	return os.FileMode(owner) | os.FileMode(group) | os.FileMode(other)
}
