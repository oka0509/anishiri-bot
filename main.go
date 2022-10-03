package main

import(
	"github.com/ChimeraCoder/anaconda"
    "net/url"
    "log"
    _"fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

type Word struct {
    ID int
    Word string
}

func main(){
    //環境変数を読み込み
    Loadenv()

    //dbに接続
    db := ConnectDb()
    defer db.Close()

    //認証・apiを作成
    api := anaconda.NewTwitterApiWithCredentials(
    os.Getenv("access-token"), 
    os.Getenv("access-token-secret"),
    os.Getenv("consumer-key"),
    os.Getenv("consumer-key-secret"),
    )

    //自分への一連のメンションを取得
	mentions, err2 := api.GetMentionsTimeline(url.Values{})
	if err2 != nil {
		log.Fatalf("Failed to get mentions: %s", err2)
	}

    //自分への各メンションについて返信する(またはしない)
    for _, mention := range mentions {
        rs := []rune(mention.FullText)
        mentionExcludedText := ""
        if len(rs) < 16 || string(rs[:16]) != "@testbot14878693" {
            mentionExcludedText = mention.Text
        } else {
            for i, c :=range mention.Text {
                if string(c) == " " ||string(c) == "　" {
                    mentionExcludedText = string(rs[i+1:])
                    break
                } 
            }
        }
        if mentionExcludedText == "" {
            continue
        }
        //既に返信済みであればcontinue
        if CheckReply(api, mention) {
            continue
        }
        sending :=url.Values{}
        //できればmaxとりたい(どこまでいけるかは不明)
        sending.Add("count", "100")
        var row Word
        rs2 := []rune(mentionExcludedText)
        db.Where("word LIKE ?", string(rs2[len(rs2)-1])+"%").First(&row)
        if row.Word == "" {
            continue
        }
        sending.Add("in_reply_to_status_id", mention.IdStr)
        text := "@" + mention.User.ScreenName + " " + row.Word
        _, err3 := api.PostTweet(text, sending)
        if err3 != nil {
            panic(err2)
        }
    } 
}
