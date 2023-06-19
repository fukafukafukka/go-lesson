package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "unix"
}

func main() {
	// os := getOsName()
	// switch os {
	// case "mac":
	// fmt.Println("Mac!")
	// case "windows":
	// fmt.Println("Windows!")
	// default:
	// fmt.Println("Default!")
	// }

	// 上記を以下のように短く書ける
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!")
	case "Windows!":
		fmt.Println("Windows!")
	default:
		fmt.Println("default", os)
	}

	// fmt.Println(os) // 上記の書き方の場合、switchの外では参照できない

	t := time.Now()
	// switchの後に値を宣言しない書き方もある
	switch {
	case t.Hour() < 12:
		fmt.Println(("Morning"))
	case t.Hour() < 17:
		fmt.Println(("Afternoon"))
	}
}
