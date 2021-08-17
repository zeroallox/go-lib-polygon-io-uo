package polyrest

import (
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"time"
)

const userAgent = "go-lib-polygon-io-uoc_0.0.1"

// do performs the HTTP request to Polygons API. We apply the API key and return
// a *APIResponse containing the API response information and a standard error
func do(apiKey string,
	req *fasthttp.Request,
	resp *fasthttp.Response,
	result interface{},
	parseOn404 bool) (*APIResponse, error) {

	if len(apiKey) == 0 {
		return nil, ErrAPIKeyNotSet
	}

	req.URI().QueryArgs().Set("apiKey", apiKey)
	//	req.Header.SetUserAgent(userAgent)

	var ar = new(APIResponse)
	defer func() {
		ar.ar.Results = nil
		ar.ar.ErrorCodeRaw = nil
	}()

	if debugMode == true {
		ar.ar.URI = req.URI().String()
		log.Debug(req.URI())
	}

	var retryCount = 0

__RETRY:

	var err = fasthttp.Do(req, resp)
	if err != nil {
		return nil, err
	}

	ar.ar.HTTPCode = resp.StatusCode()

	if ar.HTTPStatusCode() == fasthttp.StatusGatewayTimeout ||
		ar.HTTPStatusCode() == fasthttp.StatusBadGateway {

		if retryOn504 == true && retryCount < maxRetryCount {
			time.Sleep(retryInterval)
			retryCount = retryCount + 1
			goto __RETRY
		}

	}

	if ar.HTTPStatusCode() != fasthttp.StatusOK {
		if ar.HTTPStatusCode() == fasthttp.StatusNotFound && parseOn404 == false {
			return ar, ErrAPIReturnedError
		}
	}

	err = json.Unmarshal(resp.Body(), &ar.ar)
	if err != nil {
		return ar, err
	}

	// Sometimes an ErrorCode is sometimes returned sometimes from the API
	// sometimes as a string and sometimes as an int sometimes. If it exists we
	// wrap it as a string regardless alltimes.
	ar.ar.ErrorCode = jsoniter.Wrap(ar.ar.ErrorCodeRaw).ToString()

	//log.Println("Error", ar.Error())
	//log.Println("IsError", ar.IsError())
	//log.Println("OK", ar.IsOK())
	//log.Println("Status", ar.Status())

	if ar.IsError() == true {
		return ar, ErrAPIReturnedError
	}

	if ar.ar.Results != nil {
		err = json.Unmarshal(ar.ar.Results, result)
		if err != nil {
			return ar, err
		}
	}

	return ar, nil
}
