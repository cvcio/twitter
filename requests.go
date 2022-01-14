package twitter

import (
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
func NewRquest(method, url string, v url.Values) (*Request, error) {
	request, err := http.NewRequest(method, url, nil)
	query := request.URL.Query()
	for key, value := range v {
		query.Set(key, strings.Join(value, ","))
	}
	request.URL.RawQuery = query.Encode()
	return &Request{request, Data{}}, err
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
