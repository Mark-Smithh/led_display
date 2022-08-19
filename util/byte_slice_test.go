package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanDisplayNumberByteFalse(t *testing.T) {
	imp := ByteSliceImplementation{DisplayParams{NumToDisplay: 1, NumLedSegments: 0}}
	result := imp.CanDisplayNumber()

	assert.False(t, result, "expected result to equal false")
}

func TestCanDisplayNumberByteTrue(t *testing.T) {
	imp := ByteSliceImplementation{DisplayParams{NumToDisplay: 1, NumLedSegments: 10}}
	result := imp.CanDisplayNumber()

	assert.True(t, result, "expected result to equal true")
}
