package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUcase(t *testing.T) {
	input := "hello led"
	expected := "HELLO LED"
	transformedInput := Ucase(input)
	assert.Equal(t, transformedInput, expected)
}
