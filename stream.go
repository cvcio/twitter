package twitter

// GetFilterStream streams tweets in real-time based on a specific set of filter rules.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 50/15m (app)
func (api *Twitter) GetFilterStream() {}

// GetFilterStreamRules returns a list of rules currently active on the streaming endpoint, either as a list or individually.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream/rules
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream-rules
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app)
func (api *Twitter) GetFilterStreamRules() {}

// PostFilterStreamRules adds or deletes rules to your stream.
// Endpoint URL: https://api.twitter.com/2/tweets/search/stream/rules
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/post-tweets-search-stream-rules
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app)
func (api *Twitter) PostFilterStreamRules() {}

// GetSampleStream streams about 1% of all Tweets in real-time.
// Endpoint URL: https://api.twitter.com/2/tweets/sample/stream
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/sampled-stream/api-reference/get-tweets-sample-stream
// Authentication Methods: OAuth 2.0 Bearer Token
// Rate Limit: 50/15m (app)
func (api *Twitter) GetSampleStream() {}
