package internal

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
)

func CheckReply(api *anaconda.TwitterApi, mention anaconda.Tweet) (bool, error) {
	sending :=url.Values{}
	//できればmaxとりたい(どこまでいけるかは不明)
	sending.Add("count", "100")
	res, err := api.GetSearch("@" + mention.User.ScreenName, sending)
	if err != nil {
		return false, err
	}
	flag := false
	for _, tweet :=range res.Statuses {
		if tweet.User.ScreenName == os.Getenv("twitter_id") && 
		tweet.InReplyToStatusID == mention.Id {
			flag = true
		}
	}
	return flag, nil
}