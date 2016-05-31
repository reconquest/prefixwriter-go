package prefixwriter

import (
	"bytes"
	"io"
)

// Writer writes specified data, prepending prefix to each new line.
type Writer struct {
	backend io.WriteCloser
	prefix  string

	streamStarted  bool
	lineIncomplete bool
}

// New creates new Writer, that will use `writer` as backend
// and will prepend `prefix` to each line.
func New(writer io.WriteCloser, prefix string) *Writer {
	return &Writer{
		backend: writer,
		prefix:  prefix,
	}
}

// Writer writes data into Writer.
//
// Signature matches with io.Writer's Write().
func (writer *Writer) Write(data []byte) (int, error) {
	var (
		reader         = bytes.NewBuffer(data)
		eofEncountered = false
	)

	for !eofEncountered {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return 0, err
			}

			eofEncountered = true
		}

		if line == "" {
			continue
		}

		if !writer.streamStarted || !writer.lineIncomplete {
			line = writer.prefix + line

			writer.streamStarted = true
		}

		writer.lineIncomplete = eofEncountered

		_, err = writer.backend.Write([]byte(line))
		if err != nil {
			return 0, err
		}
	}

	return len(data), nil
}

// Close closes underlying backend writer.
//
// Signature matches with io.WriteCloser's Close().
func (writer *Writer) Close() error {
	return writer.backend.Close()
}
