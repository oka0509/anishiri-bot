package main

import(
	"github.com/ChimeraCoder/anaconda"
    "net/url"
    "log"
    "fmt"
    "os"
    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := os.Getenv("user")
    PASS := os.Getenv("pass")
    PROTOCOL := os.Getenv("protocol")
    DBNAME := os.Getenv("dbname")
 
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}

type Words struct {
    ID int
    Word string
}

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
    
    db, err5 := sqlConnect()
    if err5 != nil {
        panic(err5.Error())
    } else {
        fmt.Println("DB接続成功")
    }
    defer db.Close()
    error := db.Create(&Words{
        ID:     3,
        Word:   "hunya-",
    }).Error
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println("データ追加成功")
    }
    api := anaconda.NewTwitterApiWithCredentials(
    os.Getenv("access-token"), 
    os.Getenv("access-token-secret"),
    os.Getenv("consumer-key"),
    os.Getenv("consumer-key-secret"),
    )

    params := url.Values{}
	mentions, err2 := api.GetMentionsTimeline(params)
	if err2 != nil {
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
        _, err3 := api.PostTweet(text, sending)
        if err3 != nil {
            panic(err2)
        }
    } 
}