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

	runParameters := util.RunImplementationParams{
		NumToDisplay: numToDisplay,
		NumLeds:      numLeds,
	}

	util.RunImplementations(runParameters)

	api := util.RestAPI{}
	api.Start()
}

func errCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("number argument required")
		os.Exit(1)
	}
}
