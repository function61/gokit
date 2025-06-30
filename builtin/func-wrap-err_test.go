package builtin

import (
	"errors"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestFuncWrapErr(t *testing.T) {
	assert.Equal(t, FuncWrapErr(errors.New("imaginary error")).Error(), "TestFuncWrapErr: imaginary error")
}

func TestPackageFuncWrapErr(t *testing.T) {
	assert.Equal(t, PackageFuncWrapErr(errors.New("imaginary error")).Error(), "builtin.TestPackageFuncWrapErr: imaginary error")
}

func TestFunctionAndPkgName(t *testing.T) {
	for _, tc := range []struct {
		input  string
		output string
		pkg    string
	}{
		{"github.com/function61/gokit/builtin.TestFoo", "TestFoo", "builtin"},
		{"main.TestFoo", "TestFoo", "main"},
		{"github.com/function61/gokit/builtin.TestFoo.func2.TestFoo.func2.1.2", "TestFoo", "builtin"},
		{"github.com/function61/gokit/builtin.(*stuffReader).Read", "(*stuffReader).Read", "builtin"},
	} {
		t.Run(tc.input, func(t *testing.T) {
			fnName, pkgBasename := functionAndPkgName(tc.input)
			assert.Equal(t, fnName, tc.output)
			assert.Equal(t, pkgBasename, tc.pkg)
		})
	}
}

func TestStringsIndexOffset(t *testing.T) {
	assert.Equal(t, stringsIndexOffset("foo.bar.baz", ".", 0), 3)
	assert.Equal(t, stringsIndexOffset("foo.bar.baz", ".", 3), 3)
	assert.Equal(t, stringsIndexOffset("foo.bar.baz", ".", 4), 7)
	assert.Equal(t, stringsIndexOffset("foo.bar.baz", "#", 0), -1)
}
