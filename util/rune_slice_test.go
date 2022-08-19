package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanDisplayNumberRuneFalse(t *testing.T) {
	imp := RuneSliceImplementation{DisplayParams{NumToDisplay: 1, NumLedSegments: 0}}
	result := imp.CanDisplayNumber()

	assert.False(t, result, "expected result to equal false")
}

func TestCanDisplayNumberRuneTrue(t *testing.T) {
	imp := RuneSliceImplementation{DisplayParams{NumToDisplay: 1, NumLedSegments: 10}}
	result := imp.CanDisplayNumber()

	assert.True(t, result, "expected result to equal true")
}
