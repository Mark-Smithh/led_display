package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

var apiCache *cache.Cache

type API interface {
	Start()
	MsgBegin()
}

type apiResponse struct {
	Message string `json:"message"`
}

type RestAPI struct {
	Display Display
	Result  bool
}

type missingQueryStringParam struct {
	Message string `json:"message"`
}

func (r RestAPI) Start() {
	apiCache = cache.New(5*time.Minute, 10*time.Minute)
	http.HandleFunc("/candisplay", Answer)
	http.HandleFunc("/", r.MsgBegin) //localhost:8080
	fmt.Println("Starting http server to listen on port 8080.  http://localhost:8080/candisplay?num_display=[number]&num_led=[number")
	http.ListenAndServe(":8080", nil)
}

func (r RestAPI) MsgBegin(w http.ResponseWriter, req *http.Request) {
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetEscapeHTML(false) // needed to not escape problematic HTML characters should be inside JSON quoted strings.
	m := apiResponse{
		Message: msgBegin(),
	}
	jsonEncoder.Encode(m)
}

func msgBegin() string {
	return "Only two arguments required: 1) The number to display on the LED. 2) The number of LED segments."
}

func Answer(responseWriter http.ResponseWriter, req *http.Request) {
	found, numToDisplay := paramCheck(req, responseWriter, "num_display", "The number to display on the LED is required.  Example localhost:8080/canDisplay?num_display=5")
	if !found {
		return
	}

	found, numLeds := paramCheck(req, responseWriter, "num_led", "The number of LED segments is required.  Example localhost:8080/canDisplay?num_led=7")
	if !found {
		return
	}

	numToDisplayInt, _ := strconv.Atoi(numToDisplay)
	numToLedsInt, _ := strconv.Atoi(numLeds)

	runParameters := RunImplementationParams{
		NumToDisplay:       numToDisplayInt,
		NumLeds:            numToLedsInt,
		HttpResponseWriter: responseWriter,
	}

	RunImplementations(runParameters)
}

func paramCheck(req *http.Request, responseWriter http.ResponseWriter, param string, msg string) (bool, string) {
	enc := json.NewEncoder(responseWriter)
	enc.SetEscapeHTML(false) // needed to not escape problematic HTML characters should be inside JSON quoted strings.  added to not change & to \u0026
	v, ok := req.URL.Query()[param]
	if !ok {
		missingParam := missingQueryStringParam{
			Message: msg,
		}
		encodeResponse(apiResponse(missingParam), responseWriter)
		return false, ""
	}
	return true, v[0]
}

func encodeResponse(response apiResponse, responseWriter http.ResponseWriter) {
	jsonEncoder := json.NewEncoder(responseWriter)
	jsonEncoder.SetEscapeHTML(false) // needed to not escape problematic HTML characters should be inside JSON quoted strings.
	err := jsonEncoder.Encode(response)
	if err != nil {
		response.Message = err.Error()
		encodeResponse(response, responseWriter)
	}
}
