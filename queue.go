package twitter

import (
	"fmt"
	"time"
)

// Queue struct holds information for each method, such as
// 	@rate time.Duration specific for each endpoint on Twitter
// 	@delay time.Duration fallback for @rate, specific for each endpoint on Twitter
// 	@requestsChannel chan *Request the incoming (requests) channel
// 	@responseChannel chan *Response the outgoing (response) channel
type Queue struct {
	rate            time.Duration
	delay           time.Duration
	auto            bool
	requestsChannel chan *Request
	responseChannel chan *Response
}

type QueueOption func(*Queue)

func WithRate(rate time.Duration) QueueOption {
	return func(q *Queue) {
		q.rate = rate
	}
}

func WithDelay(delay time.Duration) QueueOption {
	return func(q *Queue) {
		q.delay = delay
	}
}

func WithAuto(auto bool) QueueOption {
	return func(q *Queue) {
		q.auto = auto
	}
}

func NewQueue(rate, delay time.Duration, auto bool, in chan *Request, out chan *Response, options ...QueueOption) *Queue {
	queue := &Queue{rate, delay, auto, in, out}

	for _, o := range options {
		o(queue)
	}

	return queue
}

// processRequests, processes the incoming requests and with a build-in rate-limiter
// to avoid any rate-limit errors from Twitter API
func (q *Queue) processRequests(api *Twitter) {
	// get queue throttle
	throttle := time.Tick(q.rate)
	// listen channel
	for {
		// capture input request and channel state
		req, ok := <-q.requestsChannel
		// break the loop if the channel is closed
		if !ok {
			break
		}

		// send the request on twitter api
		err := api.apiDo(req)

		// capture request errors
		if err != nil {
			// fmt.Println(err)
			if err.Code == 429 {
				// if err == rate limit then add req to channel again and continue
				go func(c *Queue, r *Request) {
					c.requestsChannel <- r
				}(q, req)

				// get the delay
				delay := time.Tick(q.delay)

				fmt.Printf("Error %s. Delay request for %s\n", err.Message, q.delay.String())
				// delay next request for q.delay duration
				<-delay

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
		<-throttle
	}
}
