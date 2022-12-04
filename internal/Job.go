package internal

import (
	"log"
	"net/url"
	"os"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	"github.com/ChimeraCoder/anaconda"
	_ "github.com/go-sql-driver/mysql"
)

type Word struct {
	ID   int
	Word string
}

func Job(db *gorm.DB) {
	//環境変数を読み込み
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("環境変数の読み込みに失敗しました: %v", err)
	}

	//認証・インスタンスを作成
	api := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("access_token"),
		os.Getenv("access_token_secret"),
		os.Getenv("api_key"),
		os.Getenv("api_key_secret"),
	)

	//自分への一連のメンションを取得
	mentions, err2 := api.GetMentionsTimeline(url.Values{})
	if err2 != nil {
		log.Printf("メンションの取得に失敗しました: %s", err2)
		return
	}

	//自分への各メンションについて返信する(またはしない)
	for _, mention := range mentions {
		//既に返信済みであればcontinue
		flag, err3:= CheckReply(api, mention)
		if err3 != nil {
			return
		}
		if(flag) {
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
			sending.Add("in_reply_to_status_id", mention.IdStr)
			text := "@" + mention.User.ScreenName + " " + "対応する単語がみつかりませんでした(末尾がひらがなの文字列のみに対応しています＞＜)"
			_, err4 := api.PostTweet(text, sending)
			if err4 != nil {
				log.Printf("tweetの送信に失敗しました: %s", err4)
				return
			}
		} else {
			sending.Add("in_reply_to_status_id", mention.IdStr)
			text := "@" + mention.User.ScreenName + " " + row.Word
			_, err4 := api.PostTweet(text, sending)
			if err4 != nil {
				log.Printf("tweetの送信に失敗しました: %s", err4)
				return
			}		
		}
	}

	log.Printf("Job終了")
}
