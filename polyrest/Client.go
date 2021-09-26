package polyrest

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"time"
)

const userAgent = "go-lib-polygon-io-uoc_0.0.1"

// do performs the HTTP request to Polygons API. We apply the API key and return
// a *Response containing the API response information and a standard error
func do(apiKey string,
	req *fasthttp.Request,
	resp *fasthttp.Response,
	parseOn404 bool,
	dest interface{}) (*Response, error) {

	if len(apiKey) == 0 {
		return nil, ErrAPIKeyNotSet
	}

	req.URI().QueryArgs().Set("apiKey", apiKey)
	//	req.Header.SetUserAgent(userAgent)

	var ar = new(Response)

	if debugMode == true {
		ar.uri = req.URI().String()
		log.Debug(req.URI())
	}

	var retryCount = 0

__RETRY:

	var err = fasthttp.Do(req, resp)
	if err != nil {
		return nil, err
	}

	ar.httpCode = resp.StatusCode()

	if ar.httpCode == fasthttp.StatusGatewayTimeout ||
		ar.httpCode == fasthttp.StatusBadGateway {

		if retryOn504 == true && retryCount < maxRetryCount {
			time.Sleep(retryInterval)
			retryCount = retryCount + 1
			goto __RETRY
		}

	}

	if ar.httpCode != fasthttp.StatusOK {
		if ar.httpCode == fasthttp.StatusNotFound && parseOn404 == false {
			return nil, ErrAPIReturnedError
		}
	}

	results, err := parseResponseBody(resp.Body(), ar)
	if err != nil {
		return nil, err
	}

	if len(results) != 0 {
		err = json.Unmarshal(results, dest)
		if err != nil {
			return nil, err
		}
	}

	return ar, nil
}

func parseResponseBody(src []byte, dest *Response) ([]byte, error) {

	var err = json.Unmarshal(src, dest)
	if err != nil {
		dest.rawResponse.RawResults = nil
		dest.rawResponse.RawErrorCode = nil
		return nil, err
	}

	// the "errorcode" field can either be a:
	//	- string
	// 	- int
	// 	- malformed int-string (example: "001")
	// jsoniter.Wrap("errorcode") thinks "001" is an array
	// using Get seems to do to decode for each case reliably
	if dest.IsError() == true {
		dest.errorCode = jsoniter.Get(dest.RawErrorCode).ToString()
		dest.err = errors.New(dest.ErrorString())
	}

	var rr = dest.rawResponse.RawResults
	dest.rawResponse.RawResults = nil
	dest.rawResponse.RawErrorCode = nil

	return rr, nil
}
