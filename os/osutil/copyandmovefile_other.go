// +build !linux

package osutil

import (
	"io/fs"
)

func getExtendedFileInfo(fi fs.FileInfo) *extendedFileInfo {
	return nil
}
