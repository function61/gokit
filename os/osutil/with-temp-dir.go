package osutil

import (
	"os"

	. "github.com/function61/gokit/builtin"
)

// creates a temporary directory to do `work` in, then deletes the directory after work is done.
func WithTempDir(pattern string, work func(dir string) error) error {
	withErr := FuncWrapErr

	tempDir, err := os.MkdirTemp("", pattern)
	if err != nil {
		return withErr(err)
	}
	defer os.RemoveAll(tempDir)

	return work(tempDir)
}
