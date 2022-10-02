package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
