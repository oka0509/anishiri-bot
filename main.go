package main

import (
	"time"
	"log"
	"twitter-bot/internal"
)

func main() {
	//dbに接続
	db, err:= internal.ConnectDb(); 
	for err != nil {
		db, err = internal.ConnectDb()
	}
	log.Printf("dbに接続しました")
	for {
		internal.Job(db)
		time.Sleep(time.Minute)
	}
}