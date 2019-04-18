// Writes a file to FS using temp-filename scheme, so destination filename should only
// appear when file write is atomically "ok"
package atomicfilewrite

import (
	"fmt"
	"io"
	"os"
)

// - creates a <filename>.part file to write to
// - write all requested content to it (content produced by callback)
// - rename temp filename to filename if everything went OK
//
// If encountering any errors, remove the partial file
func Write(filename string, produce func(io.Writer) error) error {
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

	if err := tempFile.Close(); err != nil {
		return cleanup(err, true) // do not try closing again
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
