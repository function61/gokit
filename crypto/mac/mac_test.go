package mac

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

const (
	key1 = "fooAuthenticationKey"
	key2 = "barAuthenticationKey"
)

func TestAuthenticate(t *testing.T) {
	msg := "hello world"

	signature := New(key1, msg).Sign()

	assert.Equal(t, signature, "bo4dqebLiFXPTpqv")

	assert.Ok(t, New(key1, msg).Authenticate(signature))
	assert.Equal(t, New(key1, msg).Authenticate("wrong signature"), ErrMacValidationFailed)
}

func TestDifferentMessagesProduceDifferentSignatures(t *testing.T) {
	assert.Equal(t, New(key1, "msg A").Sign(), "eq5Wd5Qh2cCgonpj")
	assert.Equal(t, New(key1, "msg B").Sign(), "HOc4MPpRdV4ytr4r")
}

func TestDifferentKeysProduceDifferentSignatures(t *testing.T) {
	msg := "message to authenticate"

	assert.Equal(t, New(key1, msg).Sign(), "PwBQuOJaF34tjmv2")
	assert.Equal(t, New(key2, msg).Sign(), "q8PkJSBswgeTrFcX")
}
