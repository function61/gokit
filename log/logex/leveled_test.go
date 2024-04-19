package logex

import (
	"bytes"
	"log"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestLevels(t *testing.T) {
	logOutput := bytes.Buffer{}
	logl := Levels(log.New(&logOutput, "", 0))

	logl.Debug.Println("currentCount=13")
	logl.Info.Println("user 123 logged in")
	logl.Warn.Println("not looking too good")
	logl.Error.Println("I/O failure on device sda1")

	assert.Equal(t, logOutput.String(), `[DEBUG] currentCount=13
[INFO] user 123 logged in
[WARN] not looking too good
[ERROR] I/O failure on device sda1
`)
}
