package main

import (
	"fmt"
	"ledDisplay/util"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Only two arguments required:\n 1) The number to display on the LED.\n 2) The number of LED segments.\n")
		os.Exit(1)
	}

	argsNoProgram := os.Args[1:]

	numToDisplay, err := strconv.Atoi(argsNoProgram[0])
	errCheck(err)

	numLeds, err := strconv.Atoi(argsNoProgram[1])
	errCheck(err)

	allDisplays := []util.Display{}
	allDisplays = append(allDisplays, util.StrSliceImplementation{Params: util.DisplayParams{NumToDisplay: numToDisplay, NumLedSegments: numLeds}})
	allDisplays = append(allDisplays, util.RuneSliceImplementation{Params: util.DisplayParams{NumToDisplay: numToDisplay, NumLedSegments: numLeds}})
	allDisplays = append(allDisplays, util.ByteSliceImplementation{Params: util.DisplayParams{NumToDisplay: numToDisplay, NumLedSegments: numLeds}})

	for _, d := range allDisplays {
		if d.CanDisplayNumber() {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func errCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("number argument required")
		os.Exit(1)
	}
}
