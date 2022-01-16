package twitter

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
)

// Request Struct
type Request struct {
	Req     *http.Request
	Results Data
}

// NewRquest returns a new Request struct
func NewRquest(method, url string, v url.Values, body []byte) (*Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	for key, value := range v {
		query.Set(key, strings.Join(value, ","))
	}

	if method == "POST" {
		// we need to set the content-type to application/json
		// to perform a post request
		request.Header.Set("Content-Type", "application/json")
	}
	request.URL.RawQuery = query.Encode()
	return &Request{request, Data{}}, nil
}

// UpdateURLValues updates request's query values
func (r *Request) UpdateURLValues(v url.Values) {
	query := r.Req.URL.Query()
	for key, value := range v {
		query.Set(key, strings.Join(value, ","))
	}
	r.Req.URL.RawQuery = query.Encode()
}

// ResetResults resets request's results
func (r *Request) ResetResults() {
	r.Results = Data{}
}
