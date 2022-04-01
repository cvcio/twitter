package twitter_test

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/cvcio/twitter"
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
		t.Fatalf("Couldn't create Twitter API HTTP Client")
	}
	if api.GetClient() == nil {
		t.Fatalf("Twitter API HTTP Client returned nil")
	}
}

func Test_NewTwitterWithContext_Client(t *testing.T) {
	api, err := twitter.NewTwitterWithContext(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	if err != nil {
		t.Fatalf("Couldn't create Twitter API HTTP Client")
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
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}

	if !ok {
		t.Fatalf("Twitter API VerifyCredentials Error: %v", ok)
	}
}

func Test_NewTwitterWithContext_VerifyCredentials(t *testing.T) {
	api, err := twitter.NewTwitterWithContext(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	ok, err := api.VerifyCredentials()
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}

	if !ok {
		t.Fatalf("Twitter API VerifyCredentials Error: %v", ok)
	}
}

func Test_GetUserFollowers(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data []*twitter.User

	v := url.Values{}
	v.Add("max_results", "50")
	v.Add("user.fields", "created_at,description,id,location,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified")
	v.Add("expansions", "pinned_tweet_id")
	v.Add("tweet.fields", "created_at,id,lang,source,public_metrics")
	res, errs := api.GetUserFollowers("44142397", v, twitter.WithRate(15*time.Minute/15), twitter.WithAuto(false)) // @andefined

	for {
		r, rok := <-res
		if !rok {
			break
		}

		if r != nil {
			b, err := json.Marshal(r.Data)
			if err != nil {
				t.Errorf("json Marshal Error: %v", err)
			}

			json.Unmarshal(b, &data)
		}

		e, eok := <-errs
		if !eok {
			break
		}

		if e != nil {
			t.Errorf("Twitter API Error: %v", e)
			break
		}
	}

	if len(data) != 50 {
		t.Fatalf("Twitter API GetUserFollowers Error. Should have returned 50, got %d", len(data))
	}
}

func Test_GetUserFollowers_Error(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}

	v := url.Values{}
	v.Add("max_results", "5000")
	res, errs := api.GetUserFollowers("44142397", v, twitter.WithRate(15*time.Minute/15), twitter.WithAuto(true)) // @andefined

	for {
		_, rok := <-res
		if !rok {
			break
		}

		e, ok := <-errs
		if !ok {
			break
		}

		if e == nil || e.Code != 400 {
			t.Fatalf("Should have returned 400: %v", e)
			break
		}
	}
}

func Test_Channels(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}

	size := 0
	v := url.Values{}
	v.Add("max_results", "1000")
	res, errs := api.GetUserFollowing("44142397", v, twitter.WithRate(time.Minute/6), twitter.WithAuto(true)) // @andefined

	for {
		select {
		case r, ok := <-res:
			if !ok {
				res = nil
				break
			}

			var d []*twitter.User
			b, err := json.Marshal(r.Data)
			if err != nil {
				t.Fatalf("json Marshal Error: %v", err)
			}

			json.Unmarshal(b, &d)
			size += len(d)

		case e, ok := <-errs:
			if !ok {
				errs = nil
				break
			}
			t.Errorf("Twitter API Error: %v", e)
		}

		if res == nil && errs == nil {
			break
		}

	}

	if size < 1000 {
		t.Fatalf("Twitter API GetUserFollowing Error. Should have returned more than 1000, got %d", size)
	}
}

func Test_GetUserFollowing(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
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
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 50 {
		t.Fatalf("Twitter API GetUserFollowing Error. Should have returned 50, got %d", len(data))
	}
}

func Test_GetUsers(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data []*twitter.User

	v := url.Values{}
	v.Add("ids", "44142397,334602996")
	res, _ := api.GetUsers(v) // @andefined, @atsipras
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 2 {
		t.Fatalf("Twitter API GetUsers Error. Should have returned 2, got %d", len(data))
	}
}

func Test_GetUsersBy(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data []*twitter.User

	v := url.Values{}
	v.Add("usernames", "andefined,atsipras")
	res, _ := api.GetUsersBy(v) // @andefined, @atsipras
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 2 {
		t.Fatalf("Twitter API GetUsersBy Error. Should have returned 2, got %d", len(data))
	}
}

func Test_GetUserByID(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data *twitter.User

	v := url.Values{}
	res, _ := api.GetUserByID("44142397", v) // @andefined
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if data.UserName != "andefined" {
		t.Fatalf("Twitter API GetUserByID Error. Should have returned andefined, got %s", data.UserName)
	}
}

func Test_GetUsersByUserName(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data *twitter.User

	v := url.Values{}
	res, _ := api.GetUsersByUserName("andefined", v) // @andefined
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if data.UserName != "andefined" {
		t.Fatalf("Twitter API GetUsersByUserName Error. Should have returned andefined, got %s", data.UserName)
	}
}

func Test_GetUserTweets(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data []*twitter.Tweet

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
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) < 1 {
		t.Fatalf("Twitter API GetUserTweets Error. Should have returned >= 1, got %d", len(data))
	}
}

func Test_GetUsesMentions(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data []*twitter.Tweet

	v := url.Values{}
	v.Add("max_results", "10")
	res, _ := api.GetUserMentions("44142397", v, twitter.WithRate(15*time.Minute/450), twitter.WithAuto(false)) // @andefined
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 10 {
		t.Fatalf("Twitter API GetUserMentions Error. Should have returned 10, got %d", len(data))
	}
}

