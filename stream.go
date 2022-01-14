package twitter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/dustin/go-jsonpointer"
)

type Stream struct {
	api *Twitter
	C   chan interface{}
	run bool
}

func (stream *Stream) Stop() {
	stream.run = false
}

func (stream *Stream) start(urlStr string, v url.Values) {
	stream.run = true
	go stream.loop(urlStr, v)
}
func (stream *Stream) loop(urlStr string, v url.Values) {
	defer close(stream.C)
	for stream.run {
		request, _ := NewRquest("GET", urlStr, v)
		r, _ := stream.api.client.Do(request.Req)
		stream.listen(*r)
	}
}

func jsonToKnownType(j []byte) interface{} {
	// TODO: DRY
	var tweet StreamData
	json.Unmarshal(j, &tweet)
	return tweet
}
func (stream *Stream) listen(response http.Response) {
	if response.Body != nil {
		defer response.Body.Close()
	}

	b, err := io.ReadAll(response.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(string(b))

	scanner := bufio.NewScanner(response.Body)

	for scanner.Scan() && stream.run {
		j := scanner.Bytes()
		if len(j) == 0 {
			break
		}

		stream.C <- jsonToKnownType(j)
	}
}
func jsonAsStruct(j []byte, path string, obj interface{}) (res bool) {
	if v, _ := jsonpointer.Find(j, path); v == nil {
		return false
	}
	err := json.Unmarshal(j, obj)
	return err == nil
}

func (api Twitter) newStream(urlStr string, v url.Values) *Stream {
	stream := Stream{
		api: &api,
		C:   make(chan interface{}),
	}

	stream.start(urlStr, v)
	return &stream
}

// GetFilterStream streams tweets in real-time based on a specific set of filter rules.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 50/15m (app)
func (api *Twitter) GetFilterStream(v url.Values) *Stream {
	return api.newStream(
		fmt.Sprintf("%s/tweets/search/stream", api.baseURL), v,
	)
}

// GetFilterStreamRules returns a list of rules currently active on the streaming endpoint, either as a list or individually.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream/rules
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream-rules
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app)
func (api *Twitter) GetFilterStreamRules() {}

// PostFilterStreamRules adds or deletes rules to your stream.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream/rules
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/post-tweets-search-stream-rules
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app)
func (api *Twitter) PostFilterStreamRules(v url.Values) (*http.Response, *APIError) {
	request, _ := NewRquest("POST", fmt.Sprintf("%s/tweets/search/stream/rules", api.baseURL), v)
	return api.apiDoRest(request)
}

// GetSampleStream streams about 1% of all Tweets in real-time.
// Endpoint URL: https://api.twitter.com/2/tweets/sample/stream
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/sampled-stream/api-reference/get-tweets-sample-stream
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 50/15m (app)
func (api *Twitter) GetSampleStream() {}
