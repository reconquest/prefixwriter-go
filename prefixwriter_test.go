package prefixwriter

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nopCloser struct {
	*bytes.Buffer
}

func (closer nopCloser) Close() error {
	return nil
}

func TestNew_ReturnsWriterWithPrefix(t *testing.T) {
	test := assert.New(t)

	writer := New(nil, "prefix")
	test.Equal("prefix", writer.prefix)
}

func TestPrefixWriter_WriteNothingAtEmptyData(t *testing.T) {
	testWriter(t, nil, "prefix", "", "")
}

func TestPrefixWriter_WriteFirstByteWithPrefix(t *testing.T) {
	testWriter(t, nil, "output: ", "1", "output: 1")
}

func TestPrefixWriter_NotInsertsPrefixOnTwoWritesWithoutNewline(t *testing.T) {
	var writer io.WriteCloser

	writer = testWriter(t, nil, "output: ", "1", "output: 1")
	writer = testWriter(t, writer, "output:", "2", "output: 12")
	writer = testWriter(t, writer, "output:", "3\n", "output: 123\n")
	_ = testWriter(t, writer, "output:", "4", "output: 123\noutput: 4")
}

func TestPrefixWriter_AddsPrefixOnEachNewLine(t *testing.T) {
	var writer io.WriteCloser

	writer = testWriter(t, nil, "p: ", "1\n", "p: 1\n")
	writer = testWriter(t, writer, "p:", "2\n", "p: 1\np: 2\n")
	_ = testWriter(t, writer, "p:", "3\n", "p: 1\np: 2\np: 3\n")
}

func TestPrefixWriter_InsertsPrefixWhenNewlineComesFirst(t *testing.T) {
	var writer io.WriteCloser

	writer = testWriter(t, nil, "p: ", "1", "p: 1")
	writer = testWriter(t, writer, "p:", "\n2", "p: 1\np: 2")
	_ = testWriter(t, writer, "p:", "\n3", "p: 1\np: 2\np: 3")
}

func testWriter(
	t *testing.T,
	writer io.WriteCloser,
	prefix string,
	data string,
	expected string,
) io.WriteCloser {
	test := assert.New(t)

	if writer == nil {
		writer = New(nopCloser{&bytes.Buffer{}}, prefix)
	}

	written, err := writer.Write([]byte(data))
	test.Nil(err)
	test.Equal(len(data), written)

	buffer := writer.(*PrefixWriter).backend.(nopCloser).Buffer

	test.Equal(expected, buffer.String())

	return writer
}
