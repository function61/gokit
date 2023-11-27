package ezhttp

import (
	"context"
	"strings"
	"testing"

	. "github.com/function61/gokit/builtin"
	"github.com/function61/gokit/testing/assert"
)

func TestCURLEquivalent(t *testing.T) {
	curlCmd := Must(NewPost(context.Background(), "https://example.net/hello", Header("x-correlation-id", "123")).CURLEquivalent())

	assert.Equal(t, strings.Join(curlCmd, " "), "curl --request=POST --header=X-Correlation-Id=123 https://example.net/hello")
}
