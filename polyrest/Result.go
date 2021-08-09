package polyrest

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"math"
)

type APIResponse struct {
	ar apiResponse
}

type apiResponse struct {
	Results      jsoniter.RawMessage `json:"results"`
	Status       string              `json:"status"`
	Ticker       string              `json:"ticker"`
	Success      bool                `json:"success"`
	RequestID    string              `json:"request_id"`
	Count        int                 `json:"count"`
	ResultCount  int                 `json:"result_count"`
	Error        string              `json:"error"`
	ErrorCode    string
	NextURL      string `json:"next_url"`
	DBLatency    int    `json:"db_latency"`
	HTTPCode     int
	ErrorCodeRaw jsoniter.RawMessage `json:"errorcode"`
	URI          string
}

// URI the URI of the request that made the API call. Useful for debugging
// and only populated if debug is enabled.
func (r *APIResponse) URI() string {
	return r.ar.URI
}

// IsOK returns if the call is ok / success / otherwise happy.
func (r *APIResponse) IsOK() bool {
	return r.ar.Status == "OK"
}

// Success returns if the API reported the call as a success.
func (r *APIResponse) Success() bool {
	return r.ar.Success
}

// IsError returns if the response is an API-level (not library level) error.
func (r *APIResponse) IsError() bool {
	return r.ar.Status == "ERROR"
}

// Error returns the error string if the API provided one.
func (r *APIResponse) Error() string {
	return r.ar.Error
}

// ErrorCode returns the ErrorCode provided by the API call if one exists.
func (r *APIResponse) ErrorCode() string {
	return r.ar.ErrorCode
}

// Status returns the Status if the API provided one.
func (r *APIResponse) Status() string {
	return r.ar.Status
}

// HTTPStatusCode returns the http code of the API call.
func (r *APIResponse) HTTPStatusCode() int {
	return r.ar.HTTPCode
}

// RequestID returns the request ID if the API provided one.
func (r *APIResponse) RequestID() string {
	return r.ar.RequestID
}

// DBLatency returns the database latency if the API provided it.
func (r *APIResponse) DBLatency() int {
	return r.ar.DBLatency
}

// Count returns the result count if the API provided it.
func (r *APIResponse) Count() uint {
	return uint(math.Max(float64(r.ar.Count), float64(r.ar.ResultCount)))
}

// NextURL returns the next URL provided by the new pagination API.
func (r *APIResponse) NextURL() string {
	return r.ar.NextURL
}

// Ticker returns the ticker associated with the call if one was provided.
func (r *APIResponse) Ticker() string {
	return r.ar.Ticker
}

// String dumps APIResponse as a string.
func (r *APIResponse) String() string {
	return fmt.Sprintf("%v", r.ar)
}
