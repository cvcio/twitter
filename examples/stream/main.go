package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/cvcio/twitter"
)

func main() {
	consumerKey := flag.String("consumer-key", "", "twitter API consumer key")
	consumerSecret := flag.String("consumer-secret", "", "twitter API consumer secret")

	flag.Parse()

	api, err := twitter.NewTwitter(*consumerKey, *consumerSecret)
	if err != nil {
		panic(err)
	}

	rulesToDelete, err := api.GetFilterStreamRules(nil)
	if err != nil {
		panic(err)
	}
	var ids []string
	for _, v := range rulesToDelete.Data {
		ids = append(ids, v.ID)
		fmt.Printf("Delete rule: %s\n", v.ID)
	}

	if len(ids) > 0 {
		rulesdel := new(twitter.Rules)
		rulesdel.Delete = &twitter.RulesDelete{
			Ids: ids,
		}

		deleted, err := api.PostFilterStreamRules(nil, rulesdel)
		if err != nil {
			panic(err)
		}

		if deleted == nil {
			panic("rules not deleted")
		}
	}

	rules := new(twitter.Rules)
	rules.Add = append(rules.Add, &twitter.RulesData{
		Value: "elon musk",
		Tag:   "test-client",
	})

	added, err := api.PostFilterStreamRules(nil, rules)
	if err != nil {
		panic(err)
	}

	if added == nil {
		panic("rules not added")
	}

	v := url.Values{}
	v.Add("expansions", "author_id,attachments.media_keys")

	v.Add("tweet.fields", "id,text,edit_history_tweet_ids,attachments,author_id,context_annotations,conversation_id,created_at,edit_controls,entities,in_reply_to_user_id,lang,non_public_metrics,organic_metrics,possibly_sensitive,promoted_metrics,public_metrics,referenced_tweets,reply_settings,source,withheld")
	v.Add("user.fields", "id,name,username,created_at,description,entities,location,pinned_tweet_id,profile_image_url,protected,public_metrics,url,verified,withheld")
	v.Add("media.fields", "media_key,type,url,duration_ms,height,width,non_public_metrics,organic_metrics,preview_image_url,promoted_metrics,public_metrics,alt_text,variants")

	s, err := api.GetFilterStream(v)
	if err != nil {
		panic(err)
	}
	for t := range s.C {
		f, _ := t.(twitter.StreamData)
		if f.Data.Includes != nil {
			for _, x := range f.Data.Includes.Tweets {
				fmt.Printf("%s\n", x.Text)
			}
		} else {
			fmt.Printf("%s\n", f.Data.Text)
		}
	}
	s.Stop()
}
