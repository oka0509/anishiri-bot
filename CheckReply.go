package main

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

func CheckReply(api *anaconda.TwitterApi, mention anaconda.Tweet) (bool) {
	sending2 :=url.Values{}
	//できればmaxとりたい(どこまでいけるかは不明)
	sending2.Add("count", "100")
	res, err := api.GetSearch("@" + mention.User.ScreenName, sending2)
	if err != nil {
		panic(err)
	}
	flag := false
	for _, tweet :=range res.Statuses {
		if tweet.User.ScreenName == "testbot14878693" && 
		tweet.InReplyToStatusID == mention.Id {
			flag = true
		}
	}
	return flag
}