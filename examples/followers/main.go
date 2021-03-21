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
