package util

import (
	"strings"
)

func Ucase(input string) string {
	return strings.ToUpper(input)
}

type Display interface {
	CanDisplayNumber() bool
}

type DisplayParams struct {
	NumToDisplay   int
	NumLedSegments int
}

func numLedsNeededToDisplay(numToDisplay int) int {
	ledsRequiredToDisplay := make(map[int]int)
	ledsRequiredToDisplay[1] = 2
	ledsRequiredToDisplay[2] = 5
	ledsRequiredToDisplay[3] = 5
	ledsRequiredToDisplay[4] = 4
	ledsRequiredToDisplay[5] = 5
	ledsRequiredToDisplay[6] = 6
	ledsRequiredToDisplay[7] = 3
	ledsRequiredToDisplay[8] = 7
	ledsRequiredToDisplay[9] = 9

	return ledsRequiredToDisplay[numToDisplay]
}
