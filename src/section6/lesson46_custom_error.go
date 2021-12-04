package main

import "fmt"

type UserNotFound struct {
	Username string
}

// カスタムエラー
func (e *UserNotFound) Error() string {
	// 単に値返すだけでも良いし、コメント書いても良い。
	// return e.Username
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	// Something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
