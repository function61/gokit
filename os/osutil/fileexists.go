// Wrapper for doing "file exists" -dance in Go (which has a bit too much ceremony)
package osutil

import (
	"os"
)

// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
// follows symlinks
func Exists(path string) (bool, error) {
	return exists(path, os.Stat)
}

func ExistsNoLinkFollow(path string) (bool, error) {
	return exists(path, os.Lstat)
}

func exists(path string, statter func(string) (os.FileInfo, error)) (bool, error) {
	_, err := statter(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err // other error
	}

	return true, nil
}
