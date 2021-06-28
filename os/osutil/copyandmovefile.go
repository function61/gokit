package osutil

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
)

// tries its best to preserve metadata like file permissions & timestamps
func CopyFile(sourcePath string, destinationPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close() // double close intentional

	metadata, err := source.Stat()
	if err != nil {
		return err
	}

	// preserve permissions & timestamps.
	// TODO: preserve xattrs?
	opts := []WriteFileOption{
		WriteFileMode(metadata.Mode().Perm()),
	}

	// not all platforms have uid/gid + atime + ctime available
	if statT, ok := metadata.Sys().(*syscall.Stat_t); ok {
		opts = append(
			opts,
			WriteFileChown(int(statT.Uid), int(statT.Gid)),
			WriteFileTimes(timespecToTime(statT.Atim), timespecToTime(statT.Mtim), timespecToTime(statT.Ctim)),
		)
	} else {
		opts = append(
			opts,
			WriteFileTimes(metadata.ModTime(), metadata.ModTime(), metadata.ModTime()), // all we have is mtime
		)
	}

	if err := WriteFileAtomic(
		destinationPath,
		func(sink io.Writer) error {
			_, err := io.Copy(sink, source)
			return err
		},
		opts...,
	); err != nil {
		return err
	}

	return source.Close()
}

// first tries by renaming, if that errors falls back to copying (so it works across filesystems)
func MoveFile(sourcePath string, destinationPath string) error {
	// first try by renaming (it is the most efficient)
	if err := os.Rename(sourcePath, destinationPath); err == nil {
		return nil
	}

	// rename errored.
	// there is no sane way to detect if this was because the file crosses filesystems (beyond matching
	// error string...) so we're going to just assume that that was the case, and try copying instead.

	if err := CopyFile(sourcePath, destinationPath); err != nil {
		return fmt.Errorf("MoveFile: copy: %w", err)
	}

	if err := os.Remove(sourcePath); err != nil {
		return fmt.Errorf("MoveFile: %w", err)
	}

	return nil
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}
