package osutil

import (
	"os"

	. "github.com/function61/gokit/builtin"
)

// does the same as "$ touch"
func CreateEmptyFile(path string) error {
	withErr := FuncWrapErr

	file, err := os.Create(path)
	if err != nil {
		return withErr(err)
	}

	if err := file.Close(); err != nil {
		return withErr(err)
	}

	return nil
}
