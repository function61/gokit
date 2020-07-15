package mac

import (
	"testing"

	"github.com/function61/gokit/assert"
)

const (
	key1 = "fooAuthenticationKey"
	key2 = "barAuthenticationKey"
)

func TestAuthenticate(t *testing.T) {
	msg := "hello world"

	signature := New(key1, msg).Sign()

	assert.EqualString(t, signature, "bo4dqebLiFXPTpqv")

	assert.Ok(t, New(key1, msg).Authenticate(signature))
	assert.Assert(t, New(key1, msg).Authenticate("wrong signature") == ErrMacValidationFailed)
}

func TestDifferentMessagesProduceDifferentSignatures(t *testing.T) {
	assert.EqualString(t, New(key1, "msg A").Sign(), "eq5Wd5Qh2cCgonpj")
	assert.EqualString(t, New(key1, "msg B").Sign(), "HOc4MPpRdV4ytr4r")
}

func TestDifferentKeysProduceDifferentSignatures(t *testing.T) {
	msg := "message to authenticate"

	assert.EqualString(t, New(key1, msg).Sign(), "PwBQuOJaF34tjmv2")
	assert.EqualString(t, New(key2, msg).Sign(), "q8PkJSBswgeTrFcX")
}
