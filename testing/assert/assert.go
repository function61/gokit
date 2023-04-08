// Assertion helpers
package assert

import (
	"encoding/json"
	"regexp"
	"testing"
)

func Equal[T comparable](t testing.TB, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Fatalf("exp=%v; got=%v", expected, actual)
	}
}

func EqualJSON(t testing.TB, obj interface{}, expected string) {
	t.Helper()

	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}

	Equal(t, string(jsonBytes), expected)
}

func Ok(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func Matches(t testing.TB, actual string, pattern string) {
	t.Helper()

	if !regexp.MustCompile(pattern).MatchString(actual) {
		t.Fatalf("expected %s to match %s", pattern, actual)
	}
}

// Deprecated: use fn with JSON uppercased
func EqualJson(t testing.TB, obj interface{}, expected string) {
	t.Helper()

	EqualJSON(t, obj, expected)
}

// Deprecated: use `Equal(t, expr, true)`
// replacing `Assert(t, attempts == 1)` with `Equal(t, attempts, 1)` allows for better error messages.
func Assert(t testing.TB, expr bool) {
	t.Helper()

	Equal(t, expr, true)
}

// Deprecated: use Equal()
func EqualString(t testing.TB, actual string, expected string) {
	t.Helper()

	Equal(t, actual, expected)
}

// Deprecated: use Equal()
func EqualInt(t testing.TB, actual int, expected int) {
	t.Helper()

	Equal(t, actual, expected)
}
