package twitter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Stream struct {
	api *Twitter
	C   chan interface{}
	run bool
}

func (stream *Stream) Stop() {
	stream.run = false
}

func (stream *Stream) start(urlStr string, v url.Values) *APIError {
	stream.run = true

	request, err := NewRquest("GET", urlStr, v, nil)
	if err != nil {
		return &APIError{0, err.Error()}
	}
	r, err := stream.api.client.Do(request.Req)
	if err != nil {
		return &APIError{0, err.Error()}
	}

	go stream.listen(r)

	return nil
}

// func (stream *Stream) loop(urlStr string, v url.Values) {
// 	defer close(stream.C)
// 	for stream.run {
// 		request, err := NewRquest("GET", urlStr, v, nil)
// 		if err != nil {
// 			panic(err)
// 		}
// 		r, err := stream.api.client.Do(request.Req)
// 		if err != nil {
// 			panic(err)
// 		}
// 		stream.listen(r)
// 	}
// }

func jsonToKnownType(j []byte) interface{} {
	var tweet StreamData
	json.Unmarshal(j, &tweet)
	return tweet
}

func (stream *Stream) listen(response *http.Response) {
	if response.Body != nil {
		defer response.Body.Close()
	}
	defer close(stream.C)

	// created the scanner to read each line
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() && stream.run {
		line := scanner.Bytes()

		// Contuinue if empty bytes returned from the stream
		if len(line) == 0 {
			continue
		}

		// stream.C <- jsonToKnownType(line)
		stream.C <- jsonToKnownType(bytes.TrimRight(line, "\r\n"))
	}
}

func (api Twitter) newStream(urlStr string, v url.Values) (*Stream, *APIError) {
	stream := Stream{
		api: &api,
		C:   make(chan interface{}),
	}

	err := stream.start(urlStr, v)
	if err != nil {
		return nil, err
	}
	return &stream, nil
}

// GetFilterStream streams tweets in real-time based on a specific set of filter rules.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 50/15m (app)
func (api *Twitter) GetFilterStream(v url.Values) (*Stream, *APIError) {
	return api.newStream(
		fmt.Sprintf("%s/tweets/search/stream", api.baseURL), v,
	)
}

// GetFilterStreamRules returns a list of rules currently active on the streaming endpoint, either as a list or individually.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream/rules
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream-rules
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app)
func (api *Twitter) GetFilterStreamRules(v url.Values) (*Rules, *APIError) {
	request, _ := NewRquest("GET", fmt.Sprintf("%s/tweets/search/stream/rules", api.baseURL), v, nil)

	res, err := api.apiDoWithResponse(request)
	if err != nil {
		return nil, err
	}

	rules := new(Rules)

	if err := json.Unmarshal(res, &rules); err != nil {
		return nil, &APIError{0, err.Error()}
	}

	return rules, nil
}

// PostFilterStreamRules adds or deletes rules to your stream.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream/rules
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/post-tweets-search-stream-rules
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app)
func (api *Twitter) PostFilterStreamRules(v url.Values, body []byte) (*Rules, *APIError) {
	request, _ := NewRquest("POST", fmt.Sprintf("%s/tweets/search/stream/rules", api.baseURL), v, body)

	res, err := api.apiDoWithResponse(request)
	if err != nil {
		return nil, err
	}

	rules := new(Rules)

	if err := json.Unmarshal(res, &rules); err != nil {
		return nil, &APIError{0, err.Error()}
	}

	return rules, nil
}

// GetSampleStream streams about 1% of all Tweets in real-time.
// Endpoint URL: https://api.twitter.com/2/tweets/sample/stream
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/sampled-stream/api-reference/get-tweets-sample-stream
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 50/15m (app)
func (api *Twitter) GetSampleStream(v url.Values) (*Stream, *APIError) {
	return api.newStream(
		fmt.Sprintf("%s/tweets/sample/stream", api.baseURL), v,
	)
}
