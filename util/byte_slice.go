package util

import "strconv"

type ByteSliceImplementation struct {
	Params DisplayParams
}

func (b ByteSliceImplementation) CanDisplayNumber() bool {
	numToDisplayByteSlice := []byte(strconv.Itoa(b.Params.NumToDisplay))

	totalNumLedRequired := 0
	for _, num := range numToDisplayByteSlice {
		numInt, _ := strconv.Atoi(string(num))
		numLedRequired := numLedsNeededToDisplay(numInt)
		totalNumLedRequired += numLedRequired
	}
	return totalNumLedRequired <= b.Params.NumLedSegments
}
