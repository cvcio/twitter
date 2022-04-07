package twitter

import (
	"fmt"
	"net/url"
	"time"
)

// GetUserFollowers returns a list of users who are followers of the specified user ID.
// Endpoint URL: https://api.twitter.com/2/users/:id/followers
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 15/15m (app), 15/15m (user)
func (api *Twitter) GetUserFollowers(id string, v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/15, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/users/%s/followers", api.baseURL, id), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
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

			// if there is a next page, transform the original request object
			// by setting the `pagination_token` parameter to get the next page
			if res.Results.Meta.NextToken != "" && q.auto {
				// create new url values and add the pagination token
				nv := url.Values{}
				nv.Add("pagination_token", res.Results.Meta.NextToken)

				// update request's url Values
				req.UpdateURLValues(nv)
				// reset request's results
				req.ResetResults()

				// add next request to the channel
				q.requestsChannel <- req

				//go to start
				continue
			}
			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}

// GetUserFollowing returns a list of users the specified user ID is following.
// Endpoint URL: https://api.twitter.com/2/users/:id/following
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 15/15m (app), 15/15m (user)
func (api *Twitter) GetUserFollowing(id string, v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/15, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/users/%s/following", api.baseURL, id), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
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

			// if there is a next page, transform the original request object
			// by setting the `pagination_token` parameter to get the next page
			if res.Results.Meta.NextToken != "" && q.auto {
				// create new url values and add the pagination token
				nv := url.Values{}
				nv.Add("pagination_token", res.Results.Meta.NextToken)

				// update request's url Values
				req.UpdateURLValues(nv)
				// reset request's results
				req.ResetResults()

				// add next request to the channel
				q.requestsChannel <- req

				//go to start
				continue
			}
			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}

// GetUsers returns a variety of information about one or more users specified by the requested IDs.
// Endpoint URL: https://api.twitter.com/2/users
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 900/15m (user)
func (api *Twitter) GetUsers(v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/15, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/users", api.baseURL), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
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

// GetUsersByUserName returns a variety of information about one or more users specified by their usernames.
// Endpoint URL: https://api.twitter.com/2/users/by
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 900/15m (user)
func (api *Twitter) GetUsersBy(v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/15, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/users/by", api.baseURL), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
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

// GetUserByID returns a variety of information about a single user specified by the requested ID.
// Endpoint URL: https://api.twitter.com/2/users/:id
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 900/15m (user)
func (api *Twitter) GetUserByID(id string, v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/15, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/users/%s", api.baseURL, id), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
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

// GetUserByUserName returns a variety of information about one or more users specified by their usernames.
// Endpoint URL: https://api.twitter.com/2/users/by/username/:username
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 900/15m (user)
func (api *Twitter) GetUsersByUserName(username string, v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/15, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/users/by/username/%s", api.baseURL, username), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
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
