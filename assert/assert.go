package assert

import (
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

func Assert(t *testing.T, expr bool) {
	t.Helper()

	if !expr {
		t.Fatal("Expecting true; got false")
	}
}
