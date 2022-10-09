package internal

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "root"
    PASS := "mypassword"
    PROTOCOL := "tcp(db)"
    DBNAME := "testdb"
 
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}

func ConnectDb() (database *gorm.DB) {
    db, err5 := sqlConnect()
    if err5 != nil {
        panic(err5.Error())
    } else {
        fmt.Println("DB接続成功")
    }
	return db
}
