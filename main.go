package main

import(
	"github.com/ChimeraCoder/anaconda"
    "net/url"
    "log"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    "strings"
)

type Word struct {
    ID int
    Word string
}

func main(){
    //環境変数を読み込み
    Loadenv()

    //dbに接続
    db := connectDb()
    defer db.Close()

    
    api := anaconda.NewTwitterApiWithCredentials(
    os.Getenv("access-token"), 
    os.Getenv("access-token-secret"),
    os.Getenv("consumer-key"),
    os.Getenv("consumer-key-secret"),
    )

    params := url.Values{}
	mentions, err2 := api.GetMentionsTimeline(params)
	if err2 != nil {
		log.Fatalf("Failed to get mentions: %s", err2)
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
        arr1 := strings.Split(mention.Text, " ")
        text2 := ""
        if len(arr1)>=2 {
            text2 = arr1[1]
        }
        fmt.Println(text2)
        var row Word
        db.Where("word LIKE ?", "h"+"%").First(&row)
        fmt.Println(row.Word)
        if flag {
            continue
        }
        sending.Add("in_reply_to_status_id", mention.IdStr)
        text := "@" + mention.User.ScreenName +" raaaaas"
        _, err3 := api.PostTweet(text, sending)
        if err3 != nil {
            panic(err2)
        }
    } 
}