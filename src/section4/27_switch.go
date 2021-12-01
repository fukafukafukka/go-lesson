package main

import "fmt"

func getOsName() string {
	return "default"
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
	case "default":
		fmt.Println("default")
	}
}
