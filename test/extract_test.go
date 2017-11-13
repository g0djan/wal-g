package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/x4m/wal-g/internal"
	"github.com/x4m/wal-g/testtools"
	"io"
	"io/ioutil"
	"testing"
)

func TestNoFilesProvided(t *testing.T) {
	buf := &testtools.NOPTarInterpreter{}
	err := internal.ExtractAll(buf, []internal.ReaderMaker{})
	assert.IsType(t, err, internal.NoFilesToExtractError{})
}

// Tests roundtrip for a tar file.
func TestTar(t *testing.T) {
	// Generate and save random bytes compare against compression-decompression cycle.
	sb := testtools.NewStrideByteReader(10)
	lr := &io.LimitedReader{
		R: sb,
		N: int64(1024),
	}
	b, err := ioutil.ReadAll(lr)

	// Copy generated bytes to another slice to make the test more robust against modifications of "b".
	bCopy := make([]byte, len(b))
	copy(bCopy, b)
	assert.NoError(t, err)

	// Make a tar in memory.
	member := &bytes.Buffer{}
	testtools.CreateTar(member, &io.LimitedReader{
		R: bytes.NewBuffer(b),
		N: int64(len(b)),
	})

	// Extract the generated tar and check that its one member is the same as the bytes generated to begin with.
	brm := &BufferReaderMaker{member, "/usr/local/file.tar"}
	buf := &testtools.BufferTarInterpreter{}
	files := []internal.ReaderMaker{brm}
	err = internal.ExtractAll(buf, files)
	if err != nil {
		t.Log(err)
	}

	assert.Equalf(t, bCopy, buf.Out, "extract: Unbundled tar output does not match input.")
}

//func TestExtractAll(t *testing.T) {
//	os.Setenv("WALE_GPG_KEY_ID", "3C19717A2B308DF0")
//	os.Setenv("WALG_DOWNLOAD_CONCURRENCY", "1")
//	defer os.Unsetenv("WALE_GPG_KEY_ID")
//	defer os.Unsetenv("WALG_DOWNLOAD_CONCURRENCY")
//	readerMaker := &testtools.FileReaderMaker{Key: "testdata/part_009.tar.zst"}
//	err := internal.ExtractAll(&testtools.NOPTarInterpreter{}, []internal.ReaderMaker {readerMaker})
//	assert.NoError(t, err)
//}

//func TestDecryptAndDecompressTar(t *testing.T) {
//	os.Setenv("WALE_GPG_KEY_ID", "3C19717A2B308DF0")
//	defer os.Unsetenv("WALE_GPG_KEY_ID")
//	readerMaker := &testtools.FileReaderMaker{Key: "testdata/part_021.tar.zst"}
//	result, err := os.Create("testdata/part_021.tar")
//	assert.NoError(t, err)
//	defer result.Close()
//	var crypter internal.OpenPGPCrypter
//	err = internal.DecryptAndDecompressTar(result, readerMaker, &crypter)
//	assert.NoError(t, err)
//}

// Used to mock files in memory.
type BufferReaderMaker struct {
	Buf *bytes.Buffer
	Key string
}

func (b *BufferReaderMaker) Reader() (io.ReadCloser, error) { return ioutil.NopCloser(b.Buf), nil }
func (b *BufferReaderMaker) Path() string                   { return b.Key }
