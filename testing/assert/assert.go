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

func EqualJson(t testing.TB, obj interface{}, expected string) {
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

type Ass struct {
	testing.TB
}

func New(t testing.TB) *Ass {
	return &Ass{t}
}

func (a *Ass) EqualString(actual string, expected string) {
	a.Helper()
	EqualString(a, actual, expected)
}

func (a *Ass) EqualInt(actual int, expected int) {
	a.Helper()
	EqualInt(a, actual, expected)
}

func (a *Ass) EqualJson(obj interface{}, expected string) {
	a.Helper()
	EqualJson(a, obj, expected)
}

func (a *Ass) Assert(expr bool) {
	a.Helper()
	Assert(a, expr)
}

func (a *Ass) Ok(err error) {
	a.Helper()
	Ok(a, err)
}

func (a *Ass) Matches(t testing.TB, actual string, pattern string) {
	a.Helper()
	Matches(a, actual, pattern)
}
