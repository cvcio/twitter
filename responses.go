package twitter

import "time"

// Response Struct
type Response struct {
	Results Data
	Error   error
}

// Meta Struct
type Meta struct {
	ResultCount   int    `json:"result_count"`
	NextToken     string `json:"next_token"`
	PreviousToken string `json:"previous_token"`
}

// Data Struct
type Data struct {
	Data     interface{} `json:"data"`
	Includes Includes    `json:"includes"`
	Meta     Meta        `json:"meta"`
}

// Twitter Specific Data

// Coordinates response object.
type Coordinates struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Geo response object.
type Geo struct {
	Coordinates Coordinates `json:"coordinates"`
	PlaceID     string      `json:"place_id"`
}

// ReferencedTweet response object.
type ReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Attachment response object.
type Attachment struct {
	MediaKeys []string `json:"media_keys"`
	PollIDs   []string `json:"poll_ids"`
}

// Annotation response object.
type Annotation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ContextAnnotation response object.
type ContextAnnotation struct {
	Domain Annotation `json:"domain"`
	Entity Annotation `json:"entity"`
}

// Entity response object.
type Entity struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// EntityAnnotation response object.
type EntityAnnotation struct {
	Entity
	Probability    float64 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

// EntityURL response object.
type EntityURL struct {
	Entity
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	UnwoundURL  string `json:"unwound_url"`
}

// EntityTag response object.
type EntityTag struct {
	Entity
	Tag string `json:"tag"`
}

// EntityMention response object.
type EntityMention struct {
	Entity
	UserName string `json:"username"`
}

// Entities response object.
type Entities struct {
	Annotations []EntityAnnotation `json:"annotations,omitempty"`
	URLs        []EntityURL        `json:"urls,omitempty"`
	HashTags    []EntityTag        `json:"hashtags,omitempty"`
	Mentions    []EntityMention    `json:"mentions,omitempty"`
	CashTags    []EntityTag        `json:"cashtags,omitempty"`
}

// Withheld response object.
type Withheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

// TweetMetrics response object.
type TweetMetrics struct {
	Retweets          int `json:"retweet_count,omitempty"`
	Replies           int `json:"reply_count,omitempty"`
	Likes             int `json:"like_count,omitempty"`
	Quotes            int `json:"quote_count,omitempty"`
	Impressions       int `json:"impression_count,omitempty"`
	URLLinkClicks     int `json:"url_link_clicks,omitempty"`
	UserProfileClicks int `json:"user_profile_clicks,omitempty"`
}

// UserMetrics response object.
type UserMetrics struct {
	Followers int `json:"followers_count"`
	Following int `json:"following_count"`
	Tweets    int `json:"tweet_count"`
	Listed    int `json:"listed_count"`
}

// Includes response object.
type Includes struct {
	Tweets []Tweet `json:"tweets,omitempty"`
	Users  []User  `json:"users,omitempty"`
}

// Error response object.
type Error struct{}

// Tweet response object as returned from /2/tweets endpoint. For detailed information
// refer to https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets.
type Tweet struct {
	ID                 string               `json:"id"`
	Text               string               `json:"text"`
	CreatedAt          string               `json:"created_at,omitempty"`
	AuthorID           string               `json:"author_id,omitempty"`
	ConversationID     string               `json:"converstation_id,omitempty"`
	InReplyToUserID    string               `json:"in_reply_to_user_id,omitempty"`
	ReferencedTweets   []*ReferencedTweet   `json:"referenced_tweets,omitempty"`
	Attachments        *Attachment          `json:"attachments,omitempty"`
	Geo                *Geo                 `json:"geo,omitempty"`
	ContextAnnotations []*ContextAnnotation `json:"context_annotations,omitempty"`
	Entities           *Entities            `json:"entities,omitempty"`
	Withheld           *TweetMetrics        `json:"withheld,omitempty"`
	PublicMetrics      *TweetMetrics        `json:"public_metrics,omitempty"`
	NonPublicMetrics   *TweetMetrics        `json:"non_public_metrics,omitempty"`
	OrganicMetrics     *TweetMetrics        `json:"organic_metrics,omitempty"`
	PromotedMetrics    *TweetMetrics        `json:"promoted_metrics,omitempty"`
	PossibySensitive   bool                 `json:"possiby_sensitive,omitempty"`
	Lang               string               `json:"lang,omitempty"`
	ReplySettings      string               `json:"reply_settings,omitempty"`
	Source             string               `json:"source,omitempty"`
	Includes           *Includes            `json:"includes,omitempty"`
	Errors             *Error               `json:"errors,omitempty"`
}

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (t Tweet) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.CreatedAt)
}

// User response object as returned from /2/users endpoint. For detailed information
// refer to https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users.
type User struct {
	ID              string       `json:"id"`
	Name            string       `json:"name"`
	UserName        string       `json:"username"`
	CreatedAt       string       `json:"created_at,omitempty"`
	Protected       bool         `json:"protected,omitempty"`
	Withheld        *Withheld    `json:"withheld,omitempty"`
	Location        string       `json:"location,omitempty"`
	URL             string       `json:"url,omitempty"`
	Description     string       `json:"description,omitempty"`
	Verified        bool         `json:"verified,omitempty"`
	Entities        *Entities    `json:"entities,omitempty"`
	ProfileImageURL string       `json:"profile_image_url,omitempty"`
	PublicMetrics   *UserMetrics `json:"public_metrics,omitempty"`
	PinnedTweetID   string       `json:"pinned_tweet_id,omitempty"`
	Includes        *Includes    `json:"includes,omitempty"`
	Errors          *Error       `json:"errors,omitempty"`
}

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (u User) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, u.CreatedAt)
}
