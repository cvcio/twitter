package twitter

import (
	"time"
)

// Response Struct
type Response struct {
	Results Data
	Error   *APIError
}

// Meta Struct
type Meta struct {
	ResultCount   int    `json:"result_count,omitempty"`
	NextToken     string `json:"next_token,omitempty"`
	PreviousToken string `json:"previous_token,omitempty"`
}

// Data Struct
type Data struct {
	Data     interface{} `json:"data,omitempty"`
	Includes Includes    `json:"includes,omitempty"`
	Meta     Meta        `json:"meta,omitempty"`
}

// Twitter Specific Data

// Coordinates response object.
type Coordinates struct {
	Type        string    `json:"type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"`
}

// Geo response object.
type Geo struct {
	Coordinates Coordinates `json:"coordinates,omitempty"`
	PlaceID     string      `json:"place_id,omitempty"`
}

// ReferencedTweet response object.
type ReferencedTweet struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Attachment response object.
type Attachment struct {
	MediaKeys []string `json:"media_keys,omitempty"`
	PollIDs   []string `json:"poll_ids,omitempty"`
}

// Annotation response object.
type Annotation struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// ContextAnnotation response object.
type ContextAnnotation struct {
	Domain Annotation `json:"domain,omitempty"`
	Entity Annotation `json:"entity,omitempty"`
}

// Entity response object.
type Entity struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// EntityAnnotation response object.
type EntityAnnotation struct {
	Entity
	Probability    float64 `json:"probability,omitempty"`
	Type           string  `json:"type,omitempty"`
	NormalizedText string  `json:"normalized_text,omitempty"`
}

// EntityURL response object.
type EntityURL struct {
	Entity
	URL         string `json:"url,omitempty"`
	ExpandedURL string `json:"expanded_url,omitempty"`
	DisplayURL  string `json:"display_url,omitempty"`
	UnwoundURL  string `json:"unwound_url,omitempty"`
}

// EntityTag response object.
type EntityTag struct {
	Entity
	Tag string `json:"tag,omitempty"`
}

// EntityMention response object.
type EntityMention struct {
	Entity
	UserName string `json:"username,omitempty"`
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
	Copyright    bool     `json:"copyright,omitempty"`
	CountryCodes []string `json:"country_codes,omitempty"`
	Scope        string   `json:"scope,omitempty"`
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
	Followers int `json:"followers_count,omitempty"`
	Following int `json:"following_count,omitempty"`
	Tweets    int `json:"tweet_count,omitempty"`
	Listed    int `json:"listed_count,omitempty"`
}

// Includes response object.
type Includes struct {
	Tweets []Tweet `json:"tweets,omitempty"`
	Users  []User  `json:"users,omitempty"`
}

// Error response object.
type Error struct{}
type StreamData struct {
	Tweet Tweet `json:"data"`
}

// Tweet response object as returned from /2/tweets endpoint. For detailed information
// refer to https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets.
type Tweet struct {
	ID                 string               `json:"id"`
	Text               string               `json:"text,omitempty"`
	CreatedAt          string               `json:"created_at,omitempty"`
	AuthorID           string               `json:"author_id,omitempty"`
	ConversationID     string               `json:"converstation_id,omitempty"`
	InReplyToUserID    string               `json:"in_reply_to_user_id,omitempty"`
	ReferencedTweets   []*ReferencedTweet   `json:"referenced_tweets,omitempty"`
	Attachments        *Attachment          `json:"attachments,omitempty"`
	Geo                *Geo                 `json:"geo,omitempty"`
	ContextAnnotations []*ContextAnnotation `json:"context_annotations,omitempty"`
	Entities           *Entities            `json:"entities,omitempty"`
	Withheld           *Withheld            `json:"withheld,omitempty"`
	PublicMetrics      *TweetMetrics        `json:"public_metrics,omitempty"`
	NonPublicMetrics   *TweetMetrics        `json:"non_public_metrics,omitempty"`
	OrganicMetrics     *TweetMetrics        `json:"organic_metrics,omitempty"`
	PromotedMetrics    *TweetMetrics        `json:"promoted_metrics,omitempty"`
	PossibySensitive   bool                 `json:"possibly_sensitive,omitempty"`
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
	Name            string       `json:"name,omitempty"`
	UserName        string       `json:"username,omitempty"`
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
	Sent        time.Time     `json:"sent,omitempty"`
	ResultCount int           `json:"result_count,omitempty"`
	Summary     *RulesSummary `json:"summary,omitempty"`
}
type RulesError struct {
	Value string `json:"value,omitempty"`
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
}
type RulesDelete struct {
	Ids []string `json:"ids,omitempty"`
}

type Rules struct {
	Add    []*RulesData             `json:"add"`
	Delete *RulesDelete             `json:"delete"`
	Data   []*RulesData             `json:"data"`
	Meta   *RulesMeta               `json:"meta"`
	Errors []map[string]interface{} `json:"errors"`
}

/*
func (rules *Rules) Make(m map[string]interface{}) error {
	for k, v := range m {
		err := Field(rules, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func Field(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
*/
