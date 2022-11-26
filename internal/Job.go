package internal

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"github.com/ChimeraCoder/anaconda"
	_ "github.com/go-sql-driver/mysql"
)

type Word struct {
	ID   int
	Word string
}

func Job() {
	//環境変数を読み込み
	Loadenv()
	//dbに接続
	db := ConnectDb()
	defer db.Close()

	//認証・apiを作成
	api := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("access_token"),
		os.Getenv("access_token_secret"),
		os.Getenv("api_key"),
		os.Getenv("api_key_secret"),
	)

	//自分への一連のメンションを取得
	mentions, err := api.GetMentionsTimeline(url.Values{})
	if err != nil {
		log.Fatalf("Failed to get mentions: %s", err)
	}

	fmt.Println(len(mentions))
	//自分への各メンションについて返信する(またはしない)
	for _, mention := range mentions {
		//既に返信済みであればcontinue
		if CheckReply(api, mention) {
			continue
		}
		rs := []rune(mention.FullText)
		mentionExcludedText := ""
		for i, c := range mention.FullText {
			if string(c) == " " || string(c) == "　" {
				mentionExcludedText = string(rs[i+1:])
				break
			}
		}
		if mentionExcludedText == "" {
			continue
		}
		sending := url.Values{}
		//できればmaxとりたい(どこまでいけるかは不明)
		sending.Add("count", "100")
		row := Word{0, ""}
		rs2 := []rune(mentionExcludedText)
		db.Where("word LIKE ?", string(rs2[len(rs2)-1])+"%").First(&row)
		if row.Word == "" {
			continue
		}
		sending.Add("in_reply_to_status_id", mention.IdStr)
		text := "@" + mention.User.ScreenName + " " + row.Word
		_, err2 := api.PostTweet(text, sending)
		if err2 != nil {
			panic(err2)
		}
	}

	fmt.Println("Job")
}
