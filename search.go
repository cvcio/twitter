package twitter

// GetTweetsSearchRecent returns Tweets from the last 7 days that match a search query.
// Endpoint URL: https://api.twitter.com/2/tweets/search/recent
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app), 180/15m (user)
func (api *Twitter) GetTweetsSearchRecent() {}

// GetTweetsSearchAll returns the complete history of public Tweets matching a search query;
// since the first Tweet was created March 26, 2006. This endpoint is part of a
// **private beta** for academic researchers only.
// *Please do not share this documentation.*
// Endpoint URL: https://api.twitter.com/2/tweets/search/all
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/full-archive-search/api-reference/get-tweets-search-all
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 1/1s (user)
func (api *Twitter) GetTweetsSearchAll() {}
