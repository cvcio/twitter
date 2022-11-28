package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/cvcio/twitter"
)

func main() {
	consumerKey := flag.String("consumer-key", "", "twitter API consumer key")
	consumerSecret := flag.String("consumer-secret", "", "twitter API consumer secret")

	id := flag.String("id", "", "user id")

	flag.Parse()

	api, err := twitter.NewTwitter(*consumerKey, *consumerSecret)
	if err != nil {
		panic(err)
	}

	fmt.Print("ID,CreatedAt,Name,UserName,Protected,Verified,Followers,Following,Tweets,Listed,ProfileImageURL\n")

	v := url.Values{}
	// set size of response ids to 1000
	v.Add("max_results", "1000")
	// set user fields to return
	v.Add("user.fields", "created_at,description,id,location,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified")
	// set tweet fields to return
	v.Add("tweet.fields", "created_at,id,lang,source,public_metrics")

	followers, _ := api.GetUserFollowers(*id, v, twitter.WithRate(15*time.Minute/15), twitter.WithAuto(true))

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
			fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
				v.ID,
				v.CreatedAt,
				v.Name,
				v.UserName,
				strconv.FormatBool(v.Protected),
				strconv.FormatBool(v.Verified),
				strconv.Itoa(v.PublicMetrics.Followers),
				strconv.Itoa(v.PublicMetrics.Following),
				strconv.Itoa(v.PublicMetrics.Tweets),
				strconv.Itoa(v.PublicMetrics.Listed),
				v.ProfileImageURL,
			)
		}

		log.Printf("Result Count: %d Next Token: %s\n", r.Meta.ResultCount, r.Meta.NextToken)
	}
}
