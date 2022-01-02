// Assertion helpers
package assert

import (
	"encoding/json"
	"regexp"
	"testing"
)

// waiting for generics to arrive to Go..

func EqualString(t testing.TB, actual string, expected string) {
	t.Helper()

	if actual != expected {
		t.Fatalf("exp=%v; got=%v", expected, actual)
	}
}

func EqualInt(t testing.TB, actual int, expected int) {
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

	EqualString(t, string(jsonBytes), expected)
}

func Assert(t testing.TB, expr bool) {
	t.Helper()

	if !expr {
		t.Fatal("Expecting true; got false")
	}
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
