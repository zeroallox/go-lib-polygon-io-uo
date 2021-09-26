package polyrest

import (
	jsoniter "github.com/json-iterator/go"
	"math"
)

type Response struct {
	rawResponse

	httpCode  int
	uri       string
	err       error
	errorCode string
}

func (ar *Response) IsError() bool {
	return ar.rawResponse.Status == "ERROR" || len(ar.rawResponse.ErrorString) != 0
}

func (ar *Response) ErrorCode() string {
	return ar.errorCode
}

func (ar *Response) Error() error {
	return ar.err
}

func (ar *Response) ErrorString() string {
	return ar.rawResponse.ErrorString
}

func (ar *Response) HttpCode() int {
	return ar.httpCode
}

func (ar *Response) URI() string {
	return ar.uri
}

func (ar *Response) Status() string {
	return ar.rawResponse.Status
}

func (ar *Response) Ticker() string {
	return ar.rawResponse.Ticker
}

func (ar *Response) Success() bool {
	return ar.rawResponse.Success
}

func (ar *Response) RequestID() string {
	return ar.rawResponse.RequestID
}

func (ar *Response) ResultCount() int {
	return int(math.Max(float64(ar.rawResponse.Count), float64(ar.rawResponse.ResultCount)))
}

func (ar *Response) NextURL() string {
	return ar.rawResponse.NextURL
}

func (ar *Response) DBLatency() int {
	return ar.rawResponse.DBLatency
}

type rawResponse struct {
	Status      string `json:"status"`
	Ticker      string `json:"ticker"`
	Success     bool   `json:"success"`
	RequestID   string `json:"request_id"`
	Count       int    `json:"count"`
	ResultCount int    `json:"result_count"`
	NextURL     string `json:"next_url"`
	DBLatency   int    `json:"db_latency"`
	ErrorString string `json:"error"`

	RawResults   jsoniter.RawMessage `json:"results"`
	RawErrorCode jsoniter.RawMessage `json:"errorcode"`
}

//// URI the URI of the request that made the API call. Useful for debugging
//// and only populated if debug is enabled.
//func (r *Response) URI() string {
//	return r.raw.URI
//}
//
//// IsOK returns if the call is ok / success / otherwise happy.
//func (r *Response) IsOK() bool {
//	return r.raw.Status == "OK"
//}
//
//// Success returns if the API reported the call as a success.
//func (r *Response) Success() bool {
//	return r.raw.Success
//}
//
//// IsError returns if the response is an API-level (not library level) error.
//func (r *Response) IsError() bool {
//	return r.raw.Status == "ERROR"
//}
//
//// ErrorString returns the Response as an error.
//func (r *Response) Error() error {
//	return r.raw.err
//}
//
//// ErrorString returns the error string if the API provided one.
//func (r *Response) ErrorString() string {
//	return r.raw.err.Error()
//}
//
//// ErrorCode returns the ErrorCode provided by the API call if one exists.
//func (r *Response) ErrorCode() string {
//	return r.raw.Ticker
//}
//
//// Status returns the Status if the API provided one.
//func (r *Response) Status() string {
//	return r.raw.Status
//}
//
//// HTTPStatusCode returns the http code of the API call.
//func (r *Response) HTTPStatusCode() int {
//	return r.raw.HTTPCode
//}
//
//// RequestID returns the request ID if the API provided one.
//func (r *Response) RequestID() string {
//	return r.raw.RequestID
//}
//
//// DBLatency returns the database latency if the API provided it.
//func (r *Response) DBLatency() int {
//	return r.raw.DBLatency
//}
//
//// Count returns the result count if the API provided it.
//func (r *Response) Count() uint {
//	return uint(math.Max(float64(r.raw.Count), float64(r.raw.ResultCount)))
//}
//
//// NextURL returns the next URL provided by the new pagination API.
//func (r *Response) NextURL() string {
//	return r.raw.NextURL
//}
//
//// Ticker returns the ticker associated with the call if one was provided.
//func (r *Response) Ticker() string {
//	return r.raw.Ticker
//}
//
//// String dumps Response as a string.
//func (r *Response) String() string {
//	return fmt.Sprintf("%v", r.raw)
//}
