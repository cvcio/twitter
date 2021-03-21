# Twitter API v2 Client for [Go](https://golang.org/)

**twitter** is a Go package for the [Twitter v2 API](https://developer.twitter.com/en/docs/twitter-api/early-access), inspired by [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda). This library uses channels for both, to retrieve data from Twitter API and return the results, with a built-in throttle to avoid rate-limit errors return from Twitter. You can bypass the throttling by usgin `twitter.WithRate(time.Duration)` option. The library will auto paginate results unless you use the `twitter.WithAuto(false)` option.

### Installation

```bash
go get github.com/cvcio/twitter
```

### Supported Endpoints

| Method | Implemented | OAuth | Rate Limit | Official Documentation |
|--------|----------|------|------------|------------------------|
| `VerifyCredentials` | Yes | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | - | - |
| `GetUserFollowers` | Yes | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 15/15m (app), 15/15m (user) | [Get User Followers](https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers) |
| `GetUserFollowing` | Yes | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 15/15m (app), 15/15m (user) | [Get User Following](https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following) |
| `GetUsers` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 300/15m (app), 900/15m (user) | [Get Users](https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following) |
| `GetUsersBy` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 300/15m (app), 900/15m (user) | [Get Users By](https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following) |
| `GetUserByID` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 300/15m (app), 900/15m (user) | [Get User By Id](https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following) |
| `GetUsersByUserName` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 300/15m (app), 900/15m (user) | [Get Users By Screen Name](https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following) |
| `GetUserTweets` | Yes | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 1500/15m (app), 900/15m (user) | [Get User Tweets](https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets) |
| `GetUserMentions` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 450/15m (app), 180/15m (user) | [Get User Mentions](https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-mentions) |
| `GetTweets` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 300/15m (app), 900/15m (user) | [Get Tweets](https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets) |
| `GetTweetByID` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 300/15m (app), 900/15m (user) | [Get Tweets By Id](https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id) |
| `GetFilterStream` | Not Yet | OAuth 2.0 Bearer Token | 50/15m (app) | [Filter Stream](https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream) |
| `GetFilterStreamRules` | Not Yet | OAuth 2.0 Bearer Token | 450/15m (app) | [Get Filter Stream Rules](https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream-rules)
| `PostFilterStreamRules` | Not Yet | OAuth 2.0 Bearer Token | 450/15m (app) | [Post Filter Stream Rules](https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/post-tweets-search-stream-rules)
| `GetSampleStream` | Not Yet | OAuth 2.0 Bearer Token | 50/15m (app) | [Sample Stream](https://developer.twitter.com/en/docs/twitter-api/tweets/sampled-stream/api-reference/get-tweets-sample-stream)
| `GetTweetsSearchRecent` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 450/15m (app), 180/15m (user) | [Sample Stream](https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent)
| `GetTweetsSearchAll` | Not Yet | OAuth 1.0a User Context, OAuth 2.0 Bearer Token | 300/15m (app), 1/1s (user) | [Sample Stream](https://developer.twitter.com/en/docs/twitter-api/tweets/full-archive-search/api-reference/get-tweets-search-all)

### Usage

#### Authentication

```go
api, err := twitter.NewTwitterWithContext(*consumerKey, *consumerSecret, *accessToken, *accessTokenSecret)
if err != nil {
	panic(err)
}
```

#### Methods
Each method returns 2 channels, one for results and one for errors (`twitter.APIError`).
```go
v := url.Values{}
v.Add("max_results", "1000")
resultsChannel, _ := api.GetUserFollowers(*id, v)
for {
	// resultsChannel
	r, ok := <-resultsChannel
	if !ok {
		break
	}
	...
}
```
#### Options

- `twitter.WithDealy(time.Duration)` 
- `twitter.WithRate(time.Duration)`
- `twitter.WithAuth(bool)`

### Examples

```go
// get all followers for a specific user id 
// and print the results
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"time"

	"github.com/cvcio/twitter"
)

func main() {
	consumerKey := flag.String("consumer-key", "", "twitter API consumer key")
	consumerSecret := flag.String("consumer-secret", "", "twitter API consumer secret")
	accessToken := flag.String("access-token", "", "twitter API access token")
	accessTokenSecret := flag.String("access-token-secret", "", "twitter API access token secret")

	id := flag.String("id", "", "user id")

	flag.Parse()

	start := time.Now()

	api, err := twitter.NewTwitterWithContext(*consumerKey, *consumerSecret, *accessToken, *accessTokenSecret)
	if err != nil {
		panic(err)
	}

	v := url.Values{}
	v.Add("max_results", "1000")
	followers, _ := api.GetUserFollowers(*id, v)

	for {
		r, ok := <-followers

		if !ok {
			break
		}

		b, err := json.Marshal(r.Data)
		if err != nil {
			panic(err)
		}

		var data []*twitter.User
		json.Unmarshal(b, &data)
		for _, v := range data {
			fmt.Printf("%s,%s,%s\n", v.ID, v.UserName, v.Name)
		}

		fmt.Println()
		fmt.Printf("Result Count: %d Next Token: %s\n", r.Meta.ResultCount, r.Meta.NextToken)
	}

	end := time.Now()

	fmt.Printf("Done in %s", end.Sub(start))
}
```

## Contribution

If you're new to contributing to Open Source on Github, [this guide](https://opensource.guide/how-to-contribute/) can help you get started. Please check out the contribution guide for more details on how issues and pull requests work. Before contributing be sure to review the [code of conduct](https://github.com/cvcio/twitter/blob/main/CODE_OF_CONDUCT.md).

## Contributors

- Dimitris Papaevagelou ([@andefined](https://github.com/andefined))
- Ilias Dimos ([@dosko64](https://github.com/dosko64))

### License

This library is distributed under the MIT license found in the [LICENSE](https://github.com/cvcio/twitter/blob/main/LICENSE) file.