func Test_GetTweets(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data []*twitter.Tweet

	v := url.Values{}
	v.Add("ids", "1370136892432322569,1370704815983038469")
	res, _ := api.GetTweets(v, twitter.WithRate(15*time.Minute/300), twitter.WithAuto(false)) // https://twitter.com/andefined/status/1370136892432322569, https://twitter.com/andefined/status/1370704815983038469
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 2 {
		t.Fatalf("Twitter API GetTweets Error. Should have returned 2, got %d", len(data))
	}

	if data[0].Text != "interesting comparison @kmitsotakis vs @atsipras https://t.co/WWcQjOgAtz" {
		t.Fatalf("Twitter API GetTweets Error. Should have returned `interesting comparison @kmitsotakis vs @atsipras https://t.co/WWcQjOgAtz`, got %s", data[0].Text)
	}
}

func Test_GetTweetByID(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data *twitter.Tweet

	v := url.Values{}
	res, _ := api.GetTweetByID("1370136892432322569", v) // https://twitter.com/andefined/status/1370136892432322569
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if data.Text != "interesting comparison @kmitsotakis vs @atsipras https://t.co/WWcQjOgAtz" {
		t.Fatalf("Twitter API GetTweets Error. Should have returned `interesting comparison @kmitsotakis vs @atsipras https://t.co/WWcQjOgAtz`, got %s", data.Text)
	}
}

func Test_GetTweetsSearchRecent(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}
	var data []*twitter.Tweet

	v := url.Values{}
	v.Add("query", "covid")
	v.Add("max_results", "100")
	res, _ := api.GetTweetsSearchRecent(v, twitter.WithAuto(false))
	for {
		r, ok := <-res

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			t.Fatalf("json Marshal Error: %v", err)
		}

		json.Unmarshal(b, &data)
	}

	if len(data) != 100 {
		t.Fatalf("Twitter API GetTweetsSearchRecent Error. Should have returned 100, got %d", len(data))
	}
}

func Test_PostFilterStreamRules(t *testing.T) {
	api, _ := twitter.NewTwitter(consumerKey, consumerSecret)

	rules := new(twitter.Rules)
	rules.Add = append(rules.Add, &twitter.RulesData{
		Value: "greece",
		Tag:   "test-client",
	})
	v := url.Values{}
	res, err := api.PostFilterStreamRules(v, rules)
	if err != nil {
		t.Fatalf("Twitter API PostFilterStreamRules Error: %s", err.Message)
	}

	if len(res.Data) != 1 {
		t.Fatalf("Twitter API PostFilterStreamRules Error. Should have returned 1, got %d", len(res.Data))
	}

	if len(res.Errors) != 0 {
		t.Fatalf("Twitter API PostFilterStreamRules Error. Should have returned 0 error message, got %d", len(res.Errors))
	}

}

func Test_GetFilterStreamRules(t *testing.T) {
	api, _ := twitter.NewTwitter(consumerKey, consumerSecret)
	_, err := api.GetFilterStreamRules(nil)
	if err != nil {
		t.Fatalf("Twitter API GetFilterStreamRules Error: %s", err.Message)
	}
}

func Test_GetFilterStream(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}

	v := url.Values{}
	v.Add("user.fields", "created_at,description,id,location,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified")
	v.Add("expansions", "author_id,in_reply_to_user_id")
	v.Add("tweet.fields", "created_at,id,lang,source,public_metrics")

	var f twitter.StreamData
	s, serr := api.GetFilterStream(v)
	if serr != nil {
		t.Fatalf("Twitter API GetSampleStream Error: %s", serr.Message)
	}
	for t := range s.C {
		f, _ = t.(twitter.StreamData)
		break
	}
	s.Stop()

	if f.Tweet.ID == "" {
		t.Fatalf("Twitter API GetSampleStream Error. Should have return a valid tweet struct, got %v", f.Tweet)
	}
}

func Test_GetSampleStream(t *testing.T) {
	api, err := twitter.NewTwitter(consumerKey, consumerSecret)
	if err != nil {
		t.Fatalf("Twitter API VerifyCredentials Error: %s", err.Message)
	}

	v := url.Values{}
	v.Add("user.fields", "created_at,description,id,location,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified")
	v.Add("expansions", "author_id,in_reply_to_user_id")
	v.Add("tweet.fields", "created_at,id,lang,source,public_metrics")

	var f twitter.StreamData
	s, serr := api.GetSampleStream(v)
	if serr != nil {
		t.Fatalf("Twitter API GetSampleStream Error: %s", serr.Message)
	}
	for t := range s.C {
		f, _ = t.(twitter.StreamData)
		break
	}
	s.Stop()

	if f.Tweet.ID == "" {
		t.Fatalf("Twitter API GetSampleStream Error. Should have return a valid tweet struct, got %v", f.Tweet)
	}
}

func Test_DeleteFilterStreamRules(t *testing.T) {
	api, _ := twitter.NewTwitter(consumerKey, consumerSecret)

	rulesToDelete, err := api.GetFilterStreamRules(nil)
	if err != nil {
		t.Fatalf("Twitter API GetFilterStreamRules Error: %s", err.Message)
	}
	var ids []string
	for _, v := range rulesToDelete.Data {
		ids = append(ids, v.ID)
	}
	rules := new(twitter.Rules)
	rules.Delete = &twitter.RulesDelete{
		Ids: ids,
	}

	res, err := api.PostFilterStreamRules(nil, rules)

	if err != nil {
		t.Fatalf("Twitter API PostFilterStreamRules Error: %s", err.Message)
	}

	if res.Meta.Summary.Deleted != 1 {
		t.Fatalf("Twitter API PostFilterStreamRules Error. Should have returned 1, got %d", res.Meta.Summary.Deleted)
	}

	if len(res.Errors) > 0 {
		t.Fatalf("Twitter API PostFilterStreamRules Error. Should have returned 0 error message, got %d", len(res.Errors))
	}
}
