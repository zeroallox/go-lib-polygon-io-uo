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
