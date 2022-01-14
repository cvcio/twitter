package twitter

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mrjones/oauth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Constants
const (
	BaseURL            = "https://api.twitter.com/2"
	RequestTokenURL    = "https://api.twitter.com/oauth/request_token"
	AuthorizeTokenURL  = "https://api.twitter.com/oauth/authorize"
	AccessTokenURL     = "https://api.twitter.com/oauth/access_token"
	TokenURL           = "https://api.twitter.com/oauth2/token"
	RateLimitStatusURL = "https://api.twitter.com/1.1/application/rate_limit_status.json"
)

// Twitter API Client
type Twitter struct {
	client  *http.Client
	baseURL string
	queue   *Queue
}

// NewTwitter returns a new Twitter API v2 Client using OAuth 2.0 based authentication.
// This method is usufull when you only need to make Application-Only requests.
// Official Documentation: https://developer.twitter.com/en/docs/authentication/oauth-2-0
func NewTwitter(consumerKey, consumerSecret string) (*Twitter, *APIError) {
	// create new context
	ctx := context.Background()

	// init new Twitter client
	api := &Twitter{
		baseURL: BaseURL,
	}

	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     TokenURL,
	}

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 30 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	// http.Client will automatically authorize Requests
	api.client = config.Client(ctx)
	return api, nil
}

// NewTwitterWithContext returns a new Twitter API v2 Client using OAuth 1.0 based authentication.
// This method is usefull when you need to make API requests, on behalf of a Twitter account.
// Official Documentation: https://developer.twitter.com/en/docs/authentication/oauth-1-0a
func NewTwitterWithContext(consumerKey, consumerSecret, accessToken, accessTokenSecret string) (*Twitter, *APIError) {
	// init new Twitter client
	api := &Twitter{
		baseURL: BaseURL,
	}

	// create the consumer
	oauthConsumer := oauth.NewConsumer(consumerKey, consumerSecret, oauth.ServiceProvider{
		RequestTokenUrl:   RequestTokenURL,
		AuthorizeTokenUrl: AuthorizeTokenURL,
		AccessTokenUrl:    AccessTokenURL,
	})

	//set tokens
	oauthToken := oauth.AccessToken{
		Token:  accessToken,
		Secret: accessTokenSecret,
	}

	// Use the custom HTTP client with tokens
	httpClient, err := oauthConsumer.MakeHttpClient(&oauthToken)
	if err != nil {
		return nil, &APIError{0, err.Error()}
	}

	// http.Client will automatically authorize Requests
	api.client = httpClient
	return api, nil
}

// GetClient Get HTTP Client
func (api *Twitter) GetClient() *http.Client {
	return api.client
}

// VerifyCredentials returns bool upon successful request. This method will make a request
// on the rate-limit endpoint since there is no official token validation method.
func (api *Twitter) VerifyCredentials() (bool, *APIError) {
	response, err := api.client.Get(RateLimitStatusURL)
	if err != nil {
		return false, &APIError{0, err.Error()}
	}
	defer response.Body.Close()
	return err == nil, nil
}

// parseResponse returns an error while unmarshaling response body to the results interface.
func parseResponse(resp *http.Response, results *Data) *APIError {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &APIError{0, err.Error()}
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		return &APIError{0, err.Error()}
	}
	return nil
}

// apiDo send's the request to Twitter API and returns an error.
// The results are processed by `parseResponse` and written to the temporary
// `req.Results` interaface.
func (api *Twitter) apiDo(req *Request) *APIError {
	// fmt.Println(req.Req.URL.String())
	resp, err := api.client.Do(req.Req)
	if err != nil {
		return &APIError{0, err.Error()}
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return &APIError{resp.StatusCode, resp.Status}
	}

	return parseResponse(resp, &req.Results)
}

// apiDoRest send's the request to Twitter API and returns an error.
// The results are processed by `parseResponse` and written to the temporary
// `req.Results` interaface.
func (api *Twitter) apiDoRest(req *Request) (*http.Response, *APIError) {
	// fmt.Println(req.Req.URL.String())
	resp, err := api.client.Do(req.Req)
	if err != nil {
		return nil, &APIError{0, err.Error()}
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, &APIError{resp.StatusCode, resp.Status}
	}

	return resp, nil
}
