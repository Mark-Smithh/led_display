package util

import "strconv"

type RuneSliceImplementation struct {
	Params DisplayParams
}

func (r RuneSliceImplementation) CanDisplayNumber() bool {
	numToDisplayRuneSlice := []rune(strconv.Itoa(r.Params.NumToDisplay))

	totalNumLedRequired := 0
	for _, num := range numToDisplayRuneSlice {
		numInt, _ := strconv.Atoi(string(num))
		numLedRequired := numLedsNeededToDisplay(numInt)
		totalNumLedRequired += numLedRequired
	}
	return totalNumLedRequired <= r.Params.NumLedSegments
}
