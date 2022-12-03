package internal

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("環境変数の読み込みに失敗しました: %v", err)
	}
}
