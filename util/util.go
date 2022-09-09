package util

import (
	"fmt"
	"net/http"
	"strings"
)

func Ucase(input string) string {
	return strings.ToUpper(input)
}

type RunImplementationParams struct {
	NumToDisplay       int
	NumLeds            int
	HttpResponseWriter http.ResponseWriter
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

func RunImplementations(params RunImplementationParams) {
	allDisplays := []Display{}
	allDisplays = append(allDisplays, StrSliceImplementation{Params: DisplayParams{NumToDisplay: params.NumToDisplay, NumLedSegments: params.NumLeds}})
	allDisplays = append(allDisplays, RuneSliceImplementation{Params: DisplayParams{NumToDisplay: params.NumToDisplay, NumLedSegments: params.NumLeds}})
	allDisplays = append(allDisplays, ByteSliceImplementation{Params: DisplayParams{NumToDisplay: params.NumToDisplay, NumLedSegments: params.NumLeds}})

	if params.HttpResponseWriter == nil {
		for _, d := range allDisplays {
			if d.CanDisplayNumber() {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
		return
	}

	for _, d := range allDisplays {
		if d.CanDisplayNumber() {
			response := apiResponse{
				Message: "true",
			}
			encodeResponse(response, params.HttpResponseWriter)
		} else {
			response := apiResponse{
				Message: "false",
			}
			encodeResponse(response, params.HttpResponseWriter)
		}
	}
}
