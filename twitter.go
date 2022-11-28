package twitter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
func NewTwitter(consumerKey, consumerSecret string) (*Twitter, error) {
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

// NewTwitter returns a new Twitter API v2 Client using OAuth 2.0 based authentication.
// This method is usufull when you only need to make Application-Only requests.
// Official Documentation: https://developer.twitter.com/en/docs/authentication/oauth-2-0
// Scopes: https://developer.twitter.com/en/docs/authentication/oauth-2-0/authorization-code
// func NewTwitterWithPKCE(consumerKey, consumerSecret, accessToken, accessTokenSecret string) (*Twitter, error) {
// 	// create new context
// 	ctx := context.Background()

// 	// init new Twitter client
// 	api := &Twitter{
// 		baseURL: BaseURL,
// 	}

// 	// oauth2 configures a client that uses app credentials to keep a fresh token
// 	config := &oauth2.Config{
// 		ClientID:     consumerKey,
// 		ClientSecret: consumerSecret,
// 		Scopes: []string{
// 			"tweet.read",
// 			"users.read",
// 			"offline.access",
// 		},
// 		Endpoint: oauth2.Endpoint{
// 			AuthURL:  AuthorizeTokenURL,
// 			TokenURL: TokenURL,
// 		},
// 	}

// 	// Redirect user to consent page to ask for permission
// 	// for the scopes specified above.
// 	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline,
// 		oauth2.SetAuthURLParam("code_challenge", "challenge"),
// 		oauth2.SetAuthURLParam("code_challenge_method", "plain"),
// 		oauth2.SetAuthURLParam("response_type", "code"),
// 	)
// 	fmt.Printf("Visit the URL for the auth dialog: %v", url)

// 	// Use the authorization code that is pushed to the redirect
// 	// URL. Exchange will do the handshake to retrieve the
// 	// initial access token. The HTTP Client returned by
// 	// conf.Client will refresh the token as necessary.
// 	var code string
// 	if _, err := fmt.Scan(&code); err != nil {
// 		return nil, err
// 	}
// 	tok, err := config.Exchange(ctx, code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	api.client = config.Client(ctx, tok)
// 	return api, nil
// }

// NewTwitterWithContext returns a new Twitter API v2 Client using OAuth 1.0 based authentication.
// This method is useful when you need to make API requests, on behalf of a Twitter account.
// Official Documentation: https://developer.twitter.com/en/docs/authentication/oauth-1-0a
func NewTwitterWithContext(consumerKey, consumerSecret, accessToken, accessTokenSecret string) (*Twitter, error) {
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
		return nil, err
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
func (api *Twitter) VerifyCredentials() (bool, error) {
	response, err := api.client.Get(RateLimitStatusURL)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	return err == nil, nil
}

// parseResponse returns an error while unmarshaling response body to the results interface.
func (api *Twitter) parseResponse(resp *http.Response, results *Data) error {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		return err
	}

	return nil
}

// parseResponseWithInterface
func (api *Twitter) parseResponseWithInterface(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// apiDo send's the request to Twitter API and returns an error.
// The results are processed by `parseResponse` and written to the temporary
// `req.Results` interaface.
func (api *Twitter) apiDo(req *Request) error {
	resp, err := api.client.Do(req.Req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, resp.Status))
	}

	return api.parseResponse(resp, &req.Results)
}

// apiDoWithResponse send's the request to Twitter API and returns an error.
// The results are processed by `parseResponse` and written to the temporary
// `req.Results` interaface.
func (api *Twitter) apiDoWithResponse(req *Request) ([]byte, error) {
	resp, err := api.client.Do(req.Req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, resp.Status))
	}

	return api.parseResponseWithInterface(resp)
}
