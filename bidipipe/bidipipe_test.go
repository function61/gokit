package bidipipe

import (
	"errors"
	"io"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestPipe(t *testing.T) {
	person1 := &ReaderWriterMock{writes: "Nice to meet you, Person2"}
	person2 := &ReaderWriterMock{writes: "Thanks Person1, pleasure is all mine"}

	err := Pipe(person1, "person1", person2, "person2")

	assert.Ok(t, err)
	assert.EqualString(t, person2.received, "Nice to meet you, Person2")
	assert.EqualString(t, person1.received, "Thanks Person1, pleasure is all mine")
}

func TestPipeWithErrorFromPerson1(t *testing.T) {
	person1 := &ReaderWriterMock{writes: "Nice to meet you, Person2", readsWillError: true}
	person2 := &ReaderWriterMock{writes: "Thanks Person1, pleasure is all mine"}

	err := Pipe(person1, "person1", person2, "person2")

	assert.EqualString(t, err.Error(), "bidipipe: person1 -> person2 error: hurr durr, I broke")
	assert.EqualString(t, person2.received, "")
	assert.EqualString(t, person1.received, "Thanks Person1, pleasure is all mine")
}

func TestPipeWithErrorFromPerson2(t *testing.T) {
	person1 := &ReaderWriterMock{writes: "Nice to meet you, Person2"}
	person2 := &ReaderWriterMock{writes: "Thanks Person1, pleasure is all mine", readsWillError: true}

	err := Pipe(person1, "person1", person2, "person2")

	assert.EqualString(t, err.Error(), "bidipipe: person2 -> person1 error: hurr durr, I broke")
	assert.EqualString(t, person2.received, "Nice to meet you, Person2")
	assert.EqualString(t, person1.received, "")
}

type ReaderWriterMock struct {
	writes         string
	received       string
	readAlready    bool
	readsWillError bool
}

// implement io.ReadWriteCloser

func (m *ReaderWriterMock) Read(p []byte) (int, error) {
	if m.readsWillError {
		return 0, errors.New("hurr durr, I broke")
	}

	if m.readAlready {
		return 0, io.EOF
	}

	copy(p[:], m.writes)
	m.readAlready = true
	return len(m.writes), nil
}

func (m *ReaderWriterMock) Write(p []byte) (int, error) {
	m.received = string(p)
	return len(p), nil
}

func (m *ReaderWriterMock) Close() error {
	return nil
}
