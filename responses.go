package twitter

import "time"

// Response Struct
type Response struct {
	Results Data
	Error   *APIError
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
	Annotations []EntityAnnotation `json:"annotations"`
	URLs        []EntityURL        `json:"urls"`
	HashTags    []EntityTag        `json:"hashtags"`
	Mentions    []EntityMention    `json:"mentions"`
	CashTags    []EntityTag        `json:"cashtags"`
}

// Withheld response object.
type Withheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

// TweetMetrics response object.
type TweetMetrics struct {
	Retweets          int `json:"retweet_count"`
	Replies           int `json:"reply_count"`
	Likes             int `json:"like_count"`
	Quotes            int `json:"quote_count"`
	Impressions       int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
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
	Tweets []Tweet `json:"tweets"`
	Users  []User  `json:"users"`
}

// Error response object.
type Error struct{}
type StreamData struct {
	Data Tweet `json:"data"`
}

// Tweet response object as returned from /2/tweets endpoint. For detailed information
// refer to https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets.
type Tweet struct {
	ID                 string              `json:"id"`
	Text               string              `json:"text"`
	CreatedAt          string              `json:"created_at"`
	AuthorID           string              `json:"author_id"`
	ConversationID     string              `json:"converstation_id"`
	InReplyToUserID    string              `json:"in_reply_to_user_id"`
	ReferencedTweets   []ReferencedTweet   `json:"referenced_tweets"`
	Attachments        Attachment          `json:"attachments"`
	Geo                Geo                 `json:"geo"`
	ContextAnnotations []ContextAnnotation `json:"context_annotations"`
	Entities           Entities            `json:"entities"`
	Withheld           TweetMetrics        `json:"withheld"`
	PublicMetrics      TweetMetrics        `json:"public_metrics"`
	NonPublicMetrics   TweetMetrics        `json:"non_public_metrics"`
	OrganicMetrics     TweetMetrics        `json:"organic_metrics"`
	PromotedMetrics    TweetMetrics        `json:"promoted_metrics"`
	PossibySensitive   bool                `json:"possiby_sensitive"`
	Lang               string              `json:"lang"`
	ReplySettings      string              `json:"reply_settings"`
	Source             string              `json:"source"`
	Includes           Includes            `json:"includes"`
	Errors             Error               `json:"errors"`
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
	CreatedAt       string       `json:"created_at"`
	Protected       bool         `json:"protected"`
	Withheld        *Withheld    `json:"withheld"`
	Location        string       `json:"location"`
	URL             string       `json:"url"`
	Description     string       `json:"description"`
	Verified        bool         `json:"verified"`
	Entities        *Entities    `json:"entities"`
	ProfileImageURL string       `json:"profile_image_url"`
	PublicMetrics   *UserMetrics `json:"public_metrics"`
	PinnedTweetID   string       `json:"pinned_tweet_id"`
	Includes        *Includes    `json:"includes"`
	Errors          *Error       `json:"errors"`
}

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (u User) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, u.CreatedAt)
}

type RulesData struct {
	Value string `json:"value,omitempty"`
	Tag   string `json:"tag,omitempty"`
	ID    string `json:"id,omitempty"`
}
type RulesSummary struct {
	Created    int `json:"created,omitempty"`
	NotCreated int `json:"not_created,omitempty"`
	Deleted    int `json:"deleted,omitempty"`
	NotDeleted int `json:"not_deleted,omitempty"`
}
type RulesMeta struct {
	Sent    time.Time     `json:"sent,omitempty"`
	Summary *RulesSummary `json:"summary,omitempty"`
}

type Rules struct {
	Data []RulesData `json:"data,omitempty"`
	Meta *RulesMeta  `json:"meta,omitempty"`
}
