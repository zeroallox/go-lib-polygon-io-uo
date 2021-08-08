package polyrest

import (
	"encoding/json"
	"fmt"
	"math"
)

type Result struct {
	rez result
}

type result struct {
	Results     json.RawMessage `json:"results"`
	Status      string          `json:"status"`
	Ticker      string          `json:"ticker"`
	Success     bool            `json:"success"`
	RequestID   string          `json:"request_id"`
	Count       int             `json:"count"`
	ResultCount int             `json:"result_count"`
	Error       string          `json:"error"`
	NextURL     string          `json:"next_url"`
	DBLatency   int             `json:"db_latency"`
	HTTPCode    int
}

func (r *Result) IsOK() bool {
	return r.rez.Status == "OK"
}

func (r *Result) IsError() bool {
	return r.rez.Status == "ERROR" || r.rez.Success == false
}

func (r *Result) Error() string {
	return r.rez.Error
}

func (r *Result) Status() string {
	return r.rez.Status
}

func (r *Result) HTTPStatusCode() int {
	return r.rez.HTTPCode
}

func (r *Result) RequestID() string {
	return r.rez.RequestID
}

func (r *Result) DBLatency() int {
	return r.rez.DBLatency
}

func (r *Result) Count() uint {
	return uint(math.Max(float64(r.rez.Count), float64(r.rez.ResultCount)))
}

func (r *Result) NextURL() string {
	return r.rez.NextURL
}

func (r *Result) Ticker() string {
	return r.rez.Ticker
}

func (r *Result) String() string {
	return fmt.Sprintf("%v", r.rez)
}

func (r *Result) takeResultData() json.RawMessage {
	var data = r.rez.Results
	r.rez.Results = nil
	return data
}
