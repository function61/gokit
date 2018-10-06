package assert

import (
	"testing"
)

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

func True(t *testing.T, actual bool) {
	t.Helper()

	if !actual {
		t.Fatal("Expecting true; got false")
	}
}

func False(t *testing.T, actual bool) {
	t.Helper()

	if actual {
		t.Fatal("Expecting true; got false")
	}
}
