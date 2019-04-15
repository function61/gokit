// Wrapper for doing "file exists" -dance in Go (which has a bit too much ceremony)
package fileexists

import (
	"os"
)

// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err // other error
	}

	return true, nil
}
