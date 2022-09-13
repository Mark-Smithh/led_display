package util

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/patrickmn/go-cache"
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

type RunImplementationsResp struct {
	CanDisplayNumber bool
	CacheHit         bool
}

func RunImplementations(params RunImplementationParams) RunImplementationsResp {
	allDisplays := []Display{}
	allDisplays = append(allDisplays, StrSliceImplementation{Params: DisplayParams{NumToDisplay: params.NumToDisplay, NumLedSegments: params.NumLeds}})
	allDisplays = append(allDisplays, RuneSliceImplementation{Params: DisplayParams{NumToDisplay: params.NumToDisplay, NumLedSegments: params.NumLeds}})
	allDisplays = append(allDisplays, ByteSliceImplementation{Params: DisplayParams{NumToDisplay: params.NumToDisplay, NumLedSegments: params.NumLeds}})

	canDisplayNumberComputedValue := true
	foundInCache := false

	if params.HttpResponseWriter == nil {
		for _, d := range allDisplays {
			if d.CanDisplayNumber() {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
				canDisplayNumberComputedValue = false
			}
		}
		return RunImplementationsResp{CanDisplayNumber: canDisplayNumberComputedValue, CacheHit: foundInCache}
	}

	cacheKey := strconv.Itoa(params.NumToDisplay) + "-" + strconv.Itoa(params.NumLeds)

	canDisplayNumberCacheValue, foundInCache := apiCache.Get(cacheKey)
	if foundInCache {
		response := apiResponse{
			Message: strconv.FormatBool(canDisplayNumberCacheValue.(bool)),
		}
		encodeResponse(response, params.HttpResponseWriter)
		return RunImplementationsResp{CanDisplayNumber: canDisplayNumberCacheValue.(bool), CacheHit: foundInCache}
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
			canDisplayNumberComputedValue = false
			encodeResponse(response, params.HttpResponseWriter)
		}
	}

	fmt.Printf("adding cache key %s\n", cacheKey)
	apiCache.Set(cacheKey, canDisplayNumberComputedValue, cache.NoExpiration)
	return RunImplementationsResp{CanDisplayNumber: canDisplayNumberComputedValue, CacheHit: foundInCache}
}
