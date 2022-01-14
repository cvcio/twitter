package twitter

import (
	"fmt"
	"net/url"
	"time"
)

// GetTweetsSearchRecent returns Tweets from the last 7 days that match a search query.
// Endpoint URL: https://api.twitter.com/2/tweets/search/recent
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app), 180/15m (user)
func (api *Twitter) GetTweetsSearchRecent(v url.Values, options ...QueueOption) (chan *Data, chan *APIError) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/450, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan *APIError)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/tweets/search/recent", api.baseURL), v)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan *APIError, req *Request) {
		// on done close channels
		// close data channel
		defer close(d)
		// close error channel
		defer close(e)

		// listen channel
		for {
			// capture the response and channel state
			res, ok := <-q.responseChannel
			// break the loop if the channel is closed
			if !ok {
				break
			}

			// send the results to the data channel
			d <- &res.Results
			// send errors to error channel
			if res.Error != nil {
				e <- res.Error
			}

			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}

// GetTweetsSearchAll returns the complete history of public Tweets matching a search query;
// since the first Tweet was created March 26, 2006. This endpoint is part of a
// **private beta** for academic researchers only.
// *Please do not share this documentation.*
// Endpoint URL: https://api.twitter.com/2/tweets/search/all
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/full-archive-search/api-reference/get-tweets-search-all
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 1/1s (user)
func (api *Twitter) GetTweetsSearchAll(v url.Values, options ...QueueOption) (chan *Data, chan *APIError) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/300, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan *APIError)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/tweets/search/all", api.baseURL), v)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan *APIError, req *Request) {
		// on done close channels
		// close data channel
		defer close(d)
		// close error channel
		defer close(e)

		// listen channel
		for {
			// capture the response and channel state
			res, ok := <-q.responseChannel
			// break the loop if the channel is closed
			if !ok {
				break
			}

			// send the results to the data channel
			d <- &res.Results
			// send errors to error channel
			if res.Error != nil {
				e <- res.Error
			}

			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}
