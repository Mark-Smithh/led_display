package util

type Display interface {
	CanDisplayNumber() bool
}

type DisplayParams struct {
	NumToDisplay   int
	NumLedSegments int
}
