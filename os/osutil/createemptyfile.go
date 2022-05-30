package osutil

import (
	"os"

	. "github.com/function61/gokit/builtin"
)

// does the same as "$ touch"
func CreateEmptyFile(path string) error {
	return ErrorWrap("CreateEmptyFile", func() error {
		file, err := os.Create(path)
		if err != nil {
			return err
		}

		return file.Close()
	}())
}
