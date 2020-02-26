// Assertion helpers
package assert

import (
	"encoding/json"
	"regexp"
	"testing"
)

// waiting for generics to arrive to Go..

func EqualString(t *testing.T, actual string, expected string) {
	t.Helper()

	if actual != expected {
		t.Fatalf("exp=%v; got=%v", expected, actual)
	}
}

func EqualInt(t *testing.T, actual int, expected int) {
	t.Helper()

	if actual != expected {
		t.Fatalf("exp=%v; got=%v", expected, actual)
	}
}

func EqualJson(t *testing.T, obj interface{}, expected string) {
	t.Helper()

	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}

	EqualString(t, string(jsonBytes), expected)
}

func Assert(t *testing.T, expr bool) {
	t.Helper()

	if !expr {
		t.Fatal("Expecting true; got false")
	}
}

func Ok(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func Matches(t *testing.T, actual string, pattern string) {
	t.Helper()

	if !regexp.MustCompile(pattern).MatchString(actual) {
		t.Fatalf("expected %s to match %s", pattern, actual)
	}
}
