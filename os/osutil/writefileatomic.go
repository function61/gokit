// Writes a file to FS using temp-filename scheme, so destination filename should only
// appear when file write is atomically "ok"
package osutil

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"time"
)

// - creates a <filename>.part file to write to
// - write all requested content to it (content produced by callback)
// - rename temp filename to filename if everything went OK
//
// If encountering any errors, remove the partial file
func WriteFileAtomic(filename string, produce func(io.Writer) error, options ...WriteFileOption) error {
	opts := writeFileOptions{}
	for _, option := range options {
		option(&opts)
	}

	return FileAtomicOperationByRename(filename, func(filenameTemp string) error {
		file, err := os.Create(filenameTemp)
		if err != nil {
			return err
		}
		defer file.Close()

		// why WriteNopCloser? while we appreciate that some producers are good citizens, when
		// given io.Writer they still might check for io.Closer and call Close(), messing
		// up our cleanup since we return error if close failed. and it'd be messy to try to
		// check for Close() error being "already closed"
		if err := produce(newWriteNopCloser(file)); err != nil {
			return err
		}

		if opts.mode != nil {
			if err := file.Chmod(*opts.mode); err != nil {
				return err
			}
		}

		if opts.uid != nil {
			if err := file.Chown(*opts.uid, *opts.gid); err != nil && !opts.ignoreIfChownErrors {
				return err
			}
		}

		if err := file.Close(); err != nil { // double close intentional
			return err
		}

		if opts.atime != nil { // Chtimes() can't be done on the handle so this has to be done after Close()
			if err := os.Chtimes(filenameTemp, *opts.atime, *opts.mtime); err != nil {
				return err
			}
		}

		return nil
	})
}

// - ensures a file with given name only appears on the filesystem if all its file operations succeed
// - tries to clean up temp file on failure of *operations* or the atomic rename
func FileAtomicOperationByRename(path string, operations func(pathTemp string) error) error {
	pathTemp := path + ".part"

	retErrorWithCleanup := func(err error) error { // err is non-nil here
		if errCleanup := os.Remove(pathTemp); errCleanup != nil && !os.IsNotExist(errCleanup) {
			// IsNotExist is acceptable from remove, the *operations* didn't manage to start creating
			// the temp file (or was not authorized so)

			return fmt.Errorf("%w; additionally FileAtomicOperationByRename failed cleaning up: %v", err, errCleanup)
		} else {
			return err
		}
	}

	// the heavy lifting happens here: *operations* tries to produce <file>.part, and only if it succeeds
	// we'll try to rename the temp file to the actual requested filename
	if err := operations(pathTemp); err != nil {
		return retErrorWithCleanup(err)
	}

	if err := os.Rename(pathTemp, path); err != nil {
		return retErrorWithCleanup(err)
	}

	return nil
}

type writeNopCloser struct {
	io.Writer
}

func (writeNopCloser) Close() error { return nil }

// NopCloser returns a WriteCloser with a no-op Close method wrapping
// the provided Writer w.
func newWriteNopCloser(w io.Writer) io.WriteCloser {
	return writeNopCloser{w}
}

type writeFileOptions struct {
	mode                *fs.FileMode
	uid                 *int
	gid                 *int
	ignoreIfChownErrors bool
	atime               *time.Time
	mtime               *time.Time
}

type WriteFileOption func(opt *writeFileOptions)

func WriteFileMode(mode fs.FileMode) WriteFileOption {
	return func(opt *writeFileOptions) {
		opt.mode = &mode
	}
}

func WriteFileChown(uid int, gid int) WriteFileOption {
	return func(opt *writeFileOptions) {
		opt.uid = &uid
		opt.gid = &gid
	}
}

// same as WriteFileChown() but if copying to filesystem that doesn't support Chown(), we'll get an
// error which we'll ignore
func WriteFileChownAndIgnoreIfErrors(uid int, gid int) WriteFileOption {
	return func(opt *writeFileOptions) {
		opt.uid = &uid
		opt.gid = &gid
		opt.ignoreIfChownErrors = true
	}
}

// warning: ctime is ignored for now
func WriteFileTimes(atime time.Time, mtime time.Time, ctime time.Time) WriteFileOption {
	return func(opt *writeFileOptions) {
		opt.atime = &atime
		opt.mtime = &mtime
	}
}

// if running under '$ sudo', set the sudo-invoking "original" user as the UID & GID
func WriteFileIfSudoPreserveInvokingUsersUIDAndGID() WriteFileOption {
	uid, gid := getInvokingUserUidAndGidIfRunningInSudo()

	return func(opt *writeFileOptions) {
		opt.uid = &uid
		opt.gid = &gid
	}
}

// if running under '$ sudo', return invoking user's uid:gid pair.
func getInvokingUserUidAndGidIfRunningInSudo() (int, int) {
	standardUidAndGid := func() (int, int) {
		return os.Getuid(), os.Getgid()
	}

	// documented in https://www.sudo.ws/docs/man/sudo.man/#SUDO_UID
	if os.Getenv("SUDO_UID") == "" { // not running under sudo
		return standardUidAndGid()
	}

	uidSudo, err := strconv.Atoi(os.Getenv("SUDO_UID"))
	if err != nil {
		return standardUidAndGid()
	}

	gidSudo, err := strconv.Atoi(os.Getenv("SUDO_GID"))
	if err != nil {
		return standardUidAndGid()
	}

	return uidSudo, gidSudo
}
