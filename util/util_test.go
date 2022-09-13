package util

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestUcase(t *testing.T) {
	input := "hello led"
	expected := "HELLO LED"
	transformedInput := Ucase(input)
	assert.Equal(t, transformedInput, expected)
}

func TestRunImplementationsConsoleFalse(t *testing.T) {
	params := RunImplementationParams{
		NumToDisplay: 1,
		NumLeds:      0,
	}

	resp := RunImplementations(params)
	assert.False(t, resp.CanDisplayNumber, "expected CanDisplayNumber to be false.")
}

func TestRunImplementationsConsoleTrue(t *testing.T) {
	params := RunImplementationParams{
		NumToDisplay: 1,
		NumLeds:      2,
	}

	resp := RunImplementations(params)
	assert.True(t, resp.CanDisplayNumber, "expected CanDisplayNumber to be true.")
}

func TestRunImplementationsApiTrue(t *testing.T) {
	apiCache = cache.New(5*time.Minute, 10*time.Minute)
	writer := httptest.NewRecorder()
	params := RunImplementationParams{
		NumToDisplay:       1,
		NumLeds:            2,
		HttpResponseWriter: writer,
	}

	resp := RunImplementations(params)
	assert.True(t, resp.CanDisplayNumber, "expected CanDisplayNumber to be true.")
}

func TestRunImplementationsApiFalse(t *testing.T) {
	apiCache = cache.New(5*time.Minute, 10*time.Minute)
	writer := httptest.NewRecorder()
	params := RunImplementationParams{
		NumToDisplay:       1,
		NumLeds:            0,
		HttpResponseWriter: writer,
	}

	resp := RunImplementations(params)
	assert.False(t, resp.CanDisplayNumber, "expected CanDisplayNumber to be false.")
}

func TestRunImplementationsApiTrueCacheHit(t *testing.T) {
	apiCache = cache.New(5*time.Minute, 10*time.Minute)
	writer := httptest.NewRecorder()
	params := RunImplementationParams{
		NumToDisplay:       1,
		NumLeds:            2,
		HttpResponseWriter: writer,
	}

	resp := RunImplementations(params)
	assert.False(t, resp.CacheHit, "expected CacheHit to be false.")

	resp = RunImplementations(params)
	assert.True(t, resp.CacheHit, "expected CacheHit to be true.")
}
