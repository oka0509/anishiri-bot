package main

import (
	"github.com/robfig/cron/v3"
	"time"
	"log"
	"twitter-bot/internal"
)

func main() {
	time.Sleep(time.Second * 10)
	//dbに接続
	db, err:= internal.ConnectDb(); 
	for err != nil {
		log.Printf("dbの接続に失敗しました: %s", err)
		db, err = internal.ConnectDb()
	}
	c := cron.New()
    c.AddFunc("@every 1m", func(){internal.Job(db)})
	c.Start()
	for {
		time.Sleep(10000000000000)
	}
}