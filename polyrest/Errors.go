package polyrest

import "errors"

var ErrAPIReturnedError = errors.New("polygon api returned an error")
var ErrInvalidDateParam = errors.New("invalid date")

// ErrNot200 is returned when Polygon sends back a non 200 status code and also didn't
// send back error information from the API. An example of this would be if the API
// itself was broken.
var ErrNot200 = errors.New("status code not 200")

// ErrAPIKeyNotSet is returned when an API call was made with an empty API key.
var ErrAPIKeyNotSet = errors.New("api key missing")
