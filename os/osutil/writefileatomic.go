// Writes a file to FS using temp-filename scheme, so destination filename should only
// appear when file write is atomically "ok"
package osutil

import (
	"fmt"
	"io"
	"io/fs"
	"os"
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

	tempFilename := filename + ".part"

	tempFile, err := os.Create(tempFilename)
	if err != nil {
		return err
	}

	cleanup := func(errRet error, fileClosed bool) error {
		if !fileClosed {
			if err := tempFile.Close(); err != nil {
				errRet = fmt.Errorf("%v; also failed closing: %v", errRet, err)
			}
		}
		if err := os.Remove(tempFilename); err != nil {
			errRet = fmt.Errorf("%v; also failed removing temp file: %v", errRet, err)
		}

		return errRet
	}

	// why WriteNopCloser? while we appreciate that some producers are good citizens, when
	// given io.Writer they still might check for io.Closer and call Close(), messing
	// up our cleanup since we return error if close failed. and it'd be messy to try to
	// check for Close() error being "already closed"
	if err := produce(newWriteNopCloser(tempFile)); err != nil {
		return cleanup(err, false)
	}

	if opts.mode != nil {
		if err := tempFile.Chmod(*opts.mode); err != nil {
			return err
		}
	}

	if opts.uid != nil {
		if err := tempFile.Chown(*opts.uid, *opts.gid); err != nil && !opts.ignoreIfChownErrors {
			return err
		}
	}

	if err := tempFile.Close(); err != nil {
		return cleanup(err, true) // do not try closing again
	}

	if opts.atime != nil { // unfortunately this has to be done after Close()
		if err := os.Chtimes(tempFilename, *opts.atime, *opts.mtime); err != nil {
			return err
		}
	}

	if err := os.Rename(tempFilename, filename); err != nil {
		return cleanup(err, true)
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
