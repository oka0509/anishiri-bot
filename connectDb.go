package main

import (
	"fmt"
	"os"
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

func connectDb() (database *gorm.DB) {
    db, err5 := sqlConnect()
    if err5 != nil {
        panic(err5.Error())
    } else {
        fmt.Println("DB接続成功")
    }
	return db
}
