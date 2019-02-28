package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/x4m/wal-g/internal"
	"github.com/x4m/wal-g/internal/walparser"
	"testing"
)

func TestExtractBlockLocations(t *testing.T) {
	record, _ := GetXLogRecordData()
	expectedLocations := []walparser.BlockLocation{record.Blocks[0].Header.BlockLocation}
	actualLocations := internal.ExtractBlockLocations([]walparser.XLogRecord{record})
	assert.Equal(t, expectedLocations, actualLocations)
}
