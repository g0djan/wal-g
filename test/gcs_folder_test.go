package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/g0djan/wal-g/internal"
)

func TestGSFolder(t *testing.T) {
	t.Skip("Credentials needed to run GCP tests")

	storageFolder, err := internal.ConfigureGSFolder("gs://g0djan-test/walg-bucket")

	assert.NoError(t, err)

	testStorageFolder(storageFolder, t)
}
