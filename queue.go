package twitter

import (
	"time"
)

// Queue struct holds information for each method, such as
// @rate time.Duration specific for each endpoint on Twitter
// @delay time.Duration fallback for @rate, specific for each endpoint on Twitter
// @requestsChannel chan *Request the incoming (requests) channel
// @responseChannel chan *Response the outgoing (response) channel
type Queue struct {
	rate            time.Duration
	delay           time.Duration
	auto            bool
	closeChannels   bool
	requestsChannel chan *Request
	responseChannel chan *Response
}

// QueueOption queue options struct
type QueueOption func(*Queue)

// WithRate (default: according to endpoint) adjusts the the duration between each request to avoid
// rate limits from Twitter API
func WithRate(rate time.Duration) QueueOption {
	return func(q *Queue) {
		q.rate = rate
	}
}

// WithDelay (default:15 minutes) adjusts the the duration between each errored requests
// due to rate limit errors from Twitter API
func WithDelay(delay time.Duration) QueueOption {
	return func(q *Queue) {
		q.delay = delay
	}
}

// WithAuto (default:true) will auto continue to the next page if
// pagination_token exists in the response object
func WithAuto(auto bool) QueueOption {
	return func(q *Queue) {
		q.auto = auto
	}
}

// NewQueue creates a new queue
func NewQueue(rate, delay time.Duration, auto bool, in chan *Request, out chan *Response, options ...QueueOption) *Queue {
	queue := &Queue{rate, delay, auto, true, in, out}

	for _, o := range options {
		o(queue)
	}

	return queue
}

// processRequests, processes the incoming requests and with a build-in rate-limiter
// to avoid any rate-limit errors from Twitter API
func (q *Queue) processRequests(api *Twitter) {
	// close queue's channels
	defer q.Close()

	// listen channel
	for {
		// capture input request and channel state
		req := <-q.requestsChannel

		// break the loop if the channel is closed
		// if !ok {
		// 	break
		// }

		// send the request on twitter api
		err := api.apiDo(req)

		// capture request errors
		if err != nil && q.auto {
			if err.Code == 420 || err.Code == 429 || err.Code >= 500 {
				// if err == rate limit then add req to channel again and continue
				go func(c *Queue, r *Request) {
					c.requestsChannel <- r
				}(q, req)

				// delay next request for q.delay duration
				<-time.After(q.delay)

				// go to start
				continue
			} else {
				q.responseChannel <- &Response{req.Results, err}
				break
			}
		}

		// add response to channel
		q.responseChannel <- &Response{req.Results, err}

		// throttle requests to avoid rate-limit errors
		<-time.After(q.rate)
	}
}

// Close closes requests and response channels
func (q *Queue) Close() {
	close(q.requestsChannel)
	close(q.responseChannel)
}
