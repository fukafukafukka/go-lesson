package main

import "fmt"

// 最初の文字は大文字にする
const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)

	// 書き換えようとするとエラー
	// Pi = 3
}
