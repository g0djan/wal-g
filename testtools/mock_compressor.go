package testtools

import (
	"github.com/g0djan/wal-g/internal"
	"io"
)

type MockCompressor struct{}

func (compressor *MockCompressor) NewWriter(writer io.Writer) internal.ReaderFromWriteCloser {
	return &MockCompressingWriter{writer}
}

func (compressor *MockCompressor) FileExtension() string {
	return "mock"
}

type MockCompressingWriter struct {
	io.Writer
}

func (writer *MockCompressingWriter) ReadFrom(reader io.Reader) (n int64, err error) {
	return internal.FastCopy(writer.Writer, reader)
}

func (writer *MockCompressingWriter) Close() error { return nil }
