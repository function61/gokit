// +build linux

package osutil

import (
	"io/fs"
	"syscall"
	"time"
)

func getExtendedFileInfo(fi fs.FileInfo) *extendedFileInfo {
	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return nil
	}

	return &extendedFileInfo{
		Uid:  int(stat.Uid),
		Gid:  int(stat.Gid),
		Atim: timespecToTime(stat.Atim),
		Mtim: timespecToTime(stat.Mtim),
		Ctim: timespecToTime(stat.Ctim),
	}
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}
