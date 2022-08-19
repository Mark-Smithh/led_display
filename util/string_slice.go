package util

import (
	"strconv"
	"strings"
)

type StrSliceImplementation struct {
	Params DisplayParams
}

func (s StrSliceImplementation) CanDisplayNumber() bool {
	numToDisplayStringSlice := strings.Split(strconv.Itoa(s.Params.NumToDisplay), "")

	totalNumLedRequired := 0
	for _, num := range numToDisplayStringSlice {
		numInt, _ := strconv.Atoi(num)
		numLedRequired := numLedsNeededToDisplay(numInt)
		totalNumLedRequired += numLedRequired
	}
	return totalNumLedRequired <= s.Params.NumLedSegments
}
