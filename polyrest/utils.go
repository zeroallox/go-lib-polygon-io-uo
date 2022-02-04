package polyrest

import (
    "github.com/google/go-querystring/query"
    "github.com/valyala/fasthttp"
)

func buildRequest(req *fasthttp.Request, uri string, params interface{}) error {
    req.SetRequestURI(uri)

    v, err := query.Values(params)
    if err != nil {
        return err
    }
    req.URI().SetQueryString(v.Encode())

    return nil
}
