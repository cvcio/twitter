package twitter_test

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"
	"time"

	twitter "github.com/cvcio/twitter"
)

var (
	consumerKey       = os.Getenv("TEST_TWITTER_CONSUMER_KEY")
	consumerSecret    = os.Getenv("TEST_TWITTER_CONSUMER_SECRET")
	accessToken       = os.Getenv("TEST_TWITTER_ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("TEST_TWITTER_ACCESS_TOKEN_SECRET")
)

// Test_API_NewAPI_Client Test New Twitter API Client
func Test_NewTwitter_Client(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Couldn't create Twitter API HTTPClient")
	}
	if api.GetClient() == nil {
		t.Fatalf("Twitter API HTTP Client returned nil")
	}
}

func Test_NewTwitterWithContext_Client(t *testing.T) {
	api, err := twitter.NewTwitterWithContext(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	if err != nil {
		t.Fatalf("Couldn't create Twitter API HTTPClient")
	}
	if api.GetClient() == nil {
		t.Fatalf("Twitter API HTTP Client returned nil")
	}
}

// Test_API_NewAPI_VerifyCredentials Test Twitter API VerifyCredentials
func Test_NewTwitter_VerifyCredentials(t *testing.T) {
	api, _ := twitter.NewTwitter(consumerKey, consumerSecret)
	ok, err := api.VerifyCredentials()
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err)
	}

	if !ok {
		t.Fatalf("Twitter API VerifyCredentials Error: %v", ok)
	}
}

func Test_NewTwitterWithContext_VerifyCredentials(t *testing.T) {
	api, err := twitter.NewTwitterWithContext(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	ok, err := api.VerifyCredentials()
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err)
	}

	if !ok {
		t.Fatalf("Twitter API VerifyCredentials Error: %v", ok)
	}
}

func Test_GetUserFollowers(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err)
	}
	var data []*twitter.User

	v := url.Values{}
	v.Add("max_results", "50")
	res, _ := api.GetUserFollowers("44142397", v, twitter.WithRate(15*time.Minute/15), twitter.WithAuto(false)) // @andefined
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshar Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 50 {
		t.Fatalf("Twitter API GetUserFollowers Error. Should have returned 50, got %d", len(data))
	}
}

func Test_GetUserFollowing(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err)
	}
	var data []*twitter.User

	v := url.Values{}
	v.Add("max_results", "50")
	res, _ := api.GetUserFollowing("44142397", v, twitter.WithRate(15*time.Minute/15), twitter.WithAuto(false)) // @andefined
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshar Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 50 {
		t.Fatalf("Twitter API GetUserFollowing Error. Should have returned 50, got %d", len(data))
	}
}

func Test_GetUserTweets(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err)
	}
	var data []*twitter.User

	v := url.Values{}
	v.Add("max_results", "50")
	res, _ := api.GetUserTweets("44142397", v, twitter.WithRate(15*time.Minute/1500), twitter.WithAuto(false)) // @andefined
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshar Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 50 {
		t.Fatalf("Twitter API GetUserTweets Error. Should have returned 50, got %d", len(data))
	}
}
