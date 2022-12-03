package internal

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

//DB接続
func ConnectDb() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "root"
    PASS := "mypassword"
    PROTOCOL := "tcp(db)"
    DBNAME := "testdb"
 
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}
