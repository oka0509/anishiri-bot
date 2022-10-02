package main

import(
	"github.com/ChimeraCoder/anaconda"
    "net/url"
    "log"
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
    api := anaconda.NewTwitterApiWithCredentials(
    os.Getenv("access-token"), 
    os.Getenv("access-token-secret"),
    os.Getenv("consumer-key"),
    os.Getenv("consumer-key-secret"),
    )

    params := url.Values{}
	mentions, err := api.GetMentionsTimeline(params)
	if err != nil {
		log.Fatalf("Failed to get mentions: %s", err)
	}


    for _, mention := range mentions {
        sending :=url.Values{}
        //できればmaxとりたい(どこまでいけるかは不明)
        sending.Add("count", "100")
        res, err := api.GetSearch("@" + mention.User.ScreenName, sending)
        if err != nil {
            panic(err)
        }
        //ここでtweetの返信先がmentionのidであるものがあれば外側のforをcontinue
        flag := false
        for _, tweet :=range res.Statuses {
            if tweet.User.ScreenName == "testbot14878693" && 
            tweet.InReplyToStatusID == mention.Id {
                flag = true
            }
        }
        if flag {
            continue
        }
        sending.Add("in_reply_to_status_id", mention.IdStr)
        text := "@" + mention.User.ScreenName +" raaaaas"
        _, err2 := api.PostTweet(text, sending)
        if err2 != nil {
            panic(err2)
        }
    } 
}