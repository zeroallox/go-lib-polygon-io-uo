package polyrest

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// do performs the HTTP request to Polygons API. We apply the API key and return
// a *Result containing the API response information and a standard error
func do(apiKey string, req *fasthttp.Request, resp *fasthttp.Response) (*Result, error) {

	req.URI().QueryArgs().Set("apiKey", apiKey)
	var err = fasthttp.Do(req, resp)
	if err != nil {
		return nil, err
	}

	var rez = new(Result)
	err = json.Unmarshal(resp.Body(), &rez.rez)
	if err != nil {
		return nil, err
	}

	rez.rez.HTTPCode = resp.StatusCode()

	if rez.IsError() == true {
		return rez, ErrAPIReturnedError
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return rez, ErrNot200
	}

	return rez, nil
}
