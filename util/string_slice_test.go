package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanDisplayNumberStringFalse(t *testing.T) {
	imp := StrSliceImplementation{DisplayParams{NumToDisplay: 1, NumLedSegments: 0}}
	result := imp.CanDisplayNumber()

	assert.False(t, result, "expected result to equal false")
}

func TestCanDisplayNumberStringTrue(t *testing.T) {
	imp := StrSliceImplementation{DisplayParams{NumToDisplay: 1, NumLedSegments: 10}}
	result := imp.CanDisplayNumber()

	assert.True(t, result, "expected result to equal true")
}